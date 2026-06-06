<template>
  <button
    :type="type"
    :disabled="disabled || loading"
    :class="[
      'inline-flex items-center justify-center gap-2 font-semibold rounded-lg transition-all duration-200',
      'disabled:opacity-50 disabled:cursor-not-allowed focus:outline-none focus:ring-2 focus:ring-offset-2',
      variantClasses,
      sizeClasses,
      { 'w-full': block }
    ]"
    v-bind="$attrs"
  >
    <!-- Loading spinner -->
    <svg
      v-if="loading"
      class="animate-spin -ml-1 h-4 w-4"
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      aria-hidden="true"
    >
      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
    </svg>

    <slot />
  </button>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  variant: {
    type: String,
    default: 'primary',
    validator: (v) => ['primary', 'secondary', 'ghost', 'danger', 'outline'].includes(v),
  },
  size: {
    type: String,
    default: 'md',
    validator: (v) => ['sm', 'md', 'lg'].includes(v),
  },
  type:     { type: String,  default: 'button' },
  loading:  { type: Boolean, default: false },
  disabled: { type: Boolean, default: false },
  block:    { type: Boolean, default: false },
})

const variantClasses = computed(() => ({
  primary:   'bg-primary-500 hover:bg-primary-600 active:bg-primary-700 text-white shadow-sm hover:shadow-md focus:ring-primary-500',
  secondary: 'border-2 border-primary-500 text-primary-600 hover:bg-primary-50 focus:ring-primary-500',
  ghost:     'text-neutral-600 hover:text-neutral-800 hover:bg-neutral-100 focus:ring-neutral-400',
  danger:    'bg-red-500 hover:bg-red-600 active:bg-red-700 text-white focus:ring-red-500',
  outline:   'border border-neutral-200 text-neutral-700 hover:bg-neutral-50 focus:ring-neutral-400',
}[props.variant]))

const sizeClasses = computed(() => ({
  sm: 'px-3 py-1.5 text-sm',
  md: 'px-5 py-2.5 text-sm',
  lg: 'px-7 py-3 text-base',
}[props.size]))
</script>
