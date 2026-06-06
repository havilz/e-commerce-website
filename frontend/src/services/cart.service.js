import api from './api'

const cartService = {
  /**
   * Get isi cart user saat ini
   * @returns {Promise}
   */
  async getCart() {
    const { data } = await api.get('/cart')
    return data
  },

  /**
   * Tambah item ke cart
   * @param {number} productId
   * @param {number} quantity
   * @returns {Promise}
   */
  async addItem(productId, quantity = 1) {
    const { data } = await api.post('/cart', { product_id: productId, quantity })
    return data
  },

  /**
   * Update quantity item di cart
   * @param {number} cartItemId
   * @param {number} quantity
   * @returns {Promise}
   */
  async updateItem(cartItemId, quantity) {
    const { data } = await api.put(`/cart/${cartItemId}`, { quantity })
    return data
  },

  /**
   * Hapus 1 item dari cart
   * @param {number} cartItemId
   * @returns {Promise}
   */
  async removeItem(cartItemId) {
    const { data } = await api.delete(`/cart/${cartItemId}`)
    return data
  },

  /**
   * Clear seluruh cart
   * @returns {Promise}
   */
  async clearCart() {
    const { data } = await api.delete('/cart')
    return data
  },
}

export default cartService
