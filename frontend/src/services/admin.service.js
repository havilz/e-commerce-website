import api from './api'

const adminService = {
  // Category CRUD

  /**
   * Buat kategori baru
   * @param {object} payload - { name }
   * @returns {Promise}
   */
  async createCategory(payload) {
    const { data } = await api.post('/admin/categories', payload)
    return data
  },

  /**
   * Update kategori
   * @param {number} id
   * @param {object} payload - { name }
   * @returns {Promise}
   */
  async updateCategory(id, payload) {
    const { data } = await api.put(`/admin/categories/${id}`, payload)
    return data
  },

  /**
   * Hapus kategori
   * @param {number} id
   * @returns {Promise}
   */
  async deleteCategory(id) {
    const { data } = await api.delete(`/admin/categories/${id}`)
    return data
  },

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
