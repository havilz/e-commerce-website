<template>
  <header class="bg-white shadow-xl sticky top-0 z-50">
    <div class="container-page">
      <div class="flex items-center h-16 gap-4">
        <!-- Logo -->
        <RouterLink to="/" class="flex items-center gap-2 flex-shrink-0">
          <span class="text-2xl">🛍️</span>
          <span class="font-heading font-extrabold text-xl text-primary-600">ShopKu</span>
        </RouterLink>

        <!-- Search bar -->
        <form @submit.prevent="handleSearch" class="flex-1 max-w-xl hidden sm:block">
          <div class="relative">
            <input
              v-model="searchQuery"
              type="text"
              placeholder="Cari produk..."
              class="input-field pr-10 py-2 text-sm"
            />
            <button
              type="submit"
              class="absolute right-3 top-1/2 -translate-y-1/2 text-neutral-400 hover:text-primary-500 transition-colors"
            >
              <Search :size="18" />
            </button>
          </div>
        </form>

        <!-- Right actions -->
        <div class="flex items-center gap-2 ml-auto">
          <!-- Cart icon -->
          <RouterLink
            to="/cart"
            class="relative p-2 rounded-full hover:bg-primary-50 transition-colors"
            aria-label="Keranjang belanja"
          >
            <ShoppingCart :size="22" class="text-neutral-600" />
            <span
              v-if="cartStore.totalItems > 0"
              class="absolute -top-0.5 -right-0.5 bg-secondary-500 text-white text-xs font-bold
                     w-5 h-5 rounded-full flex items-center justify-center leading-none"
            >
              {{ cartStore.totalItems > 99 ? '99+' : cartStore.totalItems }}
            </span>
          </RouterLink>

          <!-- User menu -->
          <div v-if="authStore.isLoggedIn" class="relative" ref="userMenuRef">
            <button
              @click="userMenuOpen = !userMenuOpen"
              class="flex items-center gap-2 p-2 rounded-full hover:bg-neutral-100 transition-colors"
            >
              <div class="w-8 h-8 bg-primary-500 rounded-full flex items-center justify-center">
                <span class="text-white text-sm font-bold">{{ userInitial }}</span>
              </div>
            </button>

            <!-- Dropdown -->
            <Transition name="fade">
              <div
                v-if="userMenuOpen"
                class="absolute right-0 top-12 w-52 bg-white rounded-xl shadow-lg border border-neutral-100 py-2 z-50 animate-slide-up"
              >
                <div class="px-4 py-2 border-b border-neutral-100">
                  <p class="text-sm font-semibold text-neutral-800 truncate">{{ authStore.user?.name }}</p>
                  <p class="text-xs text-neutral-400 truncate">{{ authStore.user?.email }}</p>
                </div>
                <RouterLink
                  v-if="authStore.isAdmin"
                  to="/admin"
                  @click="userMenuOpen = false"
                  class="flex items-center gap-2 px-4 py-2 text-sm hover:bg-neutral-50 transition-colors"
                >
                  <LayoutDashboard :size="16" class="text-accent-500" />
                  Admin Panel
                </RouterLink>
                <RouterLink
                  to="/orders"
                  @click="userMenuOpen = false"
                  class="flex items-center gap-2 px-4 py-2 text-sm hover:bg-neutral-50 transition-colors"
                >
                  <Package :size="16" class="text-primary-500" />
                  Pesanan Saya
                </RouterLink>
                <button
                  @click="handleLogout"
                  class="w-full flex items-center gap-2 px-4 py-2 text-sm text-red-500 hover:bg-red-50 transition-colors"
                >
                  <LogOut :size="16" />
                  Logout
                </button>
              </div>
            </Transition>
          </div>

          <!-- Guest actions -->
          <template v-else>
            <RouterLink to="/login" class="btn-ghost text-sm py-2">Masuk</RouterLink>
            <RouterLink to="/register" class="btn-primary text-sm py-2 px-4 hidden sm:block">Daftar</RouterLink>
          </template>
        </div>
      </div>

      <!-- Mobile search -->
      <form @submit.prevent="handleSearch" class="sm:hidden pb-3">
        <div class="relative">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Cari produk..."
            class="input-field pr-10 py-2 text-sm"
          />
          <button type="submit" class="absolute right-3 top-1/2 -translate-y-1/2 text-neutral-400">
            <Search :size="18" />
          </button>
        </div>
      </form>
    </div>
  </header>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { Search, ShoppingCart, LogOut, Package, LayoutDashboard } from 'lucide-vue-next'
import { useAuthStore } from '@/stores/auth.store'
import { useCartStore } from '@/stores/cart.store'
import { useAuth } from '@/composables/useAuth'

const router      = useRouter()
const authStore   = useAuthStore()
const cartStore   = useCartStore()
const { logout }  = useAuth()

const searchQuery  = ref('')
const userMenuOpen = ref(false)
const userMenuRef  = ref(null)

const userInitial = computed(() => {
  const name = authStore.user?.name || ''
  return name.charAt(0).toUpperCase() || 'U'
})

function handleSearch() {
  if (searchQuery.value.trim()) {
    router.push({ name: 'ProductList', query: { search: searchQuery.value.trim() } })
  }
}

async function handleLogout() {
  userMenuOpen.value = false
  logout()
}

// Close dropdown on outside click
function handleClickOutside(e) {
  if (userMenuRef.value && !userMenuRef.value.contains(e.target)) {
    userMenuOpen.value = false
  }
}

onMounted(() => document.addEventListener('click', handleClickOutside))
onUnmounted(() => document.removeEventListener('click', handleClickOutside))
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active { transition: opacity 0.15s; }
.fade-enter-from,
.fade-leave-to { opacity: 0; }
</style>
