#  Dokumentasi Frontend E-Commerce — Arsitektur & Desain

> **Stack:** Vue 3 + Vite + Tailwind CSS + Pinia + Vue Router
> **Vibe:** Ceria & Berwarna — gaya e‑commerce konsumen (mirip Tokopedia/Shopee)

---

## 1. Teknologi yang Digunakan

| Layer | Teknologi | Keterangan |
|-------|-----------|------------|
| Framework | **Vue 3 (Composition API)** | Sintaks `<script setup>` yang reaktif & composable |
| Build Tool | **Vite 5** | Hot Module Reload super cepat, build optimal |
| Styling | **Tailwind CSS v3** | Utility‑first, token desain khusus via konfigurasi |
| State Management | **Pinia** | Manajer state resmi Vue 3, lebih simpel dibanding Vuex |
| Routing | **Vue Router 4** | Routing SPA dengan navigation guard |
| HTTP Client | **Axios** | Interceptor untuk auto‑attach token JWT |
| Form Validation | **VeeValidate + Yup** | Validasi berbasis skema |
| Icons | **Lucide Vue Next** | Ikon bersih & konsisten |
| Notifications | **Vue Toastification** | Toast untuk feedback aksi (tambah ke keranjang, dll) |
| Image Upload | **Native FileReader API** | Pratinjau sebelum mengunggah (admin) |

### Dependencies (`package.json`)

> Stack: **Vue 3 + Vite + Tailwind CSS + Pinia + Vue Router**  
> Vibe: **Colorful & Playful** — consumer e-commerce (Tokopedia/Shopee style)

---

## 1. Tech Stack

| Layer | Technology | Keterangan |
|-------|-----------|------------|
| Framework | **Vue 3 (Composition API)** | `<script setup>` syntax, reactive & composable |
| Build Tool | **Vite 5** | Super fast HMR, optimized build |
| Styling | **Tailwind CSS v3** | Utility-first, custom design token via config |
| State Management | **Pinia** | Vue 3 official state manager, lebih simple dari Vuex |
| Routing | **Vue Router 4** | SPA routing + navigation guards |
| HTTP Client | **Axios** | Interceptor untuk auto-attach JWT token |
| Form Validation | **VeeValidate + Yup** | Schema-based validation |
| Icons | **Lucide Vue Next** | Clean & consistent icon set |
| Notifications | **Vue Toastification** | Toast untuk feedback aksi (add to cart, dll) |
| Image Upload | **Native FileReader API** | Preview sebelum upload (admin) |

### Dependencies (`package.json`)

```json
{
  "dependencies": {
    "vue": "^3.4.0",
    "vue-router": "^4.3.0",
    "pinia": "^2.1.0",
    "axios": "^1.6.0",
    "vee-validate": "^4.12.0",
    "yup": "^1.4.0",
    "lucide-vue-next": "^0.378.0",
    "vue-toastification": "^2.0.0-rc.5"
  },
  "devDependencies": {
    "@vitejs/plugin-vue": "^5.0.0",
    "tailwindcss": "^3.4.0",
    "autoprefixer": "^10.4.0",
    "postcss": "^8.4.0"
  }
}
```

---

## 2. Design System

### 2.1 Color Palette

```
Primary (Green — CTA utama, tombol beli)
─────────────────────────────────────────
primary-50:   #f0fdf4   ← background light
primary-100:  #dcfce7   ← hover state
primary-400:  #4ade80   ← icon / badge
primary-500:  #22c55e   ← DEFAULT — tombol utama
primary-600:  #16a34a   ← hover tombol
primary-700:  #15803d   ← pressed / active

Secondary (Orange — promo, flash sale, badge diskon)
────────────────────────────────────────────────────
secondary-400: #fb923c   ← badge ringan
secondary-500: #f97316   ← DEFAULT — label promo
secondary-600: #ea580c   ← hover

Accent (Purple — wishlist, premium, new arrival)
────────────────────────────────────────────────
accent-400:  #c084fc
accent-500:  #a855f7   ← DEFAULT
accent-600:  #9333ea

Neutral (Background, text, border)
────────────────────────────────────────────────
neutral-50:  #fafafa   ← page background
neutral-100: #f5f5f5   ← card background
neutral-200: #e5e5e5   ← border / divider
neutral-400: #a3a3a3   ← placeholder text
neutral-600: #525252   ← secondary text
neutral-800: #262626   ← primary text
neutral-900: #171717   ← heading

Status Colors
────────────────────────────────────────────────
success: #22c55e    ← order completed
warning: #f59e0b    ← pending / processing
error:   #ef4444    ← out of stock / error
info:    #3b82f6    ← informasi
```

