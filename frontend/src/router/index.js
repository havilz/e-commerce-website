import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'

// Lazy-loaded pages
const Home           = () => import('@/pages/Home.vue')
const ProductList    = () => import('@/pages/ProductList.vue')
const ProductDetail  = () => import('@/pages/ProductDetail.vue')
const Login          = () => import('@/pages/Login.vue')
const Register       = () => import('@/pages/Register.vue')
const Cart           = () => import('@/pages/Cart.vue')
const Checkout       = () => import('@/pages/Checkout.vue')
const OrderSuccess   = () => import('@/pages/OrderSuccess.vue')
const Orders         = () => import('@/pages/Orders.vue')
const OrderDetail    = () => import('@/pages/OrderDetail.vue')
const AdminDashboard = () => import('@/pages/admin/AdminDashboard.vue')
const AdminProducts  = () => import('@/pages/admin/AdminProducts.vue')
const AdminOrders    = () => import('@/pages/admin/AdminOrders.vue')

const routes = [
  // Public
  { path: '/',             name: 'Home',         component: Home },
  { path: '/products',     name: 'ProductList',  component: ProductList },
  { path: '/products/:id', name: 'ProductDetail',component: ProductDetail },
  { path: '/login',        name: 'Login',        component: Login },
  { path: '/register',     name: 'Register',     component: Register },

  // Protected (requiresAuth)
  {
    path: '/cart',
    name: 'Cart',
    component: Cart,
    meta: { requiresAuth: true },
  },
  {
    path: '/checkout',
    name: 'Checkout',
    component: Checkout,
    meta: { requiresAuth: true },
  },
  {
    path: '/order-success',
    name: 'OrderSuccess',
    component: OrderSuccess,
    meta: { requiresAuth: true },
  },
  {
    path: '/orders',
    name: 'Orders',
    component: Orders,
    meta: { requiresAuth: true },
  },
  {
    path: '/orders/:id',
    name: 'OrderDetail',
    component: OrderDetail,
    meta: { requiresAuth: true },
  },

  // Admin (requiresAuth + requiresAdmin)
  {
    path: '/admin',
    name: 'AdminDashboard',
    component: AdminDashboard,
    meta: { requiresAuth: true, requiresAdmin: true },
  },
  {
    path: '/admin/products',
    name: 'AdminProducts',
    component: AdminProducts,
    meta: { requiresAuth: true, requiresAdmin: true },
  },
  {
    path: '/admin/orders',
    name: 'AdminOrders',
    component: AdminOrders,
    meta: { requiresAuth: true, requiresAdmin: true },
  },

  // 404 fallback
  { path: '/:pathMatch(.*)*', redirect: '/' },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  },
})

// Navigation Guards
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isLoggedIn) {
    return next({ name: 'Login', query: { redirect: to.fullPath } })
  }

  if (to.meta.requiresAdmin && !authStore.isAdmin) {
    return next({ name: 'Home' })
  }

  // Redirect logged-in user away from login/register
  if ((to.name === 'Login' || to.name === 'Register') && authStore.isLoggedIn) {
    return next({ name: 'Home' })
  }

  next()
})

export default router
