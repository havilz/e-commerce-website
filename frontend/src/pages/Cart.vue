<template>
  <div class="container-page py-8">
    <h1 class="font-heading font-bold text-2xl text-neutral-800 mb-6">
      Keranjang Belanja
      <span class="text-base text-neutral-400 font-normal ml-2">({{ cartStore.totalItems }} item)</span>
    </h1>

    <!-- Loading -->
    <div v-if="cartStore.loading" class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <div class="lg:col-span-2 card p-5 space-y-4">
        <div v-for="n in 3" :key="n" class="flex gap-4 animate-pulse">
          <div class="w-20 h-20 bg-neutral-200 rounded-lg flex-shrink-0" />
          <div class="flex-1 space-y-2">
            <div class="h-4 bg-neutral-200 rounded w-3/4" />
            <div class="h-4 bg-neutral-200 rounded w-1/2" />
            <div class="h-4 bg-neutral-200 rounded w-1/4" />
          </div>
        </div>
      </div>
      <div class="card p-5 h-48 animate-pulse bg-neutral-100" />
    </div>

    <!-- Empty cart -->
    <EmptyState
      v-else-if="!cartStore.items.length"
      :image="emptyCartImg"
      title="Keranjang kamu kosong"
      description="Yuk, tambahkan produk favorit kamu ke keranjang!"
    >
      <RouterLink to="/products">
        <BaseButton variant="primary" size="lg">Mulai Belanja</BaseButton>
      </RouterLink>
    </EmptyState>

    <!-- Cart items + summary -->
    <div v-else class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Items list -->
      <div class="lg:col-span-2 card p-5">
        <CartItem
          v-for="item in cartStore.items"
          :key="item.id"
          :item="item"
          :loading="actionLoading"
          @update="handleUpdate"
          @remove="handleRemove"
        />
      </div>

      <!-- Summary -->
      <CartSummary
        :total-items="cartStore.totalItems"
        :total-price="cartStore.totalPrice"
        @checkout="router.push('/checkout')"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '@/stores/cart.store'
import { useCart }      from '@/composables/useCart'
import CartItem         from '@/components/cart/CartItem.vue'
import CartSummary      from '@/components/cart/CartSummary.vue'
import EmptyState       from '@/components/ui/EmptyState.vue'
import BaseButton       from '@/components/ui/BaseButton.vue'
import emptyCartImg     from '@/assets/images/empty-cart.svg'

const router      = useRouter()
const cartStore   = useCartStore()
const { removeFromCart, updateQuantity } = useCart()
const actionLoading = ref(false)

async function handleUpdate(itemId, qty) {
  actionLoading.value = true
  try {
    await updateQuantity(itemId, qty)
  } finally {
    actionLoading.value = false
  }
}

async function handleRemove(itemId) {
  actionLoading.value = true
  try {
    await removeFromCart(itemId)
  } finally {
    actionLoading.value = false
  }
}

onMounted(() => cartStore.fetchCart())
</script>
