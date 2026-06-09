<template>
  <aside class="w-64 bg-neutral-800 h-screen sticky top-0 flex flex-col flex-shrink-0 overflow-y-auto">
    <!-- Logo -->
    <div class="h-16 flex items-center px-6 border-b border-neutral-700">
      <RouterLink to="/admin" class="flex items-center gap-2">
        <span class="text-2xl">🛍️</span>
        <div>
          <span class="font-heading font-bold text-white text-sm block">ShopKu</span>
          <span class="text-xs text-neutral-400">Admin Panel</span>
        </div>
      </RouterLink>
    </div>

    <!-- Nav links -->
    <nav class="flex-1 py-6 px-3 space-y-1">
      <RouterLink
        v-for="item in navItems"
        :key="item.to"
        :to="item.to"
        class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm font-medium transition-colors"
        :class="[
          isActive(item.to)
            ? 'bg-primary-600 text-white'
            : 'text-neutral-400 hover:bg-neutral-700 hover:text-white'
        ]"
      >
        <component :is="item.icon" :size="18" />
        {{ item.label }}
      </RouterLink>
    </nav>

    <!-- Bottom: back to store + logout -->
    <div class="p-3 border-t border-neutral-700 space-y-1">
      <RouterLink
        to="/"
        class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm text-neutral-400 hover:bg-neutral-700 hover:text-white transition-colors"
      >
        <Store :size="18" />
        Kembali ke Toko
      </RouterLink>
      <button
        @click="handleLogout"
        class="w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm text-red-400 hover:bg-red-900/30 hover:text-red-300 transition-colors"
      >
        <LogOut :size="18" />
        Logout
      </button>
    </div>
  </aside>
</template>

<script setup>
import { useRoute } from 'vue-router'
import { LayoutDashboard, Package, ClipboardList, Tag, Store, LogOut } from 'lucide-vue-next'
import { useAuth } from '@/composables/useAuth'

const route       = useRoute()
const { logout }  = useAuth()

const navItems = [
  { to: '/admin',            label: 'Dashboard', icon: LayoutDashboard },
  { to: '/admin/products',   label: 'Produk',    icon: Package },
  { to: '/admin/categories', label: 'Kategori',  icon: Tag },
  { to: '/admin/orders',     label: 'Pesanan',   icon: ClipboardList },
]

function isActive(path) {
  if (path === '/admin') return route.path === '/admin'
  return route.path.startsWith(path)
}

function handleLogout() {
  logout()
}
</script>