### 2.2 Typography

```
Font Family:
  Heading  → "Plus Jakarta Sans" (Google Fonts) — bold, modern
  Body     → "Inter" (Google Fonts) — readable, clean

Scale (Tailwind):
  text-xs:   12px  → label, badge, caption
  text-sm:   14px  → secondary text, meta info
  text-base: 16px  → body text default
  text-lg:   18px  → card title, section subtitle
  text-xl:   20px  → page subtitle
  text-2xl:  24px  → page title
  text-3xl:  30px  → hero heading
  text-4xl:  36px  → landing hero

Font Weight:
  font-normal:    400  → body
  font-medium:    500  → label, nav item
  font-semibold:  600  → card title, button
  font-bold:      700  → heading, price
  font-extrabold: 800  → hero title, promo
```

### 2.3 Spacing & Layout

```
Container: max-w-7xl mx-auto px-4 sm:px-6 lg:px-8

Grid System:
  Product Grid  → grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4
  Admin Layout  → sidebar (w-64) + main (flex-1)
  Auth Pages    → max-w-md mx-auto (centered card)

Spacing Scale (dipakai konsisten):
  p-2  (8px)   → badge, tag padding
  p-3  (12px)  → button padding kecil
  p-4  (16px)   → card padding

## 3. Clean Architecture Overview

### Layers
## 3. Gambaran Clean Architecture

### Lapisan
- **Domain**: Entitas bisnis, model, dan logika inti. TypeScript murni tanpa ketergantungan pada Vue, Axios, atau framework lainnya.
- **Application (Use Cases)**: Layanan yang mengorkestrasi objek domain, mengimplementasikan use‑case bisnis, dan mengekspos antarmuka (port) untuk akses data.
- **Infrastructure / Data**: Implementasi konkret dari port, misalnya klien Axios, adaptor penyimpanan, atau SDK pihak ketiga.
- **Presentation**: Komponen Vue, halaman, router, dan Pinia store. Hanya bergantung pada antarmuka layer Application.

### Struktur Folder
```text
src/
├─ domain/
│   └─ models/          # Definisi entitas
│   └─ valueObjects/    # Value objects, enums
│   └─ services/        # Service domain murni
├─ application/
│   └─ useCases/        # Aksi bisnis (mis. AddToCartUseCase)
│   └─ ports/           # Interface untuk akses data
├─ infrastructure/
│   └─ api/             # Klien Axios yang mengimplementasikan port
│   └─ storage/         # Adapter penyimpanan lokal
├─ presentation/
│   └─ components/      # UI komponen yang dapat dipakai ulang
│   └─ pages/           # Komponen halaman Vue
│   └─ router/          # Konfigurasi Vue Router
│   └─ store/           # Pinia store (wrapper tipis untuk use‑cases)
└─ assets/              # Gambar, font, asset statis
```

### Arah Dependensi
- Layer dalam **tidak boleh** mengimpor layer luar.
- Presentation mengimpor port Application; Application mengimpor Domain; Infrastructure mengimplementasikan port Application.

### Pedoman
- Simpan logika UI di layer Presentation; aturan bisnis tetap di Domain/Application.
- Tulis unit test untuk Domain dan Application tanpa memerlukan Vue.
- Gunakan dependency injection (melalui parameter konstruktor atau plugin Pinia) untuk menyuntikkan implementasi konkret.

- **Application (Use Cases)**: Services that orchestrate domain objects, implement business use‑cases, and expose interfaces (ports) for data access.
- **Infrastructure / Data**: Concrete implementations of the ports, e.g., Axios API clients, storage adapters, third‑party SDKs.
- **Presentation**: Vue components, pages, router, and Pinia stores. Depends only on the Application layer interfaces.

### Folder Structure
```text
src/
├─ domain/
│   └─ models/          # Entity definitions
│   └─ valueObjects/    # Value objects, enums
│   └─ services/        # Pure domain services
├─ application/
│   └─ useCases/        # Business actions (e.g., AddToCartUseCase)
│   └─ ports/           # Interfaces for data access
├─ infrastructure/
│   └─ api/             # Axios clients implementing ports
│   └─ storage/         # Local storage adapters
├─ presentation/
│   └─ components/      # Reusable UI components
│   └─ pages/           # Vue page components
│   └─ router/          # Vue Router setup
│   └─ store/           # Pinia stores (thin wrappers around use‑cases)
└─ assets/              # Images, fonts, static assets
```

### Dependency Direction
- Inner layers **must not** import from outer layers.
- Presentation imports Application ports; Application imports Domain; Infrastructure implements Application ports.

### Guidelines
- Keep UI logic in Presentation layer; business rules stay in Domain/Application.
- Write unit tests for Domain and Application without requiring Vue.
- Use dependency injection (via constructor parameters or Pinia plugins) to supply concrete implementations.

---

  p-6  (24px)  → section padding
  p-8  (32px)  → page section
  gap-4        → grid gap standard
  gap-6        → section gap
```

