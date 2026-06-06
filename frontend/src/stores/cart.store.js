import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import cartService from '@/services/cart.service'

export const useCartStore = defineStore('cart', () => {
  // State
  const items   = ref([])
  const loading = ref(false)

  // Getters
  const totalItems = computed(() =>
    items.value.reduce((sum, item) => sum + item.quantity, 0)
  )

  const totalPrice = computed(() =>
    items.value.reduce((sum, item) => sum + item.product.price * item.quantity, 0)
  )

  // Actions
  async function fetchCart() {
    loading.value = true
    try {
      const res = await cartService.getCart()
      // Backend returns { items: [], total_items: n }
      items.value = res.data?.items || []
    } catch (err) {
      items.value = []
    } finally {
      loading.value = false
    }
  }

  async function addItem(productId, quantity = 1) {
    const res = await cartService.addItem(productId, quantity)
    await fetchCart() // sync dengan server
    return res
  }

  async function updateQty(cartItemId, quantity) {
    if (quantity < 1) return removeItem(cartItemId)
    const res = await cartService.updateItem(cartItemId, quantity)
    await fetchCart()
    return res
  }

  async function removeItem(cartItemId) {
    const res = await cartService.removeItem(cartItemId)
    await fetchCart()
    return res
  }

  async function clearCart() {
    const res = await cartService.clearCart()
    items.value = []
    return res
  }

  return {
    items,
    loading,
    totalItems,
    totalPrice,
    fetchCart,
    addItem,
    updateQty,
    removeItem,
    clearCart,
  }
})
