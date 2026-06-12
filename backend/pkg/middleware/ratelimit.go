package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/havilz/ecommerce-backend/pkg/response"
)

// visitor tracks request counts for a single IP address.
type visitor struct {
	count    int
	lastSeen time.Time
}

// RateLimiter provides per-IP rate limiting using an in-memory store.
type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.RWMutex
	max      int           // maximum requests allowed within the window
	window   time.Duration // sliding window duration
	stopCh   chan struct{} // signal to stop cleanup goroutine
}

// NewRateLimiter creates a new rate limiter with the given max requests per window.
// It also starts a background goroutine to clean up expired entries every minute.
func NewRateLimiter(max int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		max:      max,
		window:   window,
		stopCh:   make(chan struct{}),
	}

	go rl.cleanupExpired()

	return rl
}

// cleanupExpired periodically removes visitor entries that have expired.
func (rl *RateLimiter) cleanupExpired() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rl.mu.Lock()
			for ip, v := range rl.visitors {
				if time.Since(v.lastSeen) > rl.window {
					delete(rl.visitors, ip)
				}
			}
			rl.mu.Unlock()
		case <-rl.stopCh:
			return
		}
	}
}

// Stop gracefully shuts down the cleanup goroutine.
func (rl *RateLimiter) Stop() {
	close(rl.stopCh)
}

// isAllowed checks if the given IP is within the rate limit.
// It resets the counter if the window has expired, otherwise increments it.
func (rl *RateLimiter) isAllowed(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	if !exists {
		rl.visitors[ip] = &visitor{count: 1, lastSeen: time.Now()}
		return true
	}

	// If the window has passed, reset the counter
	if time.Since(v.lastSeen) > rl.window {
		v.count = 1
		v.lastSeen = time.Now()
		return true
	}

	// Within the window — check if limit is exceeded
	if v.count >= rl.max {
		return false
	}

	v.count++
	v.lastSeen = time.Now()
	return true
}

// getIP extracts the client IP from the request.
// It checks X-Forwarded-For and X-Real-IP headers first (for proxied requests),
// then falls back to RemoteAddr.
func getIP(r *http.Request) string {
	// Check X-Forwarded-For header (first IP in the chain)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// Take the first IP from the comma-separated list
		for i := 0; i < len(xff); i++ {
			if xff[i] == ',' {
				return xff[:i]
			}
		}
		return xff
	}

	// Check X-Real-IP header
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Fall back to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// Limit returns an HTTP middleware that enforces the rate limit.
// When the limit is exceeded, it responds with 429 Too Many Requests
// and includes a Retry-After header.
func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)

		if !rl.isAllowed(ip) {
			w.Header().Set("Retry-After", rl.window.String())
			response.Error(w, http.StatusTooManyRequests, "Too many requests. Please try again later.")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// LimitFunc is a convenience wrapper for http.HandlerFunc.
func (rl *RateLimiter) LimitFunc(next http.HandlerFunc) http.Handler {
	return rl.Limit(http.HandlerFunc(next))
}
