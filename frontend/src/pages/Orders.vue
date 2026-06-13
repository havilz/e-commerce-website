<template>
  <div class="container-page py-8">
    <h1 class="font-heading font-bold text-2xl text-neutral-800 mb-6">Pesanan Saya</h1>

    <!-- Loading -->
    <div v-if="loading" class="space-y-4">
      <div v-for="n in 4" :key="n" class="card p-5 animate-pulse flex gap-4">
        <div class="flex-1 space-y-2">
          <div class="h-4 bg-neutral-200 rounded w-1/4" />
          <div class="h-4 bg-neutral-200 rounded w-1/2" />
          <div class="h-4 bg-neutral-200 rounded w-1/3" />
        </div>
        <div class="h-8 w-24 bg-neutral-200 rounded-lg" />
      </div>
    </div>

    <!-- Empty -->
    <EmptyState
      v-else-if="!orders.length"
      title="Belum ada pesanan"
      description="Kamu belum pernah melakukan pembelian. Yuk mulai belanja!"
    >
      <RouterLink to="/products">
        <BaseButton variant="primary">Mulai Belanja</BaseButton>
      </RouterLink>
    </EmptyState>

    <!-- List pesanan -->
    <div v-else class="space-y-4">
      <div
        v-for="order in orders"
        :key="order.id"
        class="card p-5 hover:shadow-md transition-shadow duration-200"
      >
        <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-3">
          <div>
            <div class="flex items-center gap-3 mb-1">
              <span class="font-bold text-neutral-800">#{{ order.id }}</span>
              <BaseBadge :variant="statusVariant(order.status)" pill>{{ statusLabel(order.status) }}</BaseBadge>
            </div>
            <p class="text-xs text-neutral-400 mb-1">{{ formatDate(order.created_at) }}</p>
            <p class="text-sm text-neutral-600">
              {{ order.items?.length || 0 }} item · Alamat: {{ order.address }}
            </p>
          </div>
          <div class="flex items-center gap-3">
            <span class="font-bold text-lg text-primary-600">{{ formatPrice(order.total_price) }}</span>
            <RouterLink :to="`/orders/${order.id}`">
              <BaseButton variant="outline" size="sm">Detail</BaseButton>
            </RouterLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import orderService from '@/services/order.service'
import BaseBadge  from '@/components/ui/BaseBadge.vue'
import BaseButton from '@/components/ui/BaseButton.vue'
import EmptyState from '@/components/ui/EmptyState.vue'

const orders  = ref([])
const loading = ref(false)

async function fetchOrders() {
  loading.value = true
  try {
    const res = await orderService.getOrders()
    orders.value = res.data || []
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

onMounted(fetchOrders)
</script>
