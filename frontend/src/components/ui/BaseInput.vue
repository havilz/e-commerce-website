<template>
  <div class="flex flex-col gap-1.5">
    <label
      v-if="label"
      :for="inputId"
      class="text-sm font-medium text-neutral-700"
    >
      {{ label }}
      <span v-if="required" class="text-red-500 ml-0.5">*</span>
    </label>

    <div class="relative">
      <!-- Prefix icon -->
      <div
        v-if="$slots.prefix"
        class="absolute left-3 top-1/2 -translate-y-1/2 text-neutral-400"
      >
        <slot name="prefix" />
      </div>

      <input
        :id="inputId"
        v-bind="$attrs"
        :type="inputType"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :required="required"
        :autocomplete="autocomplete"
        @input="$emit('update:modelValue', $event.target.value)"
        :class="[
          'input-field',
          $slots.prefix ? 'pl-10' : '',
          (type === 'password' || $slots.suffix) ? 'pr-10' : '',
          error    ? 'border-red-400 focus:ring-red-400' : '',
          disabled ? 'bg-neutral-50 cursor-not-allowed' : '',
        ]"
      />

      <!-- Password toggle -->
      <button
        v-if="type === 'password'"
        type="button"
        @click="showPass = !showPass"
        class="absolute right-3 top-1/2 -translate-y-1/2 text-neutral-400 hover:text-neutral-600 transition-colors"
        :aria-label="showPass ? 'Sembunyikan password' : 'Tampilkan password'"
      >
        <component :is="showPass ? EyeOff : Eye" :size="18" />
      </button>

      <!-- Suffix slot -->
      <div
        v-else-if="$slots.suffix"
        class="absolute right-3 top-1/2 -translate-y-1/2 text-neutral-400"
      >
        <slot name="suffix" />
      </div>
    </div>

    <!-- Error message -->
    <p v-if="error" class="text-xs text-red-500 flex items-center gap-1">
      <AlertCircle :size="12" />
      {{ error }}
    </p>

    <!-- Helper text -->
    <p v-else-if="hint" class="text-xs text-neutral-400">{{ hint }}</p>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Eye, EyeOff, AlertCircle } from 'lucide-vue-next'

defineOptions({ inheritAttrs: false })

const props = defineProps({
  modelValue:   { type: String,  default: '' },
  label:        { type: String,  default: '' },
  type:         { type: String,  default: 'text' },
  placeholder:  { type: String,  default: '' },
  error:        { type: String,  default: '' },
  hint:         { type: String,  default: '' },
  required:     { type: Boolean, default: false },
  disabled:     { type: Boolean, default: false },
  autocomplete: { type: String,  default: 'off' },
  id:           { type: String,  default: '' },
})

defineEmits(['update:modelValue'])

const showPass = ref(false)

const inputId   = computed(() => props.id || `input-${Math.random().toString(36).slice(2, 9)}`)
const inputType = computed(() => {
  if (props.type === 'password') return showPass.value ? 'text' : 'password'
  return props.type
})
</script>
