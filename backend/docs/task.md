# Task — E-Commerce Backend

> Status: **Done**  
> Stack: Go + GORM + SQLite + JWT  
> Pattern: Modular Architecture (Handler → Service → Repository)  
> Module: `github.com/havilz/ecommerce-backend`

---

## FASE 1 — FOUNDATION

- [x] [0] Buat file `task.md` di folder `backend/docs`
- [x] [1] Setup project Go
  - [x] `go mod init github.com/havilz/ecommerce-backend`
  - [x] `go get` semua dependencies (gorm, sqlite driver, jwt, bcrypt, godotenv)
- [x] [2] Buat GORM Models di `models/`
  - [x] `models/user.go`
  - [x] `models/product.go`
  - [x] `models/cart.go`
  - [x] `models/order.go`
  - [x] `models/models.go`
- [x] [3] `pkg/database/database.go`
- [x] [4] `config/config.go`
- [x] [5] `pkg/response/response.go`
- [x] [6] `pkg/middleware/auth.go`
- [x] [7] `cmd/main.go` skeleton

---

## FASE 2 — AUTH MODULE

- [x] [8]  `internal/modules/auth/dto.go`
- [x] [9]  `internal/modules/auth/repository.go`
- [x] [10] `internal/modules/auth/service.go`
- [x] [11] `internal/modules/auth/handler.go`
- [x] [12] Test auth endpoint (curl / Postman)

---

## FASE 3 — PRODUCT MODULE

- [x] [13] `internal/modules/product/dto.go`
- [x] [14] `internal/modules/product/repository.go`
- [x] [15] `internal/modules/product/service.go`
- [x] [16] `internal/modules/product/handler.go`
- [x] [17] Seed beberapa produk ke SQLite

---

## FASE 4 — CART MODULE

- [x] [18] `internal/modules/cart/dto.go`
- [x] [19] `internal/modules/cart/repository.go`
- [x] [20] `internal/modules/cart/service.go`
- [x] [21] `internal/modules/cart/handler.go`

---

## FASE 5 — ORDER MODULE

- [x] [22] `internal/modules/order/dto.go`
- [x] [23] `internal/modules/order/repository.go`
- [x] [24] `internal/modules/order/service.go`
- [x] [25] `internal/modules/order/handler.go`

---

## FASE 6 — ADMIN MODULE

- [x] [26] `internal/modules/admin/dto.go`
- [x] [27] `internal/modules/admin/repository.go`
- [x] [28] `internal/modules/admin/service.go`
- [x] [29] `internal/modules/admin/handler.go`

---

## FASE 7 - CATEGORY MODULE
- [ ] [30] `internal/modules/category/dto.go`
- [ ] [31] `internal/modules/category/repository.go`
- [ ] [32] `internal/modules/category/service.go`
- [ ] [33] `internal/modules/category/handler.go`

---

## FASE 8 — WIRE & FINALIZE

- [x] [34] Update `cmd/main.go` — register semua module handler
- [x] [35] Test end-to-end semua endpoint

---

# TASK - E-COMMERCE UPDATE

## FASE 9 — RATE LIMITING

- [x] [36] Buat `pkg/middleware/ratelimit.go` — rate limiter middleware (in-memory, per-IP)
  - [x] Struct `RateLimiter` dengan configurable: max requests, window duration
  - [x] Cleanup goroutine untuk expired entries
  - [x] Helper function: `NewRateLimiter(max int, window time.Duration)`
- [x] [37] Terapkan rate limiter ke **Auth** routes (`/auth/register`, `/auth/login`)
  - [x] Limit ketat: 5 req/menit per IP (brute-force protection)
- [x] [38] Terapkan rate limiter ke **Checkout** route (`/orders/checkout`)
  - [x] Limit: 10 req/menit per IP (prevent duplicate orders)
- [x] [39] Terapkan rate limiter ke **Order** routes (`/orders`, `/orders/{id}`)
  - [x] Limit: 30 req/menit per IP
- [x] [40] Terapkan rate limiter ke **Cart** routes (`/cart`)
  - [x] Limit: 30 req/menit per IP
- [x] [41] Terapkan rate limiter ke **Admin** routes (`/admin/*`)
  - [x] Limit: 30 req/menit per IP
- [x] [42] Terapkan rate limiter ke **Products** routes (public read)
  - [x] Limit: 60 req/menit per IP (more relaxed for public browsing)
- [x] [43] Terapkan rate limiter ke **Categories** routes (public read)
  - [x] Limit: 60 req/menit per IP
- [ ] [44] Test rate limiting — pastikan return `429 Too Many Requests`

---

## FASE 10 — DATABASE INDEXING

- [x] [45] Tambah index pada `models/user.go`
  - [x] `Email` — sudah ada `uniqueIndex` (verifikasi saja)
- [x] [46] Tambah index pada `models/product.go`
  - [x] `CategoryID` — index untuk JOIN/filter by category
  - [x] `Name` — index untuk search query (`LIKE` query)
  - [x] `CreatedAt` — index untuk sorting
- [x] [47] Tambah index pada `models/cart.go`
  - [x] `UserID` — index untuk lookup cart by user
  - [x] Composite index `UserID + ProductID` — unique constraint (prevent duplicate cart item)
- [x] [48] Tambah index pada `models/order.go`
  - [x] `UserID` — sudah ada `index` (verifikasi saja)
  - [x] `Status` — index untuk filter order by status (admin)
  - [x] `CreatedAt` — index untuk sorting `ORDER BY created_at DESC`
  - [x] `OrderItem.OrderID` — sudah ada `index` (verifikasi saja)
  - [x] `OrderItem.ProductID` — index untuk JOIN product
- [x] [49] Tambah index pada `models/category.go`
  - [x] `Name` — sudah ada `uniqueIndex` (verifikasi saja)
- [x] [50] Jalankan migrasi & verifikasi index terbuat di SQLite
- [x] [51] Test query performance — pastikan tidak ada regresi

---

## FASE 11 — FINALIZE UPDATE

- [ ] [52] Update Swagger docs jika ada perubahan response (429)
- [ ] [53] Test end-to-end semua endpoint setelah update

---

## FASE 12 — SECURITY ENHANCEMENTS

- [x] [54] Tambah dependency `bluemonday` untuk sanitasi HTML/XSS
- [x] [55] Buat package utility `pkg/security` (sanitasi teks & SQL wildcard escape)
- [x] [56] Terapkan sanitasi input pada registrasi user (Auth Module) & regex email
- [x] [57] Terapkan sanitasi input & validasi skema ImageURL pada produk (Admin Module)
- [x] [58] Terapkan sanitasi input pada nama kategori (Category Module)
- [x] [59] Terapkan sanitasi input pada alamat checkout (Order Module)
- [x] [60] Amankan query pencarian produk dengan SQL wildcard escape
- [x] [61] Tulis unit test untuk verifikasi fungsionalitas security

