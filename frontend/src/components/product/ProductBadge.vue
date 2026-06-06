<template>
  <span
    v-if="label"
    :class="[
      'text-xs font-bold px-2 py-0.5 rounded text-white',
      badgeClass,
    ]"
  >
    {{ label }}
  </span>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  type:     { type: String, default: '' }, // 'discount' | 'new' | 'flash'
  discount: { type: Number, default: 0 },
})

const label = computed(() => {
  if (props.type === 'flash')    return '⚡ FLASH SALE'
  if (props.type === 'new')      return 'NEW'
  if (props.discount > 0)        return `-${props.discount}%`
  return ''
})

const badgeClass = computed(() => {
  if (props.type === 'flash')    return 'bg-red-500 animate-pulse'
  if (props.type === 'new')      return 'bg-accent-500'
  if (props.discount > 0)        return 'bg-secondary-500'
  return ''
})
</script>
