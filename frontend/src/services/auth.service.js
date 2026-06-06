import api from './api'

const authService = {
  /**
   * Login user
   * @param {string} email
   * @param {string} password
   * @returns {Promise}
   */
  async login(email, password) {
    const { data } = await api.post('/auth/login', { email, password })
    return data
  },

  /**
   * Register user baru
   * @param {string} name
   * @param {string} email
   * @param {string} password
   * @returns {Promise}
   */
  async register(name, email, password) {
    const { data } = await api.post('/auth/register', { name, email, password })
    return data
  },

  /**
   * Get current user info
   * @returns {Promise}
   */
  async getMe() {
    const { data } = await api.get('/auth/me')
    return data
  },
}

export default authService