### 2.4 Component Design Tokens

```
Border Radius:
  rounded      → 4px   tombol kecil, badge
  rounded-lg   → 8px   card, input
  rounded-xl   → 12px  product card
  rounded-2xl  → 16px  modal, drawer
  rounded-full → 9999px pill badge, avatar

Shadow:
  shadow-sm    → card default
  shadow-md    → card hover
  shadow-lg    → modal, dropdown
  shadow-xl    → sticky navbar

Transition:
  transition-all duration-200 ease-in-out  → default
  transition-transform duration-300         → slide animasi
```

### 2.5 UI Components Style

#### Button Variants

```html
<!-- Primary — Beli Sekarang -->
<button class="bg-primary-500 hover:bg-primary-600 active:bg-primary-700
               text-white font-semibold px-6 py-3 rounded-lg
               transition-all duration-200 shadow-sm hover:shadow-md">
  Beli Sekarang
</button>

<!-- Secondary — Tambah ke Keranjang -->
<button class="border-2 border-primary-500 text-primary-600
               hover:bg-primary-50 font-semibold px-6 py-3 rounded-lg
               transition-all duration-200">
  + Keranjang
</button>

<!-- Danger — Hapus -->
<button class="bg-red-500 hover:bg-red-600 text-white
               font-medium px-4 py-2 rounded-lg transition-all duration-200">
  Hapus
</button>

<!-- Ghost — secondary action -->
<button class="text-neutral-600 hover:text-neutral-800
               hover:bg-neutral-100 px-4 py-2 rounded-lg
               transition-all duration-200">
  Lihat Semua
</button>
```

#### Product Card

```html
<div class="bg-white rounded-xl shadow-sm hover:shadow-md
            transition-all duration-200 overflow-hidden
            border border-neutral-100 hover:-translate-y-1 group">

  <!-- Image -->
  <div class="relative aspect-square overflow-hidden bg-neutral-100">
    <img class="w-full h-full object-cover
                group-hover:scale-105 transition-transform duration-300"/>
    <!-- Badge Diskon -->
    <span class="absolute top-2 left-2 bg-secondary-500 text-white
                 text-xs font-bold px-2 py-1 rounded">-20%</span>
    <!-- Wishlist -->
    <button class="absolute top-2 right-2 bg-white/80 backdrop-blur-sm
                   p-1.5 rounded-full shadow-sm hover:bg-white
                   transition-all duration-200"/>
  </div>

  <!-- Info -->
  <div class="p-3">
    <p class="text-sm text-neutral-600 line-clamp-2 mb-1">Nama Produk</p>
    <p class="text-base font-bold text-neutral-800">Rp 99.000</p>
    <p class="text-xs text-neutral-400 line-through">Rp 120.000</p>
    <!-- Rating -->
    <div class="flex items-center gap-1 mt-1">
      <span class="text-yellow-400 text-xs">★★★★☆</span>
      <span class="text-xs text-neutral-400">(124)</span>
    </div>
  </div>
</div>
```

