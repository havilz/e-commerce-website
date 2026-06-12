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
