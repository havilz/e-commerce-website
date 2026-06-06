<template>
  <div>
    <!-- Hero Banner -->
    <section class="bg-gradient-to-r from-orange-500 to-red-500 text-white py-12 px-4">
      <div class="container-page">
        <div class="max-w-xl">
          <div class="inline-flex items-center gap-2 bg-white/20 rounded-full px-4 py-1 text-sm font-medium mb-4">
            ⚡ FLASH SALE HARI INI
          </div>
          <h1 class="font-heading font-extrabold text-4xl md:text-5xl mb-3 leading-tight">
            Diskon s/d 70%!
          </h1>
          <p class="text-white/80 text-lg mb-6">Temukan produk terbaik dengan harga yang nggak bikin kantong kering.</p>
          <RouterLink to="/products">
            <BaseButton size="lg" class="!bg-white !text-orange-600 hover:!bg-orange-50 font-bold shadow-lg">
              Belanja Sekarang →
            </BaseButton>
          </RouterLink>
        </div>
      </div>
    </section>

    <!-- Category pills -->
    <section class="container-page py-6">
      <h2 class="font-heading font-bold text-lg text-neutral-800 mb-3">Kategori</h2>
      <div class="flex flex-wrap gap-2">
        <button
          v-for="cat in categories"
          :key="cat.value"
          @click="filterByCategory(cat.value)"
          :class="[
            'px-4 py-2 rounded-full text-sm font-medium transition-all duration-200 border',
            activeCategory === cat.value
              ? 'bg-primary-500 text-white border-primary-500'
              : 'bg-white text-neutral-600 border-neutral-200 hover:border-primary-300 hover:text-primary-600'
          ]"
        >
          {{ cat.emoji }} {{ cat.label }}
        </button>
      </div>
    </section>

    <!-- Products -->
    <section class="container-page pb-12">
      <div class="flex items-center justify-between mb-4">
        <h2 class="font-heading font-bold text-xl text-neutral-800">
          {{ activeCategory ? activeCategory : 'Produk Terbaru' }}
        </h2>
        <RouterLink to="/products" class="text-sm text-primary-600 font-medium hover:underline flex items-center gap-1">
          Lihat Semua <ChevronRight :size="16" />
        </RouterLink>
      </div>

      <ProductGrid
        :products="products"
        :loading="loading"
        @clear-filter="clearFilter"
      />
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ChevronRight } from 'lucide-vue-next'
import { useProducts } from '@/composables/useProducts'
import ProductGrid from '@/components/product/ProductGrid.vue'
import BaseButton  from '@/components/ui/BaseButton.vue'

const { products, loading, fetchAll, filterByCategory: filterByCat, filters } = useProducts()
const activeCategory = ref('')

const categories = [
  { value: '',           emoji: '🛒', label: 'Semua' },
  { value: 'Electronics',emoji: '📱', label: 'Electronics' },
  { value: 'Fashion',    emoji: '👗', label: 'Fashion' },
  { value: 'Lifestyle',  emoji: '✨', label: 'Lifestyle' },
  { value: 'Home',       emoji: '🏠', label: 'Home' },
  { value: 'Sports',     emoji: '⚽', label: 'Sports' },
  { value: 'Books',      emoji: '📚', label: 'Books' },
  { value: 'Beauty',     emoji: '💄', label: 'Beauty' },
  { value: 'Gaming',     emoji: '🎮', label: 'Gaming' },
]

function filterByCategory(val) {
  activeCategory.value = val
  filterByCat(val)
  fetchAll({ limit: 8 })
}

function clearFilter() {
  activeCategory.value = ''
  filterByCat('')
}

onMounted(() => fetchAll({ limit: 8 }))
</script>
