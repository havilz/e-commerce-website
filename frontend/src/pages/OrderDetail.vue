<template>
  <div class="container-page py-8">
    <!-- Back -->
    <RouterLink to="/orders" class="inline-flex items-center gap-2 text-sm text-neutral-500 hover:text-primary-600 transition-colors mb-6">
      <ChevronLeft :size="18" /> Kembali ke Pesanan
    </RouterLink>

    <!-- Loading -->
    <div v-if="loading" class="animate-pulse space-y-4">
      <div class="h-8 bg-neutral-200 rounded w-48" />
      <div class="card p-5 space-y-3">
        <div class="h-4 bg-neutral-200 rounded w-1/4" />
        <div class="h-4 bg-neutral-200 rounded w-1/2" />
        <div class="h-4 bg-neutral-200 rounded w-1/3" />
      </div>
    </div>

    <!-- Not found -->
    <EmptyState
      v-else-if="!order"
      title="Pesanan tidak ditemukan"
      description="Pesanan ini mungkin sudah dihapus atau tidak ada."
    >
      <RouterLink to="/orders"><BaseButton variant="primary">Lihat Semua Pesanan</BaseButton></RouterLink>
    </EmptyState>

    <template v-else>
      <!-- Header -->
      <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-3 mb-6">
        <div>
          <h1 class="font-heading font-bold text-2xl text-neutral-800">Pesanan #{{ order.id }}</h1>
          <p class="text-sm text-neutral-400 mt-1">{{ formatDate(order.created_at) }}</p>
        </div>
        <BaseBadge :variant="statusVariant(order.status)" pill class="text-sm px-4 py-1.5">
          {{ statusLabel(order.status) }}
        </BaseBadge>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Order items -->
        <div class="lg:col-span-2 card p-5">
          <h2 class="font-semibold text-neutral-700 mb-4">Item Pesanan</h2>
          <div class="space-y-4">
            <div
              v-for="item in order.items"
              :key="item.id"
              class="flex items-center gap-4 py-3 border-b border-neutral-100 last:border-0"
            >
              <img
                v-if="item.product?.image_url"
                :src="item.product.image_url"
                :alt="item.product?.name"
                class="w-16 h-16 object-cover rounded-lg flex-shrink-0"
              />
              <div v-else class="w-16 h-16 bg-neutral-100 rounded-lg flex-shrink-0 flex items-center justify-center">
                <ImageOff :size="20" class="text-neutral-300" />
              </div>
              <div class="flex-1 min-w-0">
                <p class="text-sm font-medium text-neutral-800 truncate">{{ item.product?.name }}</p>
                <p class="text-xs text-neutral-400 mt-0.5">{{ formatPrice(item.price_at_purchase) }} × {{ item.quantity }}</p>
              </div>
              <p class="font-bold text-neutral-800 flex-shrink-0">
                {{ formatPrice(item.price_at_purchase * item.quantity) }}
              </p>
            </div>
          </div>
        </div>

        <!-- Order info -->
        <div class="space-y-4">
          <div class="card p-5">
            <h2 class="font-semibold text-neutral-700 mb-3">Informasi Pesanan</h2>
            <dl class="space-y-2 text-sm">
              <div class="flex justify-between">
                <dt class="text-neutral-500">Status</dt>
                <dd class="font-medium text-neutral-800">{{ statusLabel(order.status) }}</dd>
              </div>
              <div class="flex justify-between">
                <dt class="text-neutral-500">Tanggal</dt>
                <dd class="font-medium text-neutral-800">{{ formatDate(order.created_at) }}</dd>
              </div>
              <div>
                <dt class="text-neutral-500 mb-1">Alamat</dt>
                <dd class="font-medium text-neutral-800 text-xs leading-relaxed">{{ order.address }}</dd>
              </div>
            </dl>
          </div>

          <div class="card p-5">
            <h2 class="font-semibold text-neutral-700 mb-3">Ringkasan Biaya</h2>
            <div class="space-y-2 text-sm">
              <div class="flex justify-between text-neutral-600">
                <span>Subtotal</span>
                <span>{{ formatPrice(order.total_price) }}</span>
              </div>
              <div class="flex justify-between text-neutral-600">
                <span>Ongkir</span>
                <span class="text-green-600 font-medium">GRATIS</span>
              </div>
              <div class="flex justify-between font-bold text-base text-neutral-800 pt-2 border-t border-neutral-100">
                <span>Total</span>
                <span class="text-primary-600">{{ formatPrice(order.total_price) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { ChevronLeft, ImageOff } from 'lucide-vue-next'
import orderService from '@/services/order.service'
import BaseBadge    from '@/components/ui/BaseBadge.vue'
import BaseButton   from '@/components/ui/BaseButton.vue'
import EmptyState   from '@/components/ui/EmptyState.vue'

const route   = useRoute()
const order   = ref(null)
const loading = ref(false)

async function fetchOrder() {
  loading.value = true
  try {
    const res = await orderService.getById(route.params.id)
    order.value = res.data
  } finally {
    loading.value = false
  }
}

function statusVariant(status) {
  return { pending: 'warning', processing: 'info', completed: 'success', cancelled: 'error' }[status] || 'default'
}

function statusLabel(status) {
  return { pending: 'Menunggu', processing: 'Diproses', completed: 'Selesai', cancelled: 'Dibatalkan' }[status] || status
}

function formatPrice(price) {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', maximumFractionDigits: 0 }).format(price)
}

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric', hour: '2-digit', minute: '2-digit' })
}

onMounted(fetchOrder)
</script>
