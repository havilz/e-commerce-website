<template>
  <span :class="['inline-flex items-center gap-1 font-bold rounded text-xs px-2 py-0.5', variantClasses, { 'rounded-full': pill }]">
    <slot />
  </span>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  variant: {
    type: String,
    default: 'default',
    validator: (v) => ['default', 'discount', 'flash', 'new', 'success', 'warning', 'error', 'info'].includes(v),
  },
  pill: { type: Boolean, default: false },
})

const variantClasses = computed(() => ({
  default:  'bg-neutral-100 text-neutral-700',
  discount: 'bg-secondary-500 text-white',
  flash:    'bg-red-500 text-white animate-pulse',
  new:      'bg-accent-500 text-white',
  success:  'bg-green-100 text-green-700',
  warning:  'bg-yellow-100 text-yellow-700',
  error:    'bg-red-100 text-red-700',
  info:     'bg-blue-100 text-blue-700',
}[props.variant]))
</script>
