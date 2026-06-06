import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import authService from '@/services/auth.service'

export const useAuthStore = defineStore('auth', () => {
  // State
  const token = ref(localStorage.getItem('token') || null)
  const user  = ref(JSON.parse(localStorage.getItem('user') || 'null'))

  // Getters
  const isLoggedIn = computed(() => !!token.value)
  const isAdmin    = computed(() => user.value?.role === 'admin')

  // Actions
  async function login(email, password) {
    const res = await authService.login(email, password)
    // Backend returns: { success, message, data: { token, user } }
    const { token: t, user: u } = res.data
    token.value = t
    user.value  = u
    localStorage.setItem('token', t)
    localStorage.setItem('user', JSON.stringify(u))
    return res
  }

  async function register(name, email, password) {
    const res = await authService.register(name, email, password)
    return res
  }

  async function fetchMe() {
    try {
      const res = await authService.getMe()
      user.value = res.data
      localStorage.setItem('user', JSON.stringify(res.data))
    } catch {
      logout()
    }
  }

  function logout() {
    token.value = null
    user.value  = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  return { token, user, isLoggedIn, isAdmin, login, register, fetchMe, logout }
})