#### Badge / Label

```html
<!-- Diskon -->
<span class="bg-secondary-500 text-white text-xs font-bold px-2 py-0.5 rounded">
  DISKON 20%
</span>

<!-- Flash Sale -->
<span class="bg-red-500 text-white text-xs font-bold px-2 py-0.5 rounded animate-pulse">
  ⚡ FLASH SALE
</span>

<!-- New -->
<span class="bg-accent-500 text-white text-xs font-bold px-2 py-0.5 rounded">
  NEW
</span>

<!-- Order Status -->
<span class="bg-warning/10 text-warning text-xs font-semibold px-3 py-1 rounded-full">
  Diproses
</span>
```

#### Input Field

```html
<div class="flex flex-col gap-1.5">
  <label class="text-sm font-medium text-neutral-700">Email</label>
  <input class="border border-neutral-200 rounded-lg px-4 py-3
                text-neutral-800 placeholder:text-neutral-400
                focus:outline-none focus:ring-2 focus:ring-primary-500
                focus:border-transparent transition-all duration-200
                bg-white" />
  <!-- Error state -->
  <p class="text-xs text-red-500">Email tidak valid</p>
</div>
```

---

## 3. Project Structure

```
frontend/
├── public/
│   └── favicon.ico
│
├── src/
│   ├── main.js                          # App entry, install plugins
│   ├── App.vue                          # Root component, RouterView
│   │
│   ├── router/
│   │   └── index.js                     # Route definitions + nav guards
│   │
│   ├── stores/                          # Pinia stores
│   │   ├── auth.store.js                # user, token, isLoggedIn, login(), logout()
│   │   ├── cart.store.js                # items, totalPrice, addItem(), removeItem()
│   │   └── product.store.js             # products, filters, pagination
│   │
│   ├── services/                        # Semua API call (axios)
│   │   ├── api.js                       # Axios instance + interceptors
│   │   ├── auth.service.js              # login(), register(), getMe()
│   │   ├── product.service.js           # getAll(), getById(), search()
│   │   ├── cart.service.js              # getCart(), addItem(), removeItem()
│   │   ├── order.service.js             # checkout(), getOrders(), getById()
│   │   └── admin.service.js             # CRUD product, list orders, updateStatus()
│   │
│   ├── composables/                     # Reusable logic (Vue composables)
│   │   ├── useAuth.js                   # login/logout logic
│   │   ├── useCart.js                   # cart operations
│   │   ├── useProducts.js               # fetch & filter products
│   │   └── useToast.js                  # toast notification helper
│   │
│   ├── components/
│   │   ├── layout/
│   │   │   ├── Navbar.vue               # Top navbar + cart icon + user menu
│   │   │   ├── Footer.vue               # Footer sederhana
│   │   │   └── AdminSidebar.vue         # Sidebar untuk admin panel
│   │   │
│   │   ├── product/
│   │   │   ├── ProductCard.vue          # Card produk (grid item)
│   │   │   ├── ProductGrid.vue          # Grid wrapper + skeleton loading
│   │   │   ├── ProductBadge.vue         # Badge diskon/new/flash sale
│   │   │   └── ProductSkeleton.vue      # Loading placeholder card
│   │   │
│   │   ├── cart/
│   │   │   ├── CartItem.vue             # 1 row item di cart
│   │   │   └── CartSummary.vue          # Total + tombol checkout
│   │   │
│   │   └── ui/                          # Generic reusable UI
│   │       ├── BaseButton.vue           # Button dengan variant props
│   │       ├── BaseInput.vue            # Input dengan label + error
│   │       ├── BaseBadge.vue            # Badge/pill component
│   │       ├── BaseModal.vue            # Modal wrapper
│   │       └── EmptyState.vue           # Ilustrasi kosong (cart kosong, dll)
│   │
│   ├── pages/
│   │   ├── Home.vue                     # Landing: hero + kategori + produk
│   │   ├── ProductList.vue              # List produk + filter + search
│   │   ├── ProductDetail.vue            # Detail produk + add to cart
│   │   ├── Cart.vue                     # Halaman keranjang
│   │   ├── Checkout.vue                 # Form alamat + summary order
│   │   ├── OrderSuccess.vue             # Konfirmasi order berhasil
│   │   ├── Orders.vue                   # Riwayat order user
│   │   ├── OrderDetail.vue              # Detail 1 order
│   │   ├── Login.vue                    # Form login
│   │   ├── Register.vue                 # Form register
│   │   └── admin/
│   │       ├── AdminDashboard.vue       # Stats + recent orders
│   │       ├── AdminProducts.vue        # Table produk + CRUD
│   │       └── AdminOrders.vue          # Table semua order + update status
│   │
│   └── assets/
│       ├── css/
│       │   └── main.css                 # Tailwind directives + global styles
│       └── images/
│           └── empty-cart.svg           # Ilustrasi kosong
│
├── index.html
├── vite.config.js
├── tailwind.config.js                   # Custom color tokens
├── postcss.config.js
└── package.json
```

