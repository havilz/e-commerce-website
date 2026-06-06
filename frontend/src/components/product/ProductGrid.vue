<template>
  <div>
    <!-- Loading skeletons -->
    <div
      v-if="loading"
      class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4"
      aria-label="Memuat produk..."
    >
      <ProductSkeleton v-for="n in skeletonCount" :key="n" />
    </div>

    <!-- Empty state -->
    <EmptyState
      v-else-if="!products.length"
      title="Produk tidak ditemukan"
      description="Coba kata kunci lain atau hapus filter yang aktif."
      :icon="SearchX"
    >
      <BaseButton variant="ghost" @click="$emit('clear-filter')">Hapus Filter</BaseButton>
    </EmptyState>

    <!-- Product grid -->
    <div
      v-else
      class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4"
    >
      <ProductCard
        v-for="product in products"
        :key="product.id"
        :product="product"
      />
    </div>
  </div>
</template>

<script setup>
import { SearchX } from 'lucide-vue-next'
import ProductCard    from './ProductCard.vue'
import ProductSkeleton from './ProductSkeleton.vue'
import EmptyState     from '@/components/ui/EmptyState.vue'
import BaseButton     from '@/components/ui/BaseButton.vue'

defineProps({
  products:      { type: Array,   default: () => [] },
  loading:       { type: Boolean, default: false },
  skeletonCount: { type: Number,  default: 8 },
})

defineEmits(['clear-filter'])
</script>
