import api from './api'

const adminService = {
  // Product CRUD

  /**
   * Buat produk baru
   * @param {object} payload - { name, description, price, stock, image_url, category }
   * @returns {Promise}
   */
  async createProduct(payload) {
    const { data } = await api.post('/admin/products', payload)
    return data
  },

  /**
   * Update produk
   * @param {number} id
   * @param {object} payload
   * @returns {Promise}
   */
  async updateProduct(id, payload) {
    const { data } = await api.put(`/admin/products/${id}`, payload)
    return data
  },

  /**
   * Hapus produk
   * @param {number} id
   * @returns {Promise}
   */
  async deleteProduct(id) {
    const { data } = await api.delete(`/admin/products/${id}`)
    return data
  },

  // Orders

  /**
   * Get semua order (admin)
   * @param {object} params - { page, limit }
   * @returns {Promise}
   */
  async getAllOrders(params = {}) {
    const { data } = await api.get('/admin/orders', { params })
    return data
  },

  /**
   * Update status order
   * @param {number} id
   * @param {string} status - 'pending' | 'processing' | 'completed' | 'cancelled'
   * @returns {Promise}
   */
  async updateOrderStatus(id, status) {
    const { data } = await api.patch(`/admin/orders/${id}`, { status })
    return data
  },
}

export default adminService
