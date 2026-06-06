import { useCartStore } from '@/stores/cart.store'
import { useAuthStore } from '@/stores/auth.store'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'

export function useCart() {
  const cartStore = useCartStore()
  const authStore = useAuthStore()
  const router    = useRouter()
  const toast     = useToast()

  async function addToCart(productId, quantity = 1) {
    if (!authStore.isLoggedIn) {
      toast.warning('Login dulu untuk menambahkan ke keranjang.')
      router.push('/login')
      return
    }
    try {
      await cartStore.addItem(productId, quantity)
      toast.success('Produk ditambahkan ke keranjang!')
    } catch (err) {
      const msg = err.response?.data?.message || 'Gagal menambahkan ke keranjang.'
      toast.error(msg)
    }
  }

  async function removeFromCart(cartItemId) {
    try {
      await cartStore.removeItem(cartItemId)
      toast.info('Item dihapus dari keranjang.')
    } catch {
      toast.error('Gagal menghapus item.')
    }
  }

  async function updateQuantity(cartItemId, quantity) {
    try {
      await cartStore.updateQty(cartItemId, quantity)
    } catch {
      toast.error('Gagal update quantity.')
    }
  }

  return { addToCart, removeFromCart, updateQuantity }
}
