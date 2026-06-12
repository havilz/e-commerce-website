import api from './api'

const categoryService = {
  /**
   * Get semua kategori (public)
   * @returns {Promise}
   */
  async getAll() {
    const { data } = await api.get('/categories')
    return data
  },

  /**
   * Get detail kategori by ID (public)
   * @param {number|string} id
   * @returns {Promise}
   */
  async getById(id) {
    const { data } = await api.get(`/categories/${id}`)
    return data
  },

  /**
   * Buat kategori baru (Admin Only)
   * @param {object} payload - { name }
   * @returns {Promise}
   */
  async create(payload) {
    const { data } = await api.post('/admin/categories', payload)
    return data
  },

  /**
   * Update kategori (Admin Only)
   * @param {number} id
   * @param {object} payload - { name }
   * @returns {Promise}
   */
  async update(id, payload) {
    const { data } = await api.put(`/admin/categories/${id}`, payload)
    return data
  },

  /**
   * Hapus kategori (Admin Only)
   * @param {number} id
   * @returns {Promise}
   */
  async remove(id) {
    const { data } = await api.delete(`/admin/categories/${id}`)
    return data
  },
}

export default categoryService
