> Stack: **Go + GORM + SQLite + JWT**  
> Pattern: **Modular Architecture (Handler → Service → Repository)**
---

## 1. Tech Stack

| Layer | Technology | Keterangan |
|-------|-----------|------------|
| Language | **Go 1.22+** | Statically typed, fast compile, great for API |
| ORM | **GORM** | Go ORM dengan auto-migrate & relasi |
| Database | **SQLite** | Zero-config, file-based, cukup untuk demo/porto |
| Authentication | **JWT (golang-jwt/jwt v5)** | Stateless auth via Bearer token |
| Password Hashing | **bcrypt (golang.org/x/crypto)** | Industry standard password hashing |
| Env Management | **godotenv (joho/godotenv)** | Load `.env` file ke environment |

### 1.1. Go Dependencies
```
gorm.io/gorm                      v1.25+    # ORM
gorm.io/driver/sqlite             v1.5+     # SQLite driver
github.com/golang-jwt/jwt/v5      v5.2+     # JWT
golang.org/x/crypto               latest    # bcrypt
github.com/joho/godotenv          v1.5+     # .env loader
```

## 2. Architechture Pattern

Pattern yang dipakai adalah **Modular Architecture** dengan pemisahan tanggung jawab tiap layer:

```
HTTP Request
     │
     ▼
┌─────────────────────────────────┐
│         MIDDLEWARE LAYER         │
│  JWT Auth · CORS · Logger        │
└────────────────┬────────────────┘
                 │
                 ▼
┌─────────────────────────────────┐
│           ROUTER LAYER           │
│         /api/v1/{resource}       │
└────────────────┬────────────────┘
                 │
        ┌────────┴────────┐
        ▼                 ▼
┌──────────────┐   ┌──────────────┐
│   HANDLER    │   │   HANDLER    │  ← Tiap module punya handler sendiri
│  (HTTP I/O)  │   │  (HTTP I/O)  │
└──────┬───────┘   └──────┬───────┘
       │                  │
       ▼                  ▼
┌──────────────┐   ┌──────────────┐
│   SERVICE    │   │   SERVICE    │  ← Business logic murni
│ (Logic Only) │   │ (Logic Only) │
└──────┬───────┘   └──────┬───────┘
       │                  │
       ▼                  ▼
┌──────────────┐   ┌──────────────┐
│  REPOSITORY  │   │  REPOSITORY  │  ← Database query only
│  (DB Query)  │   │  (DB Query)  │
└──────┬───────┘   └──────┬───────┘
       │                  │
       └────────┬──────────┘
                ▼
┌─────────────────────────────────┐
│          GORM + SQLite           │
└─────────────────────────────────┘
```

### 2.1 Prinsip Tiap Layer

| File | Tanggung Jawab | Boleh Akses |
|------|---------------|-------------|
| `handler.go` | Bind request, validasi input, tulis response | Service saja |
| `service.go` | Business logic, rules, kalkulasi | Repository saja |
| `repository.go` | Query GORM ke database | Database (GORM) saja |
| `dto.go` | Struct Request & Response shape | Dipakai Handler & Service |

> **Aturan**: Handler tidak boleh akses DB langsung. Service tidak boleh tahu soal HTTP. Repository tidak boleh ada business logic.

---

## 3. Project Struchture

