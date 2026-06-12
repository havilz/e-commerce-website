import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1',
  timeout: 10000,
})

// Auto-attach JWT token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Handle 401 → auto logout, 429 → rate limit toast
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      // Redirect to login if not already there
      if (window.location.pathname !== '/login') {
        window.location.href = '/login'
      }
    } else if (error.response?.status === 429) {
      // Import and trigger toast on the fly for 429 errors
      try {
        const { useToast } = require('vue-toastification')
        const toast = useToast()
        toast.error('Terlalu banyak permintaan (Rate Limit). Silakan coba lagi nanti.')
      } catch (e) {
        // Fallback if dynamic require fails (ESM environment)
        import('vue-toastification').then(({ useToast }) => {
          const toast = useToast()
          toast.error('Terlalu banyak permintaan (Rate Limit). Silakan coba lagi nanti.')
        }).catch(() => {
          console.warn('Toast failed to load for 429 error')
        })
      }
    }
    return Promise.reject(error)
  }
)

export default api