---

## 4. Routing Structure

```javascript
// router/index.js

const routes = [
  // Public routes
  { path: '/',               name: 'Home',          component: Home },
  { path: '/products',       name: 'ProductList',   component: ProductList },
  { path: '/products/:id',   name: 'ProductDetail', component: ProductDetail },
  { path: '/login',          name: 'Login',         component: Login },
  { path: '/register',       name: 'Register',      component: Register },

  // Protected routes (requires login)
  { path: '/cart',           name: 'Cart',          component: Cart,          meta: { requiresAuth: true } },
  { path: '/checkout',       name: 'Checkout',      component: Checkout,      meta: { requiresAuth: true } },
  { path: '/orders',         name: 'Orders',        component: Orders,        meta: { requiresAuth: true } },
  { path: '/orders/:id',     name: 'OrderDetail',   component: OrderDetail,   meta: { requiresAuth: true } },
  { path: '/order-success',  name: 'OrderSuccess',  component: OrderSuccess,  meta: { requiresAuth: true } },

  // Admin routes (requires admin role)
  { path: '/admin',          name: 'AdminDashboard', component: AdminDashboard, meta: { requiresAuth: true, requiresAdmin: true } },
  { path: '/admin/products', name: 'AdminProducts',  component: AdminProducts,  meta: { requiresAuth: true, requiresAdmin: true } },
  { path: '/admin/orders',   name: 'AdminOrders',    component: AdminOrders,    meta: { requiresAuth: true, requiresAdmin: true } },
]

// Navigation Guard
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  if (to.meta.requiresAuth && !authStore.isLoggedIn) return next('/login')
  if (to.meta.requiresAdmin && authStore.user?.role !== 'admin') return next('/')
  next()
})
```

---

## 5. State Management (Pinia)

```
┌─────────────────────────────────────────────────┐
│                   auth.store.js                  │
├─────────────────────────────────────────────────┤
│ STATE:   user, token, isLoggedIn                 │
│ ACTIONS: login(), register(), logout(), fetchMe()│
│ PERSIST: localStorage (token)                    │
└─────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────┐
│                   cart.store.js                  │
├─────────────────────────────────────────────────┤
│ STATE:   items[], totalItems, totalPrice         │
│ ACTIONS: addItem(), removeItem(), updateQty()    │
│          clearCart(), syncWithServer()           │
└─────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────┐
│                 product.store.js                 │
├─────────────────────────────────────────────────┤
│ STATE:   products[], currentProduct, filters     │
│          pagination { page, limit, total }       │
│ ACTIONS: fetchAll(), fetchById(), setFilter()    │
└─────────────────────────────────────────────────┘
```

