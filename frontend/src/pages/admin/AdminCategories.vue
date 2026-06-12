<template>
  <div class="flex h-screen overflow-hidden bg-neutral-50">
    <AdminSidebar />

    <main class="flex-1 p-8 overflow-y-auto">
      <!-- Header -->
      <div class="flex items-center justify-between mb-6">
        <div>
          <h1 class="font-heading font-bold text-2xl text-neutral-800">Manajemen Kategori</h1>
          <p class="text-sm text-neutral-400 mt-0.5">{{ categories.length }} kategori terdaftar</p>
        </div>
        <BaseButton variant="primary" @click="openCreateModal">
          <Plus :size="18" /> Tambah Kategori
        </BaseButton>
      </div>

      <!-- Search -->
      <div class="relative max-w-sm mb-5">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Cari kategori..."
          class="input-field pr-10 py-2.5 text-sm"
          aria-label="Cari kategori"
        />
        <Search :size="18" class="absolute right-3 top-1/2 -translate-y-1/2 text-neutral-400" />
      </div>

      <!-- Loading -->
      <div v-if="loading" class="card p-5 space-y-3">
        <div v-for="n in 5" :key="n" class="flex gap-4 animate-pulse py-2.5">
          <div class="flex-1 h-4 bg-neutral-200 rounded" />
          <div class="h-4 w-32 bg-neutral-200 rounded" />
          <div class="h-8 w-20 bg-neutral-200 rounded" />
        </div>
      </div>

      <!-- Empty -->
      <EmptyState
        v-else-if="!filteredCategories.length"
        title="Belum ada kategori"
        description="Tambahkan kategori pertama kamu."
      >
        <BaseButton variant="primary" @click="openCreateModal">
          <Plus :size="16" /> Tambah Kategori
        </BaseButton>
      </EmptyState>

      <!-- Table -->
      <div v-else class="card overflow-hidden">
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead class="bg-neutral-50 border-b border-neutral-100">
              <tr>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">ID</th>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Nama Kategori</th>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Dibuat</th>
                <th class="text-right py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Aksi</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="cat in filteredCategories"
                :key="cat.id"
                class="border-b border-neutral-50 hover:bg-neutral-50/50 transition-colors"
              >
                <td class="py-3 px-4 text-neutral-400 text-xs font-mono">{{ cat.id }}</td>
                <td class="py-3 px-4">
                  <span class="font-medium text-neutral-800">{{ cat.name }}</span>
                </td>
                <td class="py-3 px-4 text-neutral-400 text-xs">{{ formatDate(cat.created_at) }}</td>
                <td class="py-3 px-4">
                  <div class="flex items-center justify-end gap-2">
                    <button
                      @click="openEditModal(cat)"
                      class="p-1.5 rounded-lg text-neutral-500 hover:bg-blue-50 hover:text-blue-600 transition-colors"
                      aria-label="Edit kategori"
                    >
                      <Pencil :size="15" />
                    </button>
                    <button
                      @click="confirmDelete(cat)"
                      class="p-1.5 rounded-lg text-neutral-500 hover:bg-red-50 hover:text-red-600 transition-colors"
                      aria-label="Hapus kategori"
                    >
                      <Trash2 :size="15" />
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </main>
  </div>

  <!-- Create / Edit Modal -->
  <BaseModal v-model="showModal" :title="editingCategory ? 'Edit Kategori' : 'Tambah Kategori'" size="sm">
    <form @submit.prevent="handleSave" class="space-y-4" novalidate>
      <BaseInput
        v-model="form.name"
        label="Nama Kategori"
        placeholder="cth: Electronics, Fashion..."
        :error="formErrors.name"
        required
      />
      <div v-if="formError" class="text-sm text-red-500 bg-red-50 px-4 py-2 rounded-lg">{{ formError }}</div>
    </form>
    <template #footer>
      <div class="flex justify-end gap-3">
        <BaseButton variant="ghost" @click="showModal = false">Batal</BaseButton>
        <BaseButton variant="primary" :loading="saving" @click="handleSave">
          {{ editingCategory ? 'Simpan Perubahan' : 'Tambah Kategori' }}
        </BaseButton>
      </div>
    </template>
  </BaseModal>

  <!-- Delete confirm modal -->
  <BaseModal v-model="showDeleteModal" title="Hapus Kategori?" size="sm">
    <p class="text-sm text-neutral-600">
      Kamu akan menghapus kategori <strong>{{ deletingCategory?.name }}</strong>. Produk yang menggunakan kategori ini mungkin terpengaruh.
    </p>
    <template #footer>
      <div class="flex justify-end gap-3">
        <BaseButton variant="ghost" @click="showDeleteModal = false">Batal</BaseButton>
        <BaseButton variant="danger" :loading="deleting" @click="handleDelete">Hapus</BaseButton>
      </div>
    </template>
  </BaseModal>
