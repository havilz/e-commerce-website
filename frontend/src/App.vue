<template>
  <div class="min-h-screen bg-neutral-50 flex flex-col">
    <!-- Navbar — hanya tampil di halaman non-admin -->
    <Navbar v-if="!isAdminRoute" />

    <!-- Main content -->
    <main class="flex-1">
      <RouterView />
    </main>

    <!-- Footer — hanya di halaman non-admin -->
    <Footer v-if="!isAdminRoute" />
  </div>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'
import { useCartStore } from '@/stores/cart.store'
import Navbar from '@/components/layout/Navbar.vue'
import Footer from '@/components/layout/Footer.vue'

const route     = useRoute()
const authStore = useAuthStore()
const cartStore = useCartStore()

const isAdminRoute = computed(() => route.path.startsWith('/admin'))

// Sync cart saat app load (kalau sudah login)
onMounted(async () => {
  if (authStore.isLoggedIn) {
    await cartStore.fetchCart()
  }
})
</script>
