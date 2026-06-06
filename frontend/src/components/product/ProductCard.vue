<template>
  <RouterLink
    :to="`/products/${product.id}`"
    class="bg-white rounded-xl shadow-sm hover:shadow-md transition-all duration-200 overflow-hidden
           border border-neutral-100 hover:-translate-y-1 group block"
    :aria-label="`Lihat detail ${product.name}`"
  >
    <!-- Image -->
    <div class="relative aspect-square overflow-hidden bg-neutral-100">
      <img
        v-if="product.image_url"
        :src="product.image_url"
        :alt="product.name"
        loading="lazy"
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
        @error="imgError = true"
      />
      <!-- Fallback image -->
      <div v-if="!product.image_url || imgError" class="w-full h-full flex items-center justify-center bg-neutral-100">
        <ImageOff :size="40" class="text-neutral-300" />
      </div>

      <!-- Badge diskon -->
      <ProductBadge
        v-if="discountPercent > 0"
        :discount="discountPercent"
        class="absolute top-2 left-2"
      />

      <!-- Stock habis overlay -->
      <div
        v-if="product.stock === 0"
        class="absolute inset-0 bg-black/40 flex items-center justify-center"
      >
        <span class="bg-white text-neutral-700 text-xs font-bold px-3 py-1 rounded-full">Habis</span>
      </div>
    </div>

    <!-- Info -->
    <div class="p-3">
      <!-- Kategori -->
      <p class="text-xs text-primary-600 font-medium mb-0.5 uppercase tracking-wide">{{ product.category }}</p>
      <!-- Nama -->
      <p class="text-sm text-neutral-700 font-medium line-clamp-2 mb-2 leading-snug">{{ product.name }}</p>
      <!-- Harga -->
      <p class="text-base font-bold text-neutral-800">{{ formatPrice(product.price) }}</p>
      <!-- Stock info -->
      <p class="text-xs text-neutral-400 mt-1">Stok: {{ product.stock }}</p>
    </div>
  </RouterLink>
</template>

<script setup>
import { ref, computed } from 'vue'
import { ImageOff } from 'lucide-vue-next'
import ProductBadge from './ProductBadge.vue'

const props = defineProps({
  product: {
    type: Object,
    required: true,
  },
})

const imgError = ref(false)

const discountPercent = computed(() => {
  // If product has original_price, calculate discount
  if (props.product.original_price && props.product.original_price > props.product.price) {
    return Math.round((1 - props.product.price / props.product.original_price) * 100)
  }
  return 0
})

function formatPrice(price) {
  return new Intl.NumberFormat('id-ID', {
    style:    'currency',
    currency: 'IDR',
    maximumFractionDigits: 0,
  }).format(price)
}
</script>
