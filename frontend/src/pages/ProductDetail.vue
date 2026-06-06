<template>
  <div class="container-page py-8">
    <!-- Skeleton loading -->
    <div v-if="loading" class="animate-pulse">
      <div class="h-4 bg-neutral-200 rounded w-48 mb-6" />
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="aspect-square bg-neutral-200 rounded-xl" />
        <div class="space-y-4">
          <div class="h-8 bg-neutral-200 rounded w-3/4" />
          <div class="h-6 bg-neutral-200 rounded w-1/4" />
          <div class="h-10 bg-neutral-200 rounded w-1/2" />
          <div class="h-4 bg-neutral-200 rounded w-full" />
          <div class="h-4 bg-neutral-200 rounded w-5/6" />
        </div>
      </div>
    </div>

    <!-- Error -->
    <div v-else-if="!product" class="text-center py-20">
      <EmptyState title="Produk tidak ditemukan" description="Produk ini mungkin sudah tidak tersedia.">
        <RouterLink to="/products">
          <BaseButton variant="primary">Lihat Produk Lain</BaseButton>
        </RouterLink>
      </EmptyState>
    </div>

    <!-- Product detail -->
    <template v-else>
      <!-- Breadcrumb -->
      <nav class="flex items-center gap-2 text-sm text-neutral-400 mb-6" aria-label="Breadcrumb">
        <RouterLink to="/" class="hover:text-primary-600 transition-colors">Home</RouterLink>
        <ChevronRight :size="14" />
        <RouterLink to="/products" class="hover:text-primary-600 transition-colors">Produk</RouterLink>
        <ChevronRight :size="14" />
        <span class="text-neutral-600 truncate max-w-xs">{{ product.name }}</span>
      </nav>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-8 lg:gap-12">
        <!-- Image -->
        <div class="space-y-3">
          <div class="aspect-square overflow-hidden rounded-xl bg-neutral-100 border border-neutral-100">
            <img
              v-if="product.image_url && !imgError"
              :src="product.image_url"
              :alt="product.name"
              class="w-full h-full object-cover"
              @error="imgError = true"
            />
            <div v-else class="w-full h-full flex items-center justify-center">
              <ImageOff :size="64" class="text-neutral-300" />
            </div>
          </div>
        </div>

        <!-- Info -->
        <div class="space-y-4">
          <!-- Category -->
          <span class="text-xs font-semibold text-primary-600 uppercase tracking-wide bg-primary-50 px-3 py-1 rounded-full">
            {{ product.category }}
          </span>

          <!-- Name -->
          <h1 class="font-heading font-bold text-2xl text-neutral-800 leading-tight">{{ product.name }}</h1>

          <!-- Price -->
          <div>
            <p class="text-3xl font-bold text-primary-600">{{ formatPrice(product.price) }}</p>
          </div>

          <!-- Stock -->
          <div class="flex items-center gap-2">
            <span
              :class="[
                'text-xs font-semibold px-3 py-1 rounded-full',
                product.stock > 0 ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'
              ]"
            >
              {{ product.stock > 0 ? `Stok: ${product.stock}` : 'Habis' }}
            </span>
          </div>

          <!-- Quantity selector -->
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-neutral-600">Jumlah:</span>
            <div class="flex items-center gap-2 border border-neutral-200 rounded-lg overflow-hidden">
              <button
                @click="qty > 1 ? qty-- : null"
                :disabled="qty <= 1"
                class="px-3 py-2 hover:bg-neutral-50 transition-colors disabled:opacity-40 text-neutral-600"
                aria-label="Kurangi jumlah"
              >
                <Minus :size="16" />
              </button>
              <span class="w-10 text-center font-semibold text-sm">{{ qty }}</span>
              <button
                @click="qty < product.stock ? qty++ : null"
                :disabled="qty >= product.stock"
                class="px-3 py-2 hover:bg-neutral-50 transition-colors disabled:opacity-40 text-neutral-600"
                aria-label="Tambah jumlah"
              >
                <Plus :size="16" />
              </button>
            </div>
          </div>

          <!-- Action buttons -->
          <div class="flex gap-3 pt-2">
            <BaseButton
              variant="secondary"
              size="lg"
              class="flex-1"
              :disabled="product.stock === 0 || cartLoading"
              :loading="cartLoading"
              @click="handleAddToCart"
            >
              <ShoppingCart :size="18" />
              + Keranjang
            </BaseButton>
            <BaseButton
              variant="primary"
              size="lg"
              class="flex-1"
              :disabled="product.stock === 0"
              @click="handleBuyNow"
            >
              <Zap :size="18" />
              Beli Sekarang
            </BaseButton>
          </div>

          <!-- Description -->
          <div class="border-t border-neutral-100 pt-4">
            <h2 class="font-semibold text-neutral-700 mb-2">Deskripsi Produk</h2>
            <p class="text-sm text-neutral-600 leading-relaxed whitespace-pre-line">
              {{ product.description || 'Tidak ada deskripsi.' }}
            </p>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ChevronRight, Minus, Plus, ShoppingCart, Zap, ImageOff } from 'lucide-vue-next'
import { useProducts }  from '@/composables/useProducts'
import { useCart }      from '@/composables/useCart'
import BaseButton       from '@/components/ui/BaseButton.vue'
import EmptyState       from '@/components/ui/EmptyState.vue'

const route   = useRoute()
const router  = useRouter()

const { currentProduct: product, loading, fetchById } = useProducts()
const { addToCart } = useCart()

const qty        = ref(1)
const imgError   = ref(false)
const cartLoading = ref(false)

async function handleAddToCart() {
  cartLoading.value = true
  try {
    await addToCart(product.value.id, qty.value)
  } finally {
    cartLoading.value = false
  }
}

async function handleBuyNow() {
  cartLoading.value = true
  try {
    await addToCart(product.value.id, qty.value)
    router.push('/cart')
  } finally {
    cartLoading.value = false
  }
}

function formatPrice(price) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency', currency: 'IDR', maximumFractionDigits: 0,
  }).format(price)
}

onMounted(() => fetchById(route.params.id))
</script>
