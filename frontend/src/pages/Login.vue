<template>
  <div class="min-h-screen bg-gradient-to-br from-primary-50 via-white to-neutral-100 flex items-center justify-center p-4">
    <div class="w-full max-w-md">
      <!-- Logo -->
      <div class="text-center mb-8">
        <RouterLink to="/" class="inline-flex items-center gap-2">
          <span class="text-4xl">🛍️</span>
          <span class="font-heading font-extrabold text-3xl text-primary-600">ShopKu</span>
        </RouterLink>
        <p class="text-neutral-500 mt-2 text-sm">Masuk ke akun kamu</p>
      </div>

      <!-- Card -->
      <div class="card p-8 animate-fade-in">
        <h1 class="font-heading font-bold text-2xl text-neutral-800 mb-6">Selamat Datang 👋</h1>

        <form @submit.prevent="handleLogin" class="space-y-4" novalidate>
          <BaseInput
            v-model="form.email"
            label="Email"
            type="email"
            placeholder="nama@email.com"
            :error="errors.email"
            required
            autocomplete="email"
          />

          <BaseInput
            v-model="form.password"
            label="Password"
            type="password"
            placeholder="Masukkan password"
            :error="errors.password"
            required
            autocomplete="current-password"
          />

          <!-- Error global -->
          <div v-if="errorMsg" class="bg-red-50 border border-red-200 rounded-lg px-4 py-3 text-sm text-red-600 flex items-center gap-2">
            <AlertCircle :size="16" />
            {{ errorMsg }}
          </div>

          <BaseButton type="submit" variant="primary" size="lg" block :loading="loading">
            Masuk
          </BaseButton>
        </form>

        <p class="text-center text-sm text-neutral-500 mt-6">
          Belum punya akun?
          <RouterLink to="/register" class="text-primary-600 font-semibold hover:underline">Daftar sekarang</RouterLink>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { AlertCircle } from 'lucide-vue-next'
import BaseInput  from '@/components/ui/BaseInput.vue'
import BaseButton from '@/components/ui/BaseButton.vue'
import { useAuth } from '@/composables/useAuth'

const { login } = useAuth()

const loading  = ref(false)
const errorMsg = ref('')

const form = reactive({ email: '', password: '' })
const errors = reactive({ email: '', password: '' })

function validate() {
  errors.email    = ''
  errors.password = ''
  let valid = true
  if (!form.email) {
    errors.email = 'Email wajib diisi.'
    valid = false
  } else if (!/\S+@\S+\.\S+/.test(form.email)) {
    errors.email = 'Format email tidak valid.'
    valid = false
  }
  if (!form.password) {
    errors.password = 'Password wajib diisi.'
    valid = false
  }
  return valid
}

async function handleLogin() {
  if (!validate()) return
  loading.value  = true
  errorMsg.value = ''
  try {
    await login(form.email, form.password)
  } catch (err) {
    errorMsg.value = err.response?.data?.message || 'Email atau password salah.'
  } finally {
    loading.value = false
  }
}
</script>
