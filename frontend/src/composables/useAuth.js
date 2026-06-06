import { useAuthStore } from '@/stores/auth.store'
import { useCartStore } from '@/stores/cart.store'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'

export function useAuth() {
  const authStore = useAuthStore()
  const cartStore = useCartStore()
  const router    = useRouter()
  const toast     = useToast()

  async function login(email, password) {
    await authStore.login(email, password)
    await cartStore.fetchCart()
    toast.success('Selamat datang kembali!')
    router.push('/')
  }

  async function register(name, email, password) {
    await authStore.register(name, email, password)
    toast.success('Registrasi berhasil! Silakan login.')
    router.push('/login')
  }

  function logout() {
    authStore.logout()
    cartStore.items = []
    toast.info('Kamu telah logout.')
    router.push('/')
  }

  return { login, register, logout }
}
