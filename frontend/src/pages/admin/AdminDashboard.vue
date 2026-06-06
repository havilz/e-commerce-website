<template>
  <div class="flex h-screen overflow-hidden bg-neutral-50">
    <AdminSidebar />

    <main class="flex-1 p-8 overflow-y-auto">
      <!-- Header -->
      <div class="flex items-center justify-between mb-8">
        <div>
          <h1 class="font-heading font-bold text-2xl text-neutral-800">Dashboard</h1>
          <p class="text-sm text-neutral-400 mt-0.5">Selamat datang, {{ authStore.user?.name }} 👋</p>
        </div>
        <span class="text-xs text-neutral-400">{{ today }}</span>
      </div>

      <!-- Stat cards -->
      <div class="grid grid-cols-1 sm:grid-cols-3 gap-5 mb-8">
        <div
          v-for="stat in stats"
          :key="stat.label"
          class="card p-5 flex items-center gap-4"
        >
          <div :class="['w-12 h-12 rounded-xl flex items-center justify-center flex-shrink-0', stat.bg]">
            <component :is="stat.icon" :size="22" :class="stat.color" />
          </div>
          <div>
            <p class="text-2xl font-extrabold text-neutral-800">
              <span v-if="loadingStats" class="inline-block w-12 h-6 bg-neutral-200 rounded animate-pulse" />
              <span v-else>{{ stat.value }}</span>
            </p>
            <p class="text-xs text-neutral-400 mt-0.5">{{ stat.label }}</p>
          </div>
        </div>
      </div>

      <!-- Recent orders -->
      <div class="card p-5">
        <div class="flex items-center justify-between mb-4">
          <h2 class="font-heading font-bold text-lg text-neutral-800">Pesanan Terbaru</h2>
          <RouterLink to="/admin/orders" class="text-sm text-primary-600 hover:underline">Lihat Semua →</RouterLink>
        </div>

        <!-- Loading -->
        <div v-if="loadingOrders" class="space-y-3">
          <div v-for="n in 5" :key="n" class="flex gap-4 animate-pulse py-2">
            <div class="h-4 bg-neutral-200 rounded w-16" />
            <div class="h-4 bg-neutral-200 rounded flex-1" />
            <div class="h-4 bg-neutral-200 rounded w-24" />
            <div class="h-4 bg-neutral-200 rounded w-20" />
          </div>
        </div>

        <!-- Empty -->
        <p v-else-if="!recentOrders.length" class="text-sm text-neutral-400 text-center py-6">Belum ada pesanan.</p>

        <!-- Table -->
        <div v-else class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead>
              <tr class="border-b border-neutral-100">
                <th class="text-left py-2 px-2 text-xs text-neutral-400 font-medium">#ID</th>
                <th class="text-left py-2 px-2 text-xs text-neutral-400 font-medium">Pembeli</th>
                <th class="text-left py-2 px-2 text-xs text-neutral-400 font-medium">Total</th>
                <th class="text-left py-2 px-2 text-xs text-neutral-400 font-medium">Status</th>
                <th class="text-left py-2 px-2 text-xs text-neutral-400 font-medium">Tanggal</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="order in recentOrders"
                :key="order.id"
                class="border-b border-neutral-50 hover:bg-neutral-50 transition-colors"
              >
                <td class="py-2.5 px-2 font-medium text-neutral-700">#{{ order.id }}</td>
                <td class="py-2.5 px-2 text-neutral-600">{{ order.user?.name || '—' }}</td>
                <td class="py-2.5 px-2 font-semibold text-neutral-800">{{ formatPrice(order.total_price) }}</td>
                <td class="py-2.5 px-2">
                  <BaseBadge :variant="statusVariant(order.status)" pill>{{ statusLabel(order.status) }}</BaseBadge>
                </td>
                <td class="py-2.5 px-2 text-neutral-400 text-xs">{{ formatDate(order.created_at) }}</td>
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
import { Package, ClipboardList, TrendingUp } from 'lucide-vue-next'
import { useAuthStore }  from '@/stores/auth.store'
import AdminSidebar      from '@/components/layout/AdminSidebar.vue'
import BaseBadge         from '@/components/ui/BaseBadge.vue'
import adminService      from '@/services/admin.service'
import productService    from '@/services/product.service'

const authStore     = useAuthStore()
const recentOrders  = ref([])
const loadingOrders = ref(false)
const loadingStats  = ref(false)

const totalProducts = ref(0)
const totalOrders   = ref(0)
const totalRevenue  = ref(0)

const today = new Date().toLocaleDateString('id-ID', { weekday: 'long', day: 'numeric', month: 'long', year: 'numeric' })

const stats = computed(() => [
  {
    label: 'Total Produk',
    value: totalProducts.value,
    icon:  Package,
    bg:    'bg-primary-50',
    color: 'text-primary-600',
  },
  {
    label: 'Total Pesanan',
    value: totalOrders.value,
    icon:  ClipboardList,
    bg:    'bg-secondary-500/10',
    color: 'text-secondary-600',
  },
  {
    label: 'Total Pendapatan',
    value: formatPrice(totalRevenue.value),
    icon:  TrendingUp,
    bg:    'bg-accent-500/10',
    color: 'text-accent-600',
  },
])

async function loadData() {
  loadingOrders.value = true
  loadingStats.value  = true
  try {
    const [ordersRes, productsRes] = await Promise.all([
      adminService.getAllOrders({ limit: 100 }),
      productService.getAll({ limit: 1000 }),
    ])
    // getAllOrders returns pagination: { data: [...], meta: { total, ... } }
    const orders = ordersRes.data || []
    recentOrders.value  = orders.slice(0, 10)
    totalOrders.value   = ordersRes.meta?.total ?? orders.length
    totalRevenue.value  = orders.reduce((sum, o) => sum + (o.total_price || 0), 0)
    // getAll returns pagination: { data: [...], meta: { total, ... } }
    totalProducts.value = productsRes.meta?.total ?? (productsRes.data?.length ?? 0)
  } finally {
    loadingOrders.value = false
    loadingStats.value  = false
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

onMounted(loadData)
</script>