```
ecommerce/
│
├── backend/
│   ├── cmd/
│   │   └── main.go                         # Entry point, wire semua module
│   │
│   ├── config/
│   │   └── config.go                       # Load .env, DB config, JWT secret
│   │
│   ├── models/                             # Folder model database
│   │   ├── user.go                         # Model User
│   │   ├── product.go                      # Model Product
│   │   ├── cart.go                         # Model CartItem
│   │   ├── order.go                        # Model Order & OrderItem
│   │   └── models.go                       # Penyatuan/registrasi semua model untuk auto-migration
│   │
│   ├── internal/
│   │   └── modules/
│   │       ├── auth/
│   │       │   ├── handler.go              # POST /register, POST /login
│   │       │   ├── service.go              # Hash password, generate JWT
│   │       │   ├── repository.go           # FindByEmail, CreateUser
│   │       │   └── dto.go                  # RegisterRequest, LoginResponse
│   │       │
│   │       ├── product/
│   │       │   ├── handler.go              # GET /products, GET /products/:id
│   │       │   ├── service.go              # Filter, pagination logic
│   │       │   ├── repository.go           # FindAll, FindByID, Search
│   │       │   └── dto.go                  # ProductResponse, ProductListResponse
│   │       │
│   │       ├── cart/
│   │       │   ├── handler.go              # GET/POST/DELETE /cart
│   │       │   ├── service.go              # Add, remove, update quantity
│   │       │   ├── repository.go           # Cart CRUD per user
│   │       │   └── dto.go                  # CartItemRequest, CartResponse
│   │       │
│   │       ├── order/
│   │       │   ├── handler.go              # POST /checkout, GET /orders
│   │       │   ├── service.go              # Create order dari cart, hitung total
│   │       │   ├── repository.go           # Order & OrderItem CRUD
│   │       │   └── dto.go                  # CheckoutRequest, OrderResponse
│   │       │
│   │       └── admin/
│   │           ├── handler.go              # CRUD product, list orders
│   │           ├── service.go              # Admin-only business logic
│   │           ├── repository.go           # Admin queries
│   │           └── dto.go                  # AdminProductRequest, OrderListResponse
│   │
│   ├── pkg/
│   │   ├── database/
│   │   │   └── database.go                 # Init GORM SQLite, AutoMigrate
│   │   ├── middleware/
│   │   │   └── auth.go                     # JWT validation middleware
│   │   └── response/
│   │       └── response.go                 # Standard JSON response helper
│   │
│   ├── .env                                # JWT_SECRET, DB_PATH, PORT
    ├── go.mod
    └── go.sum

```
---

## 4. SQLite Database Relations

### 4.1 Tabel & Kolom

```
┌─────────────┐         ┌─────────────────┐
│    users    │         │    products      │
├─────────────┤         ├─────────────────┤
│ id (PK)     │         │ id (PK)          │
│ name        │         │ name             │
│ email       │◄──┐     │ description      │
│ password    │   │     │ price            │
│ role        │   │     │ stock            │
│ created_at  │   │     │ image_url        │
└─────────────┘   │     │ category         │
                  │     │ created_at       │
                  │     └────────┬─────────┘
                  │              │
      ┌───────────┘              │
      │                          │
      ▼                          ▼
┌─────────────────┐    ┌──────────────────┐
│   cart_items    │    │     orders        │
├─────────────────┤    ├──────────────────┤
│ id (PK)         │    │ id (PK)           │
│ user_id (FK) ───┼───►│ user_id (FK)      │
│ product_id (FK)─┼─┐  │ total_price       │
│ quantity        │ │  │ status            │
│ created_at      │ │  │ address           │
└─────────────────┘ │  │ created_at        │
                    │  └────────┬──────────┘
                    │           │
                    │           ▼
                    │  ┌──────────────────┐
                    │  │   order_items     │
                    │  ├──────────────────┤
                    │  │ id (PK)           │
                    │  │ order_id (FK)     │
                    └─►│ product_id (FK)   │
                       │ quantity          │
                       │ price_at_purchase │
                       └──────────────────┘
```

### Relasi

| Relasi | Tipe | Keterangan |
|--------|------|------------|
| User → CartItem | One to Many | 1 user bisa punya banyak item di cart |
| User → Order | One to Many | 1 user bisa punya banyak order |
| Order → OrderItem | One to Many | 1 order berisi banyak item |
| Product → CartItem | One to Many | 1 produk bisa ada di banyak cart |
| Product → OrderItem | One to Many | 1 produk bisa ada di banyak order |

### 4.2 Catatan Penting
- `price_at_purchase` di `order_items` disimpan terpisah supaya kalau harga produk berubah, history order tetap akurat.
- `role` di `users`: `"user"` atau `"admin"` — dicheck di JWT middleware.
- `status` di `orders`: `"pending"` → `"processing"` → `"completed"` / `"cancelled"`.

--- 

## 5. API Endpoints

### 5.1 Auth
```
POST   /api/v1/auth/register        # Daftar akun baru
POST   /api/v1/auth/login           # Login, dapat JWT token
GET    /api/v1/auth/me              # [Protected] Info user login
```

### 5.2 Product
```
GET    /api/v1/products             # List semua produk (+ ?search= &category=)
GET    /api/v1/products/{id}        # Detail 1 produk
```

### 5.3 Cart (Protected — user only)
```
GET    /api/v1/cart                 # Lihat isi cart milik user
POST   /api/v1/cart                 # Tambah item ke cart
PUT    /api/v1/cart/{id}            # Update quantity item
DELETE /api/v1/cart/{id}            # Hapus 1 item dari cart
DELETE /api/v1/cart                 # Clear seluruh cart
```

