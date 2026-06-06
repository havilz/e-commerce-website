import api from './api'

const orderService = {
  /**
   * Checkout — buat order dari cart
   * @param {string} address
   * @returns {Promise}
   */
  async checkout(address) {
    const { data } = await api.post('/orders/checkout', { address })
    return data
  },

  /**
   * Get riwayat order user
   * @returns {Promise}
   */
  async getOrders() {
    const { data } = await api.get('/orders')
    return data
  },

  /**
   * Get detail 1 order
   * @param {number|string} id
   * @returns {Promise}
   */
  async getById(id) {
    const { data } = await api.get(`/orders/${id}`)
    return data
  },
}

export default orderService
