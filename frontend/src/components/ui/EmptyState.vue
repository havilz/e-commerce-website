<template>
  <div class="flex flex-col items-center justify-center py-16 px-4 text-center">
    <!-- Ilustrasi / icon -->
    <div class="mb-5">
      <img
        v-if="image"
        :src="image"
        :alt="title"
        class="w-40 h-40 object-contain opacity-80"
      />
      <component
        v-else-if="icon"
        :is="icon"
        :size="64"
        class="text-neutral-300"
      />
      <div
        v-else
        class="w-20 h-20 bg-neutral-100 rounded-full flex items-center justify-center"
      >
        <PackageOpen :size="36" class="text-neutral-300" />
      </div>
    </div>

    <h3 class="font-heading font-bold text-lg text-neutral-700 mb-2">{{ title }}</h3>
    <p v-if="description" class="text-sm text-neutral-400 max-w-xs mb-6">{{ description }}</p>

    <slot />
  </div>
</template>

<script setup>
import { PackageOpen } from 'lucide-vue-next'

defineProps({
  title:       { type: String, default: 'Tidak ada data' },
  description: { type: String, default: '' },
  image:       { type: String, default: '' },
  icon:        { type: Object, default: null },
})
</script>
