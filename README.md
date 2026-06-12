# ShopKu - E-Commerce Platform

A modern, high-performance e-commerce platform built with a modular Go backend and a highly responsive Vue 3 frontend.

<div align="center">

  ![Go](https://img.shields.io/badge/Go_1.26-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
  ![Vue 3](https://img.shields.io/badge/Vue_3-%234FC08D.svg?style=for-the-badge&logo=vuedotjs&logoColor=white)
  ![SQLite](https://img.shields.io/badge/SQLite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)
  ![Vite](https://img.shields.io/badge/Vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white)
  ![TailwindCSS](https://img.shields.io/badge/Tailwind_CSS-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)
  ![JWT](https://img.shields.io/badge/JWT-black.svg?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white)

</div>

---

## Repository Structure

Proyek ini menggunakan struktur monorepo untuk menyatukan repositori frontend dan backend dalam satu workspace:

```text
e-commerce-website/
├── backend/          # RESTful API Service (Go)
│   ├── cmd/          # Entry point aplikasi (main.go)
│   ├── config/       # Konfigurasi environment (.env)
│   ├── database/     # SQLite database & seed scripts
│   ├── docs/         # Swagger API Documentation & Task List
│   ├── internal/     # Core logic modul (Auth, Cart, Order, Product, Category, Admin)
│   ├── models/       # Skema model database GORM
│   └── pkg/          # Package helper (Database, Middleware, Security, Response)
│
└── frontend/         # Single Page Application (Vue 3)
    ├── src/
    │   ├── components/ # Reusable UI Components
    │   ├── pages/      # Halaman aplikasi (Home, Login, Cart, Admin, dll.)
    │   ├── services/   # API Service Client (Axios wrapper)
    │   ├── stores/     # State management (Pinia)
    │   ├── router/     # Vue Router + Navigation Guards
    │   └── main.js     # Entry point Vue app
```

---

## Backend (Go + GORM + SQLite)

<div align="center">
  <img src="https://golang.org/doc/gopher/fiveyears.jpg" alt="Go Gophers" width="280" />
  <br />
  <br />

  ![Go](https://img.shields.io/badge/Go_1.26-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
  ![GORM](https://img.shields.io/badge/GORM-F16822.svg?style=for-the-badge&logo=gorm&logoColor=white)
  ![SQLite](https://img.shields.io/badge/SQLite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)
  ![JWT](https://img.shields.io/badge/JWT-black.svg?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white)
  ![Swagger](https://img.shields.io/badge/Swagger-%2385EA2D.svg?style=for-the-badge&logo=swagger&logoColor=black)
  ![bcrypt](https://img.shields.io/badge/bcrypt-Security-informational?style=for-the-badge&logo=letsencrypt&logoColor=white)

</div>

Service RESTful API tangguh yang mengutamakan keamanan data, kecepatan kueri, dan pembatasan traffic secara cerdas. Dibangun dengan **native Go** (tanpa framework HTTP) menggunakan `net/http` Go 1.22+ dengan pattern-based routing.

### Tech Stack & Packages

| Layer | Package | Versi | Keterangan |
|-------|---------|-------|------------|
| Language | [Go](https://golang.org/) | 1.26.4 | Statically typed, native `net/http` router |
| ORM | [gorm.io/gorm](https://gorm.io/) | v1.31.1 | Auto-migrate, relasi, parameterized query |
| Database Driver | [glebarez/sqlite](https://github.com/glebarez/sqlite) | v1.11.0 | Pure-Go SQLite, **CGO-free** |
| Authentication | [golang-jwt/jwt v5](https://github.com/golang-jwt/jwt) | v5.3.1 | Stateless Bearer token auth |
| Password Hashing | [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) (bcrypt) | v0.52.0 | Industry-standard password hashing |
| XSS Sanitization | [microcosm-cc/bluemonday](https://github.com/microcosm-cc/bluemonday) | v1.0.27 | HTML/XSS sanitizer untuk semua input text |
| Env Management | [joho/godotenv](https://github.com/joho/godotenv) | v1.5.1 | Load `.env` ke environment |
| API Docs | [swaggo/swag](https://github.com/swaggo/swag) + [http-swagger](https://github.com/swaggo/http-swagger) | v1.16.6 / v2.0.2 | Auto-generate Swagger UI dari komentar Go |
| Rate Limiting | Custom in-memory per-IP | — | Sliding window, cleanup goroutine |

### Architecture Pattern

```
HTTP Request
     │
     ▼
┌─────────────────────────────────┐
│       MIDDLEWARE LAYER           │
│  JWT Auth · CORS · Rate Limit    │
└────────────────┬────────────────┘
                 │
                 ▼
┌─────────────────────────────────┐
│  ROUTER (net/http Go 1.22+)     │
│     /api/v1/{resource}          │
└────────────────┬────────────────┘
                 │
        ┌────────┴────────┐
        ▼                 ▼
   HANDLER (HTTP I/O) → SERVICE (Business Logic) → REPOSITORY (DB Query)
                                                          │
                                                          ▼
                                                   GORM + SQLite
```

### Setup & Instalasi

1. Masuk ke folder backend:
   ```bash
   cd backend
   ```
2. Buat file `.env`:
   ```ini
   PORT=8080
   DB_PATH=database/ecommerce.db
   JWT_SECRET=supersecretkeyyangpanjangdanaman
   JWT_EXPIRY_HOURS=24
   ALLOWED_ORIGINS=http://localhost:5173
   ```
3. Unduh semua dependencies:
   ```bash
   go mod download
   ```
4. (Opsional) Seed data awal ke database:
   ```bash
   go run database/seed/main/main.go
   ```
5. Jalankan server backend:
   ```bash
   go run cmd/main.go
   ```
   - Server: `http://localhost:8080`
   - Swagger UI: `http://localhost:8080/swagger/index.html`

### API Endpoints

| Method | Endpoint | Auth | Keterangan |
|--------|----------|------|------------|
| `POST` | `/api/v1/auth/register` | — | Daftar akun baru |
| `POST` | `/api/v1/auth/login` | — | Login, dapat JWT token |
| `GET` | `/api/v1/auth/me` | JWT | Info user yang sedang login |
| `GET` | `/api/v1/products` | — | List produk (+ `?search=` `&category=`) |
| `GET` | `/api/v1/products/{id}` | — | Detail 1 produk |
| `GET` | `/api/v1/categories` | — | List semua kategori |
| `GET` | `/api/v1/categories/{id}` | — | Detail 1 kategori |
| `GET` | `/api/v1/cart` | JWT | Lihat isi keranjang |
| `POST` | `/api/v1/cart` | JWT | Tambah item ke cart |
| `PUT` | `/api/v1/cart/{id}` | JWT | Update quantity item |
| `DELETE` | `/api/v1/cart/{id}` | JWT | Hapus 1 item |
| `DELETE` | `/api/v1/cart` | JWT | Clear seluruh cart |
| `POST` | `/api/v1/orders/checkout` | JWT | Checkout dari cart |
| `GET` | `/api/v1/orders` | JWT | Riwayat order user |
| `GET` | `/api/v1/orders/{id}` | JWT | Detail 1 order |
| `POST` | `/api/v1/admin/products` | JWT + Admin | Tambah produk |
| `PUT` | `/api/v1/admin/products/{id}` | JWT + Admin | Edit produk |
| `DELETE` | `/api/v1/admin/products/{id}` | JWT + Admin | Hapus produk |
| `POST` | `/api/v1/admin/categories` | JWT + Admin | Tambah kategori |
| `PUT` | `/api/v1/admin/categories/{id}` | JWT + Admin | Edit kategori |
| `DELETE` | `/api/v1/admin/categories/{id}` | JWT + Admin | Hapus kategori |
| `GET` | `/api/v1/admin/orders` | JWT + Admin | List semua order |
| `PATCH` | `/api/v1/admin/orders/{id}` | JWT + Admin | Update status order |

### Fitur & Pengamanan Backend

- **Rate Limiting (Per-IP, In-Memory)**: Sliding window counter berbasis goroutine, melindungi dari brute-force dan spam.
  - `/auth/register` & `/auth/login` → **5 req/menit**
  - `/orders/checkout` → **10 req/menit**
  - `/cart` & `/orders` → **30 req/menit**
  - `/products` & `/categories` → **60 req/menit**
  - Admin endpoints → **30 req/menit**
- **HTML Sanitization & XSS Protection**: Semua input text (nama user, deskripsi produk, alamat order) dibersihkan menggunakan `bluemonday` sebelum disimpan ke database.
- **SQL Injection Prevention**: Parameterized Query GORM secara penuh, ditambah escape karakter `%` dan `_` pada fitur pencarian produk.
- **CORS Whitelist**: Hanya origin yang terdaftar di `ALLOWED_ORIGINS` yang diizinkan akses API.
- **Database Indexing**: Kueri pencarian, filter kategori, dan sorting dioptimasi dengan index pada SQLite.
- **price_at_purchase**: Harga saat checkout disimpan terpisah di `order_items` sehingga history order tetap akurat meski harga produk berubah.

---

## Frontend (Vue 3 + Vite + Tailwind)

<div align="center">
  <img src="https://vuejs.org/images/logo.png" alt="Vue.js Logo" width="80" />
  <br />
  <br />

  ![Vue 3](https://img.shields.io/badge/Vue_3-%234FC08D.svg?style=for-the-badge&logo=vuedotjs&logoColor=white)
  ![Vite](https://img.shields.io/badge/Vite_5-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white)
  ![TailwindCSS](https://img.shields.io/badge/Tailwind_CSS_v3-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)
  ![Pinia](https://img.shields.io/badge/Pinia-%23FFE05D.svg?style=for-the-badge&logo=pinia&logoColor=black)
  ![Vue Router](https://img.shields.io/badge/Vue_Router_4-%234FC08D.svg?style=for-the-badge&logo=vuedotjs&logoColor=white)
  ![Axios](https://img.shields.io/badge/Axios-%235A29E4.svg?style=for-the-badge&logo=axios&logoColor=white)
  ![VeeValidate](https://img.shields.io/badge/VeeValidate-%23FF4154.svg?style=for-the-badge&logo=vueuse&logoColor=white)
  ![Yup](https://img.shields.io/badge/Yup-Schema_Validation-blueviolet?style=for-the-badge)
  ![Lucide](https://img.shields.io/badge/Lucide_Icons-%23F56565.svg?style=for-the-badge&logo=lucide&logoColor=white)
  ![PostCSS](https://img.shields.io/badge/PostCSS-%23DD3A0A.svg?style=for-the-badge&logo=postcss&logoColor=white)

</div>

Antarmuka e-commerce modern, responsif, dan interaktif dengan desain UI premium dan micro-animations. Terinspirasi dari gaya consumer e-commerce (Tokopedia/Shopee).

### Tech Stack & Packages

| Layer | Package | Versi | Keterangan |
|-------|---------|-------|------------|
| Framework | [Vue 3](https://vuejs.org/) (Composition API) | ^3.4.0 | `<script setup>` syntax, reactive & composable |
| Build Tool | [Vite](https://vitejs.dev/) | ^5.0.0 | HMR super cepat, optimized production build |
| Styling | [Tailwind CSS](https://tailwindcss.com/) | ^3.4.0 | Utility-first, custom design token |
| State Management | [Pinia](https://pinia.vuejs.org/) | ^2.1.0 | Official Vue 3 state manager |
| Routing | [Vue Router](https://router.vuejs.org/) | ^4.3.0 | SPA routing + navigation guards |
| HTTP Client | [Axios](https://axios-http.com/) | ^1.6.0 | Interceptor untuk auto-attach JWT token |
| Form Validation | [VeeValidate](https://vee-validate.logaretm.com/) | ^4.12.0 | Schema-based form validation |
| Validation Schema | [Yup](https://github.com/jquense/yup) | ^1.4.0 | TypeScript-first schema builder |
| Icons | [Lucide Vue Next](https://lucide.dev/) | ^0.378.0 | Clean & consistent SVG icon set |
| Notifications | [Vue Toastification](https://github.com/Maronato/vue-toastification) | ^2.0.0-rc.5 | Toast feedback untuk aksi user |
| CSS Processing | [PostCSS](https://postcss.org/) + [Autoprefixer](https://github.com/postcss/autoprefixer) | ^8.4.0 / ^10.4.0 | Cross-browser CSS compatibility |
| Vue Plugin | [@vitejs/plugin-vue](https://github.com/vitejs/vite-plugin-vue) | ^5.0.0 | Vite integration untuk Single File Components |

### Setup & Instalasi

1. Masuk ke folder frontend:
   ```bash
   cd frontend
   ```
2. Buat file `.env`:
   ```ini
   VITE_API_URL=http://localhost:8080/api/v1
   VITE_APP_NAME=ShopKu
   VITE_APP_VERSION=1.0.0
   ```
3. Install dependencies:
   ```bash
   npm install
   ```
4. Jalankan server development:
   ```bash
   npm run dev
   ```
   - Aplikasi: `http://localhost:5173`

5. Build untuk production:
   ```bash
   npm run build
   ```

### Halaman Aplikasi

| Route | Halaman | Auth |
|-------|---------|------|
| `/` | Home — Hero banner + kategori + produk terbaru | — |
| `/products` | Product List — Grid + search + filter kategori | — |
| `/products/:id` | Product Detail — Detail + add to cart | — |
| `/login` | Login | — |
| `/register` | Register | — |
| `/cart` | Keranjang Belanja | JWT |
| `/checkout` | Form alamat + ringkasan order | JWT |
| `/orders` | Riwayat order user | JWT |
| `/orders/:id` | Detail 1 order | JWT |
| `/order-success` | Konfirmasi order berhasil | JWT |
| `/admin` | Admin Dashboard — Stats & recent orders | JWT + Admin |
| `/admin/products` | Admin Products — CRUD produk | JWT + Admin |
| `/admin/orders` | Admin Orders — List & update status | JWT + Admin |

### Fitur Keamanan Sisi Client

- **Peringatan Rate Limiting Terpusat**: Axios response interceptor otomatis menangkap error HTTP `429` dan memunculkan toast *"Terlalu banyak permintaan. Silakan coba lagi nanti."*
- **Auto Logout pada 401**: Interceptor mendeteksi token kedaluwarsa/tidak valid, menghapus session, dan redirect ke `/login`.
- **Validasi Format Email**: Form pendaftaran memvalidasi email menggunakan regex RFC 5322 via VeeValidate + Yup sebelum API call.
- **Validasi URL Gambar**: Input gambar produk diwajibkan berprotokol `http://` atau `https://` untuk mencegah XSS berbasis link dinamis.
- **Navigation Guards**: Route yang membutuhkan autentikasi atau role admin dilindungi oleh Vue Router `beforeEach` guard.

---

## Database Schema

```
┌─────────────┐         ┌──────────────────┐        ┌───────────────┐
│    users    │         │    products       │        │  categories   │
├─────────────┤         ├──────────────────┤        ├───────────────┤
│ id (PK)     │         │ id (PK)           │        │ id (PK)       │
│ name        │         │ name              │        │ name          │
│ email       │         │ description       │        │ description   │
│ password    │         │ price             │        │ created_at    │
│ role        │         │ stock             │        └───────────────┘
│ created_at  │         │ image_url         │
└──────┬──────┘         │ category_id (FK) ─┼──────► categories.id
       │                │ created_at        │
       │                └────────┬──────────┘
       │                         │
       ▼                         ▼
┌─────────────────┐    ┌──────────────────┐
│   cart_items    │    │     orders        │
├─────────────────┤    ├──────────────────┤
│ id (PK)         │    │ id (PK)           │
│ user_id (FK)    │    │ user_id (FK)      │
│ product_id (FK) │    │ total_price       │
│ quantity        │    │ status            │
│ created_at      │    │ address           │
└─────────────────┘    │ created_at        │
                       └────────┬──────────┘
                                │
                                ▼
                       ┌──────────────────┐
                       │   order_items     │
                       ├──────────────────┤
                       │ id (PK)           │
                       │ order_id (FK)     │
                       │ product_id (FK)   │
                       │ quantity          │
                       │ price_at_purchase │ ← snapshot harga saat checkout
                       └──────────────────┘
```

**Order status flow**: `pending` → `processing` → `completed` / `cancelled`

---

## Standard API Response

Semua endpoint menggunakan format JSON yang konsisten:

```json
// Success
{ "success": true, "message": "Products fetched successfully", "data": { } }

// Error
{ "success": false, "message": "Invalid credentials", "data": null }

// Pagination
{
  "success": true,
  "message": "OK",
  "data": [ ],
  "meta": { "page": 1, "limit": 10, "total": 42 }
}
```

---

<div align="center">

  Made with by **[Havilz Lating](https://havilzlating.vercel.app/)**

  © 2025 Havilz Lating. All rights reserved.

</div>
