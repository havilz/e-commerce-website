import api from './api'

const productService = {
  /**
   * Get all products dengan filter opsional
   * @param {object} params - { search, category, page, limit }
   * @returns {Promise}
   */
  async getAll(params = {}) {
    const { data } = await api.get('/products', { params })
    return data
  },

  /**
   * Get detail produk by ID
   * @param {number|string} id
   * @returns {Promise}
   */
  async getById(id) {
    const { data } = await api.get(`/products/${id}`)
    return data
  },
}

export default productService
