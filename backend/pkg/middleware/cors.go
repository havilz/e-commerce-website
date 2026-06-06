package middleware

import (
	"net/http"
	"strings"
)

// CORSMiddleware hanya mengizinkan request dari origin yang ada di whitelist.
// allowedOrigins diisi dari config ALLOWED_ORIGINS di .env
// Contoh: "http://localhost:5173,https://tokomu.com"
func CORSMiddleware(allowedOrigins string) func(http.Handler) http.Handler {
	// Pecah string menjadi slice, trim spasi
	origins := strings.Split(allowedOrigins, ",")
	originSet := make(map[string]bool, len(origins))
	for _, o := range origins {
		originSet[strings.TrimSpace(o)] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Jika origin ada di whitelist, izinkan
			if originSet[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				// Vary header penting agar browser cache tidak salah
				w.Header().Set("Vary", "Origin")
			}

			// Preflight request (OPTIONS) — jawab langsung tanpa lanjut ke handler
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
