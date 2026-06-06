<template>
  <div class="flex h-screen overflow-hidden bg-neutral-50">
    <AdminSidebar />

    <main class="flex-1 p-8 overflow-y-auto">
      <!-- Header -->
      <div class="mb-6">
        <h1 class="font-heading font-bold text-2xl text-neutral-800">Manajemen Pesanan</h1>
        <p class="text-sm text-neutral-400 mt-0.5">{{ filteredOrders.length }} pesanan</p>
      </div>

      <!-- Filters -->
      <div class="flex flex-col sm:flex-row gap-3 mb-5">
        <div class="relative max-w-sm flex-1">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Cari ID atau pembeli..."
            class="input-field pr-10 py-2.5 text-sm"
            aria-label="Cari pesanan"
          />
          <Search :size="18" class="absolute right-3 top-1/2 -translate-y-1/2 text-neutral-400" />
        </div>
        <select
          v-model="statusFilter"
          class="input-field py-2.5 text-sm max-w-xs cursor-pointer"
          aria-label="Filter status"
        >
          <option value="">Semua Status</option>
          <option value="pending">Menunggu</option>
          <option value="processing">Diproses</option>
          <option value="completed">Selesai</option>
          <option value="cancelled">Dibatalkan</option>
        </select>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="card p-5 space-y-3">
        <div v-for="n in 6" :key="n" class="flex gap-4 animate-pulse py-2.5">
          <div class="h-4 bg-neutral-200 rounded w-12" />
          <div class="h-4 bg-neutral-200 rounded flex-1" />
          <div class="h-4 bg-neutral-200 rounded w-24" />
          <div class="h-4 bg-neutral-200 rounded w-20" />
          <div class="h-8 bg-neutral-200 rounded w-28" />
        </div>
      </div>

      <!-- Empty -->
      <EmptyState
        v-else-if="!filteredOrders.length"
        title="Tidak ada pesanan"
        description="Belum ada pesanan yang masuk."
      />

      <!-- Table -->
      <div v-else class="card overflow-hidden">
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead class="bg-neutral-50 border-b border-neutral-100">
              <tr>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">#ID</th>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Pembeli</th>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Total</th>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Tanggal</th>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Status</th>
                <th class="text-right py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Update Status</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="order in filteredOrders"
                :key="order.id"
                class="border-b border-neutral-50 hover:bg-neutral-50/50 transition-colors"
              >
                <td class="py-3 px-4 font-medium text-neutral-700">#{{ order.id }}</td>
                <td class="py-3 px-4 text-neutral-600">{{ order.user?.name || '—' }}</td>
                <td class="py-3 px-4 font-semibold text-neutral-800">{{ formatPrice(order.total_price) }}</td>
                <td class="py-3 px-4 text-neutral-400 text-xs">{{ formatDate(order.created_at) }}</td>
                <td class="py-3 px-4">
                  <BaseBadge :variant="statusVariant(order.status)" pill>{{ statusLabel(order.status) }}</BaseBadge>
                </td>
                <td class="py-3 px-4">
                  <div class="flex items-center justify-end gap-2">
                    <select
                      :value="order.status"
                      @change="handleStatusChange(order, $event.target.value)"
                      :disabled="updatingId === order.id || order.status === 'completed' || order.status === 'cancelled'"
                      class="text-xs border border-neutral-200 rounded-lg px-3 py-1.5 bg-white
                             focus:outline-none focus:ring-2 focus:ring-primary-500
                             disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer"
                      :aria-label="`Update status pesanan #${order.id}`"
                    >
                      <option value="pending">Menunggu</option>
                      <option value="processing">Diproses</option>
                      <option value="completed">Selesai</option>
                      <option value="cancelled">Dibatalkan</option>
                    </select>
                    <!-- Loading spinner -->
                    <svg
                      v-if="updatingId === order.id"
                      class="animate-spin h-4 w-4 text-primary-500"
                      fill="none"
                      viewBox="0 0 24 24"
                      aria-hidden="true"
                    >
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
                    </svg>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Search } from 'lucide-vue-next'
import AdminSidebar from '@/components/layout/AdminSidebar.vue'
import BaseBadge    from '@/components/ui/BaseBadge.vue'
import EmptyState   from '@/components/ui/EmptyState.vue'
import adminService from '@/services/admin.service'
import { useToast } from '@/composables/useToast'

const toast       = useToast()
const orders      = ref([])
const loading     = ref(false)
const updatingId  = ref(null)
const searchQuery = ref('')
const statusFilter = ref('')

const filteredOrders = computed(() => {
  let list = orders.value
  if (statusFilter.value) {
    list = list.filter((o) => o.status === statusFilter.value)
  }
  if (searchQuery.value.trim()) {
    const q = searchQuery.value.toLowerCase()
    list = list.filter(
      (o) =>
        String(o.id).includes(q) ||
        (o.user?.name || '').toLowerCase().includes(q) ||
        (o.user?.email || '').toLowerCase().includes(q)
    )
  }
  return list
})

async function fetchOrders() {
  loading.value = true
  try {
    const res = await adminService.getAllOrders({ limit: 100 })
    // getAllOrders returns pagination: { data: [...], meta: { total, ... } }
    orders.value = res.data || []
  } finally {
    loading.value = false
  }
}

async function handleStatusChange(order, newStatus) {
  if (newStatus === order.status) return
  updatingId.value = order.id
  try {
    await adminService.updateOrderStatus(order.id, newStatus)
    order.status = newStatus
    toast.success(`Status pesanan #${order.id} diperbarui.`)
  } catch {
    toast.error('Gagal mengubah status pesanan.')
  } finally {
    updatingId.value = null
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
  return new Date(dateStr).toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })
}

onMounted(fetchOrders)
</script>