---

## 6. Axios Setup & Interceptors

```javascript
// services/api.js

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1',
  timeout: 10000,
})

// Auto-attach JWT token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})

// Handle 401 → auto logout
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      useAuthStore().logout()
      router.push('/login')
    }
    return Promise.reject(error)
  }
)
```

---

## 7. Page Wireframes (Lo-fi)

### 7.1 Home Page

```
┌────────────────────────────────────────────────────┐
│  🛍️ ShopKu        [Search............] 🛒3  👤     │  ← Navbar sticky
├────────────────────────────────────────────────────┤
│                                                    │
│  ┌──────────────────────────────────────────────┐  │
│  │  🎉 FLASH SALE HARI INI  ⚡                  │  │  ← Hero Banner
│  │  Diskon s/d 70% — Berakhir 12:00:00          │  │    bg: gradient orange→red
│  │  [ Belanja Sekarang ]                         │  │
│  └──────────────────────────────────────────────┘  │
│                                                    │
│  Kategori                                          │  ← Category pills
│  [📱 Elektronik] [👗 Fashion] [🏠 Rumah] [+]      │
│                                                    │
│  Produk Terbaru                      Lihat Semua → │
│  ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐             │  ← Product Grid
│  │ img  │ │ img  │ │ img  │ │ img  │             │    grid-cols-2/3/4
│  │-20%  │ │ NEW  │ │      │ │⚡    │             │
│  │Nama  │ │Nama  │ │Nama  │ │Nama  │             │
│  │Rp99k │ │Rp50k │ │Rp75k │ │Rp30k │             │
│  └──────┘ └──────┘ └──────┘ └──────┘             │
│                                                    │
└────────────────────────────────────────────────────┘
```

### 7.2 Product Detail Page

```
┌────────────────────────────────────────────────────┐
│  Navbar                                            │
├────────────────────────────────────────────────────┤
│                                                    │
│  Home > Elektronik > Nama Produk                   │  ← Breadcrumb
│                                                    │
│  ┌───────────────────┐  ┌────────────────────────┐ │
│  │                   │  │ Nama Produk Panjang     │ │
│  │   [   IMAGE   ]   │  │ ★★★★☆ (124 ulasan)     │ │
│  │                   │  │                         │ │
│  │ ○ ○ ● ○ ○         │  │ ~~Rp 120.000~~          │ │  ← Thumbnail dots
│  └───────────────────┘  │ Rp 99.000     -17%      │ │
│                         │                         │ │
│                         │ Stok: 42 tersedia       │ │
│                         │                         │ │
│                         │ Jumlah: [−] 1 [+]       │ │
│                         │                         │ │
│                         │ [+ Keranjang] (outline) │ │
│                         │ [  Beli Sekarang  ] (bg)│ │
│                         └────────────────────────┘ │
│                                                    │
│  Deskripsi Produk                                  │
│  Lorem ipsum...                                    │
│                                                    │
└────────────────────────────────────────────────────┘
```

### 7.3 Cart Page

```
┌────────────────────────────────────────────────────┐
│  Navbar                                            │
├────────────────────────────────────────────────────┤
│  🛒 Keranjang Belanja (3 item)                     │
│                                                    │
│  ┌──────────────────────────────┐ ┌─────────────┐  │
│  │ ☑  [img] Nama Produk         │ │  Ringkasan  │  │
│  │         Rp 99.000            │ │             │  │
│  │         [−] 2 [+]   🗑️      │ │ Subtotal    │  │
│  ├──────────────────────────────┤ │ Rp 198.000  │  │
│  │ ☑  [img] Nama Produk 2       │ │             │  │
│  │         Rp 50.000            │ │ Ongkir      │  │
│  │         [−] 1 [+]   🗑️      │ │ Rp 0        │  │
│  ├──────────────────────────────┤ │             │  │
│  │ ☑  [img] Nama Produk 3       │ │ Total       │  │
│  │         Rp 75.000            │ │ Rp 323.000  │  │
│  │         [−] 1 [+]   🗑️      │ │             │  │
│  └──────────────────────────────┘ │ [Checkout→] │  │
│                                   └─────────────┘  │
└────────────────────────────────────────────────────┘
```

