<template>
  <div class="card p-5 sticky top-24">
    <h3 class="font-heading font-bold text-lg text-neutral-800 mb-5">Ringkasan Belanja</h3>

    <div class="space-y-3 text-sm">
      <div class="flex justify-between text-neutral-600">
        <span>Subtotal ({{ totalItems }} item)</span>
        <span>{{ formatPrice(totalPrice) }}</span>
      </div>
      <div class="flex justify-between text-neutral-600">
        <span>Ongkos kirim</span>
        <span class="text-green-600 font-medium">GRATIS</span>
      </div>
      <div class="border-t border-neutral-100 pt-3 flex justify-between font-bold text-base text-neutral-800">
        <span>Total</span>
        <span class="text-primary-600">{{ formatPrice(totalPrice) }}</span>
      </div>
    </div>

    <BaseButton
      variant="primary"
      size="lg"
      block
      class="mt-6"
      :disabled="!totalItems"
      @click="$emit('checkout')"
    >
      <ShoppingBag :size="18" />
      Checkout Sekarang
    </BaseButton>

    <RouterLink to="/products" class="block text-center mt-3 text-sm text-neutral-500 hover:text-primary-600 transition-colors">
      ← Lanjut Belanja
    </RouterLink>
  </div>
</template>

<script setup>
import { ShoppingBag } from 'lucide-vue-next'
import BaseButton from '@/components/ui/BaseButton.vue'

defineProps({
  totalItems: { type: Number, default: 0 },
  totalPrice: { type: Number, default: 0 },
})

defineEmits(['checkout'])

function formatPrice(price) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency', currency: 'IDR', maximumFractionDigits: 0,
  }).format(price)
}
</script>
