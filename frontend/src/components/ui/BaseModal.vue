<template>
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="modelValue"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        role="dialog"
        :aria-modal="modelValue"
        :aria-label="title"
      >
        <!-- Backdrop -->
        <div
          class="absolute inset-0 bg-black/50 backdrop-blur-sm"
          @click="$emit('update:modelValue', false)"
        />

        <!-- Modal panel -->
        <div
          class="relative bg-white rounded-2xl shadow-xl w-full max-w-md animate-bounce-in"
          :class="sizeClass"
        >
          <!-- Header -->
          <div class="flex items-center justify-between p-5 border-b border-neutral-100">
            <h3 class="font-heading font-bold text-lg text-neutral-800">{{ title }}</h3>
            <button
              @click="$emit('update:modelValue', false)"
              class="p-1.5 rounded-full hover:bg-neutral-100 transition-colors text-neutral-400 hover:text-neutral-700"
              aria-label="Tutup modal"
            >
              <X :size="18" />
            </button>
          </div>

          <!-- Body -->
          <div class="p-5">
            <slot />
          </div>

          <!-- Footer -->
          <div v-if="$slots.footer" class="px-5 pb-5">
            <slot name="footer" />
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { computed } from 'vue'
import { X } from 'lucide-vue-next'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  title:      { type: String,  default: '' },
  size:       { type: String,  default: 'md', validator: (v) => ['sm', 'md', 'lg'].includes(v) },
})

defineEmits(['update:modelValue'])

const sizeClass = computed(() => ({
  sm: 'max-w-sm',
  md: 'max-w-md',
  lg: 'max-w-2xl',
}[props.size]))
</script>

<style scoped>
.modal-enter-active,
.modal-leave-active { transition: opacity 0.2s; }
.modal-enter-from,
.modal-leave-to { opacity: 0; }
</style>