### 7.4 Admin Dashboard

```
┌────────────────────────────────────────────────────┐
│  🛍️ ShopKu Admin                          👤 Admin │
├───────────────┬────────────────────────────────────┤
│               │                                    │
│  📊 Dashboard │  Overview                          │
│  📦 Produk    │  ┌──────┐ ┌──────┐ ┌──────┐        │
│  📋 Pesanan   │  │  42  │ │  18  │ │ 3.2M │        │  ← Stat cards
│               │  │Produk│ │Order │ │ Rev  │        │
│               │  └──────┘ └──────┘ └──────┘        │
│               │                                    │
│               │  Pesanan Terbaru                   │
│               │  ┌──────────────────────────────┐  │
│               │  │ #001  Budi  Rp99k  Pending   │  │
│               │  │ #002  Siti  Rp50k  Selesai   │  │
│               │  └──────────────────────────────┘  │
│               │                                    │
└───────────────┴────────────────────────────────────┘
```

---

## 8. tailwind.config.js

```javascript
/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js}'],
  theme: {
    extend: {
      colors: {
        primary: {
          50:  '#f0fdf4',
          100: '#dcfce7',
          400: '#4ade80',
          500: '#22c55e',
          600: '#16a34a',
          700: '#15803d',
        },
        secondary: {
          400: '#fb923c',
          500: '#f97316',
          600: '#ea580c',
        },
        accent: {
          400: '#c084fc',
          500: '#a855f7',
          600: '#9333ea',
        },
      },
      fontFamily: {
        sans:    ['Inter', 'system-ui', 'sans-serif'],
        heading: ['"Plus Jakarta Sans"', 'system-ui', 'sans-serif'],
      },
      animation: {
        'fade-in':    'fadeIn 0.2s ease-in-out',
        'slide-up':   'slideUp 0.3s ease-out',
        'bounce-in':  'bounceIn 0.4s cubic-bezier(0.68,-0.55,0.265,1.55)',
      },
      keyframes: {
        fadeIn:   { from: { opacity: 0 },             to: { opacity: 1 } },
        slideUp:  { from: { transform: 'translateY(16px)', opacity: 0 }, to: { transform: 'translateY(0)', opacity: 1 } },
        bounceIn: { from: { transform: 'scale(0.8)',  opacity: 0 }, to: { transform: 'scale(1)', opacity: 1 } },
      },
    },
  },
  plugins: [],
}
```

---

## 9. Environment Variables (`.env`)

```env
VITE_API_URL=http://localhost:8080/api/v1
VITE_APP_NAME=ShopKu
VITE_APP_VERSION=1.0.0
```

---

## 10. Flow Pengerjaan Frontend

