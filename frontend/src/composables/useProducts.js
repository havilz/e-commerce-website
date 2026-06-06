import { useProductStore } from '@/stores/product.store'
import { storeToRefs } from 'pinia'

export function useProducts() {
  const store = useProductStore()
  const { products, currentProduct, loading, filters, pagination } = storeToRefs(store)

  function search(query) {
    store.setFilter('search', query)
    store.fetchAll()
  }

  function filterByCategory(category) {
    store.setFilter('category', category)
    store.fetchAll()
  }

  function changePage(page) {
    store.setPage(page)
    store.fetchAll()
  }

  return {
    products,
    currentProduct,
    loading,
    filters,
    pagination,
    fetchAll: store.fetchAll,
    fetchById: store.fetchById,
    search,
    filterByCategory,
    changePage,
  }
}