### 5.4 Order (Protected — user only)
```
POST   /api/v1/orders/checkout      # Checkout dari cart → buat order
GET    /api/v1/orders               # Riwayat order milik user
GET    /api/v1/orders/{id}          # Detail 1 order
```

### 5.5 Admin (Protected — admin only)
```
POST   /api/v1/admin/products       # Tambah produk baru
PUT    /api/v1/admin/products/{id}  # Edit produk
DELETE /api/v1/admin/products/{id}  # Hapus produk
GET    /api/v1/admin/orders         # List semua order
PATCH  /api/v1/admin/orders/{id}    # Update status order
```

---

## 6. Flow Pengerjaan

Urutan dari fondasi → fitur, biar tidak ada dependency yang belum exist saat coding.

```
FASE 1 — FOUNDATION
─────────────────────────────────────────────────────
[0] Buat file **task.md** di folder `backend/docs` untuk mendefinisikan task dengan status `Draft` – prioritas utama sebelum langkah lain.
[1] Setup project Go
    └── go mod init github.com/{username}/ecommerce-backend
    └── go get semua dependencies

[2] models/ (GORM Models)
    └── Buat file model terpisah: user.go, product.go, cart.go, order.go
    └── Satukan registrasi model di models.go (misal: mengembalikan list model untuk AutoMigrate)

[3] pkg/database/database.go
    └── Init GORM + SQLite
    └── AutoMigrate semua model

[4] config/config.go
    └── Load .env (JWT_SECRET, PORT, DB_PATH)

[5] pkg/response/response.go
    └── Helper: Success(), Error(), Pagination()

[6] pkg/middleware/auth.go
    └── JWT middleware (ValidateToken, AdminOnly)

[7] cmd/main.go (skeleton)
    └── Init DB, setup net/http Mux (Go 1.22+), register routes placeholder


FASE 2 — AUTH MODULE
─────────────────────────────────────────────────────
[8]  internal/modules/auth/dto.go
     └── RegisterRequest, LoginRequest, AuthResponse

[9]  internal/modules/auth/repository.go
     └── FindByEmail(), CreateUser()

[10] internal/modules/auth/service.go
     └── Register (hash bcrypt), Login (compare + generate JWT)

[11] internal/modules/auth/handler.go
     └── POST /register, POST /login, GET /me

[12] Test auth endpoint via curl / Postman


FASE 3 — PRODUCT MODULE
─────────────────────────────────────────────────────
[13] internal/modules/product/dto.go
[14] internal/modules/product/repository.go
     └── FindAll (+ filter), FindByID
[15] internal/modules/product/service.go
[16] internal/modules/product/handler.go
     └── GET /products, GET /products/:id

[17] Seed beberapa produk ke SQLite (manual / seed script)


FASE 4 — CART MODULE
─────────────────────────────────────────────────────
[18] internal/modules/cart/dto.go
[19] internal/modules/cart/repository.go
[20] internal/modules/cart/service.go
     └── Add (cek stock), Update qty, Remove, Clear
[21] internal/modules/cart/handler.go
     └── Semua route cart (butuh JWT middleware)


FASE 5 — ORDER MODULE
─────────────────────────────────────────────────────
[22] internal/modules/order/dto.go
[23] internal/modules/order/repository.go
[24] internal/modules/order/service.go
     └── Checkout: ambil cart → buat order → buat order_items → clear cart → kurangi stock
[25] internal/modules/order/handler.go


FASE 6 — ADMIN MODULE
─────────────────────────────────────────────────────
[26] internal/modules/admin/dto.go
[27] internal/modules/admin/repository.go
[28] internal/modules/admin/service.go
[29] internal/modules/admin/handler.go
     └── Semua route admin (butuh AdminOnly middleware)


FASE 7 — WIRE & FINALIZE
─────────────────────────────────────────────────────
[30] Update cmd/main.go
     └── Register semua module handler ke router

[31] Test end-to-end semua endpoint

## 7. Environment Variables (`.env`)

```env
PORT=8080
DB_PATH=./ecommerce.db
JWT_SECRET=your-super-secret-key-change-this
JWT_EXPIRY_HOURS=24
```

---

## 8. Standard JSON Response

Semua endpoint return format yang sama:

```json
// Success
{
  "success": true,
  "message": "Products fetched successfully",
  "data": { ... }
}

// Error
{
  "success": false,
  "message": "Invalid credentials",
  "data": null
}

// Pagination
{
  "success": true,
  "message": "OK",
  "data": [ ... ],
  "meta": {
    "page": 1,
    "limit": 10,
    "total": 42
  }
}
```