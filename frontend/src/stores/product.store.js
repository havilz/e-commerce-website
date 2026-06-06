import { defineStore } from 'pinia'
import { ref, reactive } from 'vue'
import productService from '@/services/product.service'

export const useProductStore = defineStore('product', () => {
  // State
  const products       = ref([])
  const currentProduct = ref(null)
  const loading        = ref(false)
  const filters = reactive({
    search:   '',
    category: '',
  })
  const pagination = reactive({
    page:  1,
    limit: 100,
    total: 0,
  })

  // Actions
  async function fetchAll(params = {}) {
    loading.value = true
    try {
      const res = await productService.getAll({
        search:   filters.search   || undefined,
        category: filters.category || undefined,
        page:     pagination.page,
        limit:    pagination.limit,
        ...params,
      })
      products.value   = res.data || []
      // Backend meta field is "limt" (typo in backend) not "limit"
      pagination.total = res.meta?.total ?? products.value.length
    } catch {
      products.value = []
    } finally {
      loading.value = false
    }
  }

  async function fetchById(id) {
    loading.value = true
    currentProduct.value = null
    try {
      const res = await productService.getById(id)
      currentProduct.value = res.data
    } finally {
      loading.value = false
    }
  }

  function setFilter(key, value) {
    filters[key]     = value
    pagination.page  = 1
  }

  function setPage(page) {
    pagination.page = page
  }

  return {
    products,
    currentProduct,
    loading,
    filters,
    pagination,
    fetchAll,
    fetchById,
    setFilter,
    setPage,
  }
})
