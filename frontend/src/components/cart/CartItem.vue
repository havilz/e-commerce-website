<template>
  <div class="flex items-center gap-4 py-4 border-b border-neutral-100 last:border-0">
    <!-- Product image -->
    <RouterLink :to="`/products/${item.product_id}`" class="flex-shrink-0">
      <img
        v-if="item.product?.image_url"
        :src="item.product.image_url"
        :alt="item.product?.name"
        class="w-20 h-20 object-cover rounded-lg border border-neutral-100"
        @error="imgError = true"
      />
      <div v-else class="w-20 h-20 bg-neutral-100 rounded-lg flex items-center justify-center">
        <ImageOff :size="24" class="text-neutral-300" />
      </div>
    </RouterLink>

    <!-- Product info -->
    <div class="flex-1 min-w-0">
      <RouterLink
        :to="`/products/${item.product_id}`"
        class="text-sm font-medium text-neutral-800 hover:text-primary-600 transition-colors line-clamp-2 block mb-1"
      >
        {{ item.product?.name }}
      </RouterLink>
      <p class="text-base font-bold text-neutral-800">{{ formatPrice(item.product?.price) }}</p>
      <p class="text-xs text-neutral-400 mt-0.5">Stok tersedia: {{ item.product?.stock }}</p>
    </div>

    <!-- Quantity controls -->
    <div class="flex items-center gap-2 flex-shrink-0">
      <button
        @click="$emit('update', item.id, item.quantity - 1)"
        :disabled="item.quantity <= 1 || loading"
        class="w-8 h-8 rounded-full border border-neutral-200 flex items-center justify-center
               hover:bg-neutral-50 transition-colors disabled:opacity-40 disabled:cursor-not-allowed"
        aria-label="Kurangi jumlah"
      >
        <Minus :size="14" />
      </button>

      <span class="w-8 text-center text-sm font-semibold">{{ item.quantity }}</span>

      <button
        @click="$emit('update', item.id, item.quantity + 1)"
        :disabled="item.quantity >= item.product?.stock || loading"
        class="w-8 h-8 rounded-full border border-neutral-200 flex items-center justify-center
               hover:bg-neutral-50 transition-colors disabled:opacity-40 disabled:cursor-not-allowed"
        aria-label="Tambah jumlah"
      >
        <Plus :size="14" />
      </button>
    </div>

    <!-- Subtotal + delete -->
    <div class="text-right flex-shrink-0">
      <p class="text-sm font-bold text-neutral-800">
        {{ formatPrice(item.product?.price * item.quantity) }}
      </p>
      <button
        @click="$emit('remove', item.id)"
        :disabled="loading"
        class="mt-1 text-red-400 hover:text-red-600 transition-colors disabled:opacity-40"
        aria-label="Hapus item"
      >
        <Trash2 :size="16" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Minus, Plus, Trash2, ImageOff } from 'lucide-vue-next'

defineProps({
  item:    { type: Object,  required: true },
  loading: { type: Boolean, default: false },
})

defineEmits(['update', 'remove'])

const imgError = ref(false)

function formatPrice(price = 0) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency', currency: 'IDR', maximumFractionDigits: 0,
  }).format(price)
}
</script>