```
[0] Buat file **task.md** di folder `frontend/docs` untuk mendefinisikan task dengan status `Draft` – prioritas utama sebelum langkah lain.
FASE 1 — SETUP PROJECT
──────────────────────────────────────────────────────
[1]  Scaffold project
     └── npm create vite@latest frontend -- --template vue
     └── cd frontend && npm install

[2]  Install dependencies
     └── npm install vue-router@4 pinia axios vee-validate yup
     └── npm install lucide-vue-next vue-toastification
     └── npm install -D tailwindcss postcss autoprefixer
     └── npx tailwindcss init -p

[3]  Setup Tailwind
     └── Edit tailwind.config.js (custom colors, fonts, animations)
     └── Edit src/assets/css/main.css (add @tailwind directives)
     └── Import Google Fonts di index.html (Inter + Plus Jakarta Sans)

[4]  Setup Vite proxy (vite.config.js)
     └── proxy: { '/api': 'http://localhost:8080' }


FASE 2 — CORE SETUP
──────────────────────────────────────────────────────
[5]  services/api.js
     └── Axios instance + request interceptor (attach token)
     └── Response interceptor (handle 401 → logout)

[6]  stores/auth.store.js
     └── State: user, token
     └── Actions: login(), logout(), fetchMe()
     └── Persist token ke localStorage

[7]  stores/cart.store.js
     └── State: items[], totalPrice
     └── Actions: addItem(), removeItem(), updateQty(), clearCart()

[8]  router/index.js
     └── Definisi semua routes
     └── Navigation guard (requiresAuth, requiresAdmin)

[9]  main.js
     └── createApp → use(router) → use(createPinia()) → use(Toast) → mount


FASE 3 — LAYOUT COMPONENTS
──────────────────────────────────────────────────────
[10] components/layout/Navbar.vue
     └── Logo, search bar, cart icon (badge count), user dropdown

[11] components/ui/BaseButton.vue
     └── Props: variant (primary/secondary/ghost/danger), size, loading

[12] components/ui/BaseInput.vue
     └── Props: label, error, type, placeholder

[13] App.vue
     └── <Navbar /> + <RouterView /> + <Footer />


FASE 4 — AUTH PAGES
──────────────────────────────────────────────────────
[14] services/auth.service.js → login(), register()
[15] pages/Login.vue         → form + validasi + redirect
[16] pages/Register.vue      → form + validasi + redirect
[17] Test: login → token tersimpan → redirect ke Home


FASE 5 — PRODUCT PAGES
──────────────────────────────────────────────────────
[18] services/product.service.js  → getAll(), getById()
[19] stores/product.store.js      → fetchAll(), fetchById()
[20] components/product/ProductCard.vue    → card UI + badge
[21] components/product/ProductSkeleton.vue → loading state
[22] pages/Home.vue               → hero banner + product grid
[23] pages/ProductList.vue        → grid + search + filter kategori
[24] pages/ProductDetail.vue      → detail + quantity selector + add to cart
[25] Test: browse produk, cek detail


FASE 6 — CART
──────────────────────────────────────────────────────
[26] services/cart.service.js    → getCart(), addItem(), removeItem()
[27] components/cart/CartItem.vue → 1 baris item
[28] components/cart/CartSummary.vue → total + tombol checkout
[29] pages/Cart.vue              → list cart + summary
[30] Navbar cart icon update realtime dari cart.store
[31] Test: add to cart dari detail page, cek di cart


FASE 7 — CHECKOUT & ORDER
──────────────────────────────────────────────────────
[32] services/order.service.js   → checkout(), getOrders()
[33] pages/Checkout.vue          → form alamat + order summary + submit
[34] pages/OrderSuccess.vue      → konfirmasi sukses + link ke orders
[35] pages/Orders.vue            → list riwayat order + status badge
[36] pages/OrderDetail.vue       → detail 1 order + list items
[37] Test: full flow checkout → success page → cek di orders


FASE 8 — ADMIN PANEL
──────────────────────────────────────────────────────
[38] components/layout/AdminSidebar.vue
[39] services/admin.service.js   → CRUD product, list orders, updateStatus()
[40] pages/admin/AdminDashboard.vue  → stat cards + recent orders
[41] pages/admin/AdminProducts.vue   → table produk + tambah/edit/hapus
[42] pages/admin/AdminOrders.vue     → table order + update status
[43] Test: login sebagai admin → akses /admin


FASE 9 — POLISH & FINALIZE
──────────────────────────────────────────────────────
[44] Tambah loading skeleton di semua fetch
[45] Tambah empty state (cart kosong, order kosong)
[46] Tambah toast notification (add to cart, checkout berhasil, dll)
[47] Responsive check: mobile (375px), tablet (768px), desktop (1280px)
[48] Test full flow sebagai user biasa & admin
[49] Update .env, cek semua hardcoded URL dihapus
[50] Push ke GitHub
     └── git init, .gitignore (node_modules, dist, .env), push
```

---
