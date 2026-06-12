<template>
  <div class="min-h-screen bg-gradient-to-br from-primary-50 via-white to-neutral-100 flex items-center justify-center p-4">
    <div class="w-full max-w-md">
      <!-- Logo -->
      <div class="text-center mb-8">
        <RouterLink to="/" class="inline-flex items-center gap-2">
          <span class="text-4xl">🛍️</span>
          <span class="font-heading font-extrabold text-3xl text-primary-600">ShopKu</span>
        </RouterLink>
        <p class="text-neutral-500 mt-2 text-sm">Buat akun baru</p>
      </div>

      <div class="card p-8 animate-fade-in">
        <h1 class="font-heading font-bold text-2xl text-neutral-800 mb-6">Daftar Akun 🎉</h1>

        <form @submit.prevent="handleRegister" class="space-y-4" novalidate>
          <BaseInput
            v-model="form.name"
            label="Nama Lengkap"
            type="text"
            placeholder="Nama kamu"
            :error="errors.name"
            required
            autocomplete="name"
          />

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
            placeholder="Min. 6 karakter"
            :error="errors.password"
            required
            autocomplete="new-password"
          />

          <BaseInput
            v-model="form.confirmPassword"
            label="Konfirmasi Password"
            type="password"
            placeholder="Ulangi password"
            :error="errors.confirmPassword"
            required
            autocomplete="new-password"
          />

          <div v-if="errorMsg" class="bg-red-50 border border-red-200 rounded-lg px-4 py-3 text-sm text-red-600 flex items-center gap-2">
            <AlertCircle :size="16" />
            {{ errorMsg }}
          </div>

          <BaseButton type="submit" variant="primary" size="lg" block :loading="loading">
            Buat Akun
          </BaseButton>
        </form>

        <p class="text-center text-sm text-neutral-500 mt-6">
          Sudah punya akun?
          <RouterLink to="/login" class="text-primary-600 font-semibold hover:underline">Masuk di sini</RouterLink>
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

const { register } = useAuth()

const loading  = ref(false)
const errorMsg = ref('')

const form = reactive({
  name: '', email: '', password: '', confirmPassword: '',
})
const errors = reactive({
  name: '', email: '', password: '', confirmPassword: '',
})

function validate() {
  Object.keys(errors).forEach((k) => (errors[k] = ''))
  let valid = true
  if (!form.name.trim()) {
    errors.name = 'Nama wajib diisi.'; valid = false
  }
  if (!form.email) {
    errors.email = 'Email wajib diisi.'; valid = false
  } else if (!/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(form.email)) {
    errors.email = 'Format email tidak valid.'; valid = false
  }
  if (!form.password) {
    errors.password = 'Password wajib diisi.'; valid = false
  } else if (form.password.length < 6) {
    errors.password = 'Password minimal 6 karakter.'; valid = false
  }
  if (form.password !== form.confirmPassword) {
    errors.confirmPassword = 'Password tidak cocok.'; valid = false
  }
  return valid
}

async function handleRegister() {
  if (!validate()) return
  loading.value  = true
  errorMsg.value = ''
  try {
    await register(form.name, form.email, form.password)
  } catch (err) {
    errorMsg.value = err.response?.data?.message || 'Registrasi gagal. Coba lagi.'
  } finally {
    loading.value = false
  }
}
</script>
