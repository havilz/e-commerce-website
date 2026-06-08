import { defineStore } from 'pinia'
import { ref } from 'vue'
import categoryService from '@/services/category.service'

export const useCategoryStore = defineStore('category', () => {
  // State
  const categories = ref([])
  const loading    = ref(false)

  // Actions
  async function fetchAll() {
    loading.value = true
    try {
      const res    = await categoryService.getAll()
      categories.value = res.data || []
    } catch {
      categories.value = []
    } finally {
      loading.value = false
    }
  }

  async function create(payload) {
    const res = await categoryService.create(payload)
    await fetchAll()
    return res
  }

  async function update(id, payload) {
    const res = await categoryService.update(id, payload)
    await fetchAll()
    return res
  }

  async function remove(id) {
    const res = await categoryService.remove(id)
    await fetchAll()
    return res
  }

  return { categories, loading, fetchAll, create, update, remove }
})