</template>

<script setup>
import { ref, computed, reactive, onMounted } from 'vue'
import { Plus, Search, Pencil, Trash2 } from 'lucide-vue-next'
import AdminSidebar   from '@/components/layout/AdminSidebar.vue'
import BaseButton     from '@/components/ui/BaseButton.vue'
import BaseInput      from '@/components/ui/BaseInput.vue'
import BaseModal      from '@/components/ui/BaseModal.vue'
import EmptyState     from '@/components/ui/EmptyState.vue'
import adminService   from '@/services/admin.service'
import categoryService from '@/services/category.service'
import { useToast }   from '@/composables/useToast'

const toast = useToast()

const categories      = ref([])
const loading         = ref(false)
const searchQuery     = ref('')
const showModal       = ref(false)
const showDeleteModal = ref(false)
const editingCategory  = ref(null)
const deletingCategory = ref(null)
const saving  = ref(false)
const deleting = ref(false)
const formError = ref('')

const form = reactive({ name: '' })
const formErrors = reactive({ name: '' })

const filteredCategories = computed(() => {
  if (!searchQuery.value) return categories.value
  const q = searchQuery.value.toLowerCase()
  return categories.value.filter((c) => c.name.toLowerCase().includes(q))
})

async function fetchCategories() {
  loading.value = true
  try {
    const res = await categoryService.getAll()
    categories.value = res.data || []
  } finally {
    loading.value = false
  }
}

function openCreateModal() {
  editingCategory.value = null
  form.name = ''
  formErrors.name = ''
  formError.value = ''
  showModal.value = true
}

function openEditModal(cat) {
  editingCategory.value = cat
  form.name = cat.name
  formErrors.name = ''
  formError.value = ''
  showModal.value = true
}

function validateForm() {
  formErrors.name = ''
  if (!form.name.trim()) {
    formErrors.name = 'Nama kategori wajib diisi.'
    return false
  }
  return true
}

async function handleSave() {
  if (!validateForm()) return
  saving.value    = true
  formError.value = ''
  try {
    if (editingCategory.value) {
      await adminService.updateCategory(editingCategory.value.id, { name: form.name.trim() })
      toast.success('Kategori berhasil diperbarui.')
    } else {
      await adminService.createCategory({ name: form.name.trim() })
      toast.success('Kategori berhasil ditambahkan.')
    }
    showModal.value = false
    await fetchCategories()
  } catch (err) {
    formError.value = err.response?.data?.message || 'Gagal menyimpan kategori.'
  } finally {
    saving.value = false
  }
}

function confirmDelete(cat) {
  deletingCategory.value = cat
  showDeleteModal.value  = true
}

async function handleDelete() {
  deleting.value = true
  try {
    await adminService.deleteCategory(deletingCategory.value.id)
    toast.success('Kategori berhasil dihapus.')
    showDeleteModal.value = false
    await fetchCategories()
  } catch {
    toast.error('Gagal menghapus kategori.')
  } finally {
    deleting.value = false
  }
}

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })
}

onMounted(fetchCategories)
</script>
