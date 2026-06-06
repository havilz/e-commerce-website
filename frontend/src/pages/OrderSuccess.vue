<template>
  <div class="container-page py-16 text-center">
    <div class="max-w-md mx-auto animate-bounce-in">
      <!-- Icon sukses -->
      <div class="w-24 h-24 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-6">
        <CheckCircle2 :size="48" class="text-green-500" />
      </div>

      <h1 class="font-heading font-extrabold text-3xl text-neutral-800 mb-3">Pesanan Berhasil! 🎉</h1>
      <p class="text-neutral-500 mb-2">
        Terima kasih telah berbelanja di <strong class="text-primary-600">ShopKu</strong>.
      </p>
      <p v-if="orderId" class="text-sm text-neutral-400 mb-8">
        ID Pesanan: <strong class="text-neutral-700">#{{ orderId }}</strong>
      </p>

      <div class="flex flex-col sm:flex-row gap-3 justify-center">
        <RouterLink to="/orders">
          <BaseButton variant="primary" size="lg">
            <Package :size="18" />
            Lihat Pesanan
          </BaseButton>
        </RouterLink>
        <RouterLink to="/products">
          <BaseButton variant="outline" size="lg">
            Lanjut Belanja
          </BaseButton>
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { CheckCircle2, Package } from 'lucide-vue-next'
import BaseButton from '@/components/ui/BaseButton.vue'

const route   = useRoute()
const orderId = computed(() => route.query.orderId)
</script>
