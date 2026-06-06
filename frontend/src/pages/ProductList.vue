<template>
  <div class="container-page py-8">
    <!-- Header -->
    <div class="mb-6">
      <h1 class="font-heading font-bold text-2xl text-neutral-800 mb-1">Semua Produk</h1>
      <p class="text-sm text-neutral-500">
        {{ pagination.total }} produk ditemukan
        <span v-if="filters.search"> untuk "<strong>{{ filters.search }}</strong>"</span>
      </p>
    </div>

    <!-- Filter bar -->
    <div class="flex flex-col sm:flex-row gap-3 mb-6">
      <!-- Search -->
      <div class="relative flex-1 max-w-sm">
        <input
          v-model="localSearch"
          @keyup.enter="applySearch"
          type="text"
          placeholder="Cari produk..."
          class="input-field pr-10 py-2.5 text-sm"
          :aria-label="'Cari produk'"
        />
        <button @click="applySearch" class="absolute right-3 top-1/2 -translate-y-1/2 text-neutral-400 hover:text-primary-500" aria-label="Cari">
          <Search :size="18" />
        </button>
      </div>

      <!-- Category filter -->
      <select
        v-model="localCategory"
        @change="applyCategory"
        class="input-field py-2.5 text-sm max-w-xs cursor-pointer"
        aria-label="Filter kategori"
      >
        <option value="">Semua Kategori</option>
        <option v-for="cat in categories" :key="cat" :value="cat">{{ cat }}</option>
      </select>

      <!-- Clear filter -->
      <BaseButton
        v-if="filters.search || filters.category"
        variant="ghost"
        size="md"
        @click="clearFilters"
      >
        <X :size="16" /> Hapus Filter
      </BaseButton>
    </div>

    <!-- Grid -->
    <ProductGrid
      :products="products"
      :loading="loading"
      @clear-filter="clearFilters"
    />

    <!-- Pagination -->
    <div v-if="totalPages > 1" class="flex justify-center items-center gap-2 mt-8">
      <button
        @click="changePage(pagination.page - 1)"
        :disabled="pagination.page <= 1"
        class="btn-ghost px-3 py-2 rounded-lg text-sm disabled:opacity-40 disabled:cursor-not-allowed"
        aria-label="Halaman sebelumnya"
      >
        <ChevronLeft :size="18" />
      </button>

      <button
        v-for="p in displayPages"
        :key="p"
        @click="changePage(p)"
        :class="[
          'px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200',
          p === pagination.page
            ? 'bg-primary-500 text-white'
            : 'hover:bg-neutral-100 text-neutral-600'
        ]"
        :aria-current="p === pagination.page ? 'page' : undefined"
      >
        {{ p }}
      </button>

      <button
        @click="changePage(pagination.page + 1)"
        :disabled="pagination.page >= totalPages"
        class="btn-ghost px-3 py-2 rounded-lg text-sm disabled:opacity-40 disabled:cursor-not-allowed"
        aria-label="Halaman berikutnya"
      >
        <ChevronRight :size="18" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Search, X, ChevronLeft, ChevronRight } from 'lucide-vue-next'
import { useProducts } from '@/composables/useProducts'
import ProductGrid from '@/components/product/ProductGrid.vue'
import BaseButton  from '@/components/ui/BaseButton.vue'

const route  = useRoute()
const router = useRouter()

const { products, loading, filters, pagination, fetchAll, filterByCategory, changePage, search } = useProducts()

const categories = ['Electronics', 'Fashion', 'Lifestyle', 'Home', 'Sports', 'Books', 'Beauty', 'Gaming']
const localSearch  = ref('')
const localCategory = ref('')

const totalPages = computed(() => Math.ceil(pagination.total / pagination.limit) || 1)
const displayPages = computed(() => {
  const pages = []
  const start = Math.max(1, pagination.page - 2)
  const end   = Math.min(totalPages.value, pagination.page + 2)
  for (let i = start; i <= end; i++) pages.push(i)
  return pages
})

function applySearch() {
  router.replace({ query: { ...route.query, search: localSearch.value || undefined } })
  search(localSearch.value)
}

function applyCategory() {
  router.replace({ query: { ...route.query, category: localCategory.value || undefined } })
  filterByCategory(localCategory.value)
}

function clearFilters() {
  localSearch.value   = ''
  localCategory.value = ''
  router.replace({ query: {} })
  filterByCategory('')
  search('')
}

onMounted(() => {
  // Read query params from URL
  localSearch.value   = route.query.search   ? String(route.query.search)   : ''
  localCategory.value = route.query.category ? String(route.query.category) : ''
  fetchAll({
    search:   localSearch.value   || undefined,
    category: localCategory.value || undefined,
  })
})
</script>
