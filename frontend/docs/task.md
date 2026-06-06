# Task Frontend E-Commerce — ShopKu

> Status: **Done**
> Stack: Vue 3 + Vite + Tailwind CSS + Pinia + Vue Router
> Tanggal: 2026-06-06

---

## Progress Overview

| Fase | Deskripsi | Status |
|------|-----------|--------|
| FASE 0 | Buat task.md | ✅ Done |
| FASE 1 | Setup Project (scaffold, dependencies, config) | ✅ Done |
| FASE 2 | Core Setup (api.js, stores, router, main.js) | ✅ Done |
| FASE 3 | Layout Components (Navbar, Footer, BaseUI) | ✅ Done |
| FASE 4 | Auth Pages (Login, Register) | ✅ Done |
| FASE 5 | Product Pages (Home, ProductList, ProductDetail) | ✅ Done |
| FASE 6 | Cart (CartItem, CartSummary, Cart page) | ✅ Done |
| FASE 7 | Checkout & Order (Checkout, Orders, OrderDetail) | ✅ Done |
| FASE 8 | Admin Panel (Dashboard, Products, Orders) | ✅ Done |
| FASE 9 | Polish & Finalize (loading, empty state, toast, responsive) | ✅ Done |

---

## Detail Task

### FASE 1 — Setup Project
- [x] [1] Scaffold project dengan Vite template Vue
- [x] [2] Install semua dependencies
- [x] [3] Setup Tailwind CSS (config + directives + fonts)
- [x] [4] Setup Vite proxy ke backend `:8080`

### FASE 2 — Core Setup
- [x] [5] `services/api.js` — Axios instance + interceptors
- [x] [6] `stores/auth.store.js` — user, token, login/logout
- [x] [7] `stores/cart.store.js` — items, addItem, removeItem
- [x] [8] `router/index.js` — routes + navigation guards
- [x] [9] `main.js` — install semua plugins

### FASE 3 — Layout Components
- [x] [10] `components/layout/Navbar.vue`
- [x] [11] `components/ui/BaseButton.vue`
- [x] [12] `components/ui/BaseInput.vue`
- [x] [13] `App.vue`

### FASE 4 — Auth Pages
- [x] [14] `services/auth.service.js`
- [x] [15] `pages/Login.vue`
- [x] [16] `pages/Register.vue`

### FASE 5 — Product Pages
- [x] [17] `services/product.service.js`
- [x] [18] `stores/product.store.js`
- [x] [19] `components/product/ProductCard.vue`
- [x] [20] `components/product/ProductSkeleton.vue`
- [x] [21] `pages/Home.vue`
- [x] [22] `pages/ProductList.vue`
- [x] [23] `pages/ProductDetail.vue`

### FASE 6 — Cart
- [x] [24] `services/cart.service.js`
- [x] [25] `components/cart/CartItem.vue`
- [x] [26] `components/cart/CartSummary.vue`
- [x] [27] `pages/Cart.vue`

### FASE 7 — Checkout & Order
- [x] [28] `services/order.service.js`
- [x] [29] `pages/Checkout.vue`
- [x] [30] `pages/OrderSuccess.vue`
- [x] [31] `pages/Orders.vue`
- [x] [32] `pages/OrderDetail.vue`

### FASE 8 — Admin Panel
- [x] [33] `components/layout/AdminSidebar.vue`
- [x] [34] `services/admin.service.js`
- [x] [35] `pages/admin/AdminDashboard.vue`
- [x] [36] `pages/admin/AdminProducts.vue`
- [x] [37] `pages/admin/AdminOrders.vue`

### FASE 9 — Polish
- [x] [38] Loading skeleton semua fetch
- [x] [39] Empty state components
- [x] [40] Toast notifications
- [x] [41] Responsive check (mobile/tablet/desktop)
- [x] [42] Environment variables cleanup

---

## Catatan
- API base URL: `http://localhost:8080/api/v1`
- App Name: **ShopKu**
- Backend sudah siap dengan Go + SQLite, port 8080
