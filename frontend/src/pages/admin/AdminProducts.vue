<template>
  <div class="flex h-screen overflow-hidden bg-neutral-50">
    <AdminSidebar />

    <main class="flex-1 p-8 overflow-y-auto">
      <!-- Header -->
      <div class="flex items-center justify-between mb-6">
        <div>
          <h1 class="font-heading font-bold text-2xl text-neutral-800">Manajemen Produk</h1>
          <p class="text-sm text-neutral-400 mt-0.5">{{ products.length }} produk terdaftar</p>
        </div>
        <BaseButton variant="primary" @click="openCreateModal">
          <Plus :size="18" /> Tambah Produk
        </BaseButton>
      </div>

      <!-- Search -->
      <div class="relative max-w-sm mb-5">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Cari produk..."
          class="input-field pr-10 py-2.5 text-sm"
          aria-label="Cari produk"
        />
        <Search :size="18" class="absolute right-3 top-1/2 -translate-y-1/2 text-neutral-400" />
      </div>

      <!-- Loading -->
      <div v-if="loading" class="card p-5 space-y-3">
        <div v-for="n in 6" :key="n" class="flex gap-4 animate-pulse py-2">
          <div class="w-12 h-12 bg-neutral-200 rounded-lg flex-shrink-0" />
          <div class="flex-1 space-y-2">
            <div class="h-4 bg-neutral-200 rounded w-1/2" />
            <div class="h-3 bg-neutral-200 rounded w-1/4" />
          </div>
          <div class="h-8 w-16 bg-neutral-200 rounded" />
        </div>
      </div>

      <!-- Empty -->
      <EmptyState
        v-else-if="!filteredProducts.length"
        title="Belum ada produk"
        description="Tambahkan produk pertama kamu."
      >
        <BaseButton variant="primary" @click="openCreateModal">
          <Plus :size="16" /> Tambah Produk
        </BaseButton>
      </EmptyState>

      <!-- Table -->
      <div v-else class="card overflow-hidden">
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead class="bg-neutral-50 border-b border-neutral-100">
              <tr>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Produk</th>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Kategori</th>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Harga</th>
                <th class="text-left py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Stok</th>
                <th class="text-right py-3 px-4 text-xs font-semibold text-neutral-500 uppercase tracking-wide">Aksi</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="product in filteredProducts"
                :key="product.id"
                class="border-b border-neutral-50 hover:bg-neutral-50/50 transition-colors"
              >
                <!-- Product info -->
                <td class="py-3 px-4">
                  <div class="flex items-center gap-3">
                    <img
                      v-if="product.image_url"
                      :src="product.image_url"
                      :alt="product.name"
                      class="w-10 h-10 object-cover rounded-lg flex-shrink-0"
                      @error="(e) => e.target.style.display='none'"
                    />
                    <div v-else class="w-10 h-10 bg-neutral-100 rounded-lg flex-shrink-0 flex items-center justify-center">
                      <ImageOff :size="14" class="text-neutral-300" />
                    </div>
                    <span class="font-medium text-neutral-800 line-clamp-1">{{ product.name }}</span>
                  </div>
                </td>
                <td class="py-3 px-4">
                  <span class="text-xs bg-primary-50 text-primary-700 px-2 py-0.5 rounded-full font-medium">
                    {{ product.category || '—' }}
                  </span>
                </td>
                <td class="py-3 px-4 font-semibold text-neutral-800">{{ formatPrice(product.price) }}</td>
                <td class="py-3 px-4">
                  <span :class="['font-semibold text-xs', product.stock > 0 ? 'text-green-600' : 'text-red-500']">
                    {{ product.stock }}
                  </span>
                </td>
                <td class="py-3 px-4">
                  <div class="flex items-center justify-end gap-2">
                    <button
                      @click="openEditModal(product)"
                      class="p-1.5 rounded-lg text-neutral-500 hover:bg-blue-50 hover:text-blue-600 transition-colors"
                      aria-label="Edit produk"
                    >
                      <Pencil :size="15" />
                    </button>
                    <button
                      @click="confirmDelete(product)"
                      class="p-1.5 rounded-lg text-neutral-500 hover:bg-red-50 hover:text-red-600 transition-colors"
                      aria-label="Hapus produk"
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
  <BaseModal v-model="showModal" :title="editingProduct ? 'Edit Produk' : 'Tambah Produk'" size="lg">
    <form @submit.prevent="handleSave" class="space-y-4" novalidate>
      <BaseInput
        v-model="form.name"
        label="Nama Produk"
        placeholder="Nama produk"
        :error="formErrors.name"
        required
      />
      <div class="grid grid-cols-2 gap-4">
        <BaseInput
          v-model="form.price"
          label="Harga (Rp)"
          type="number"
          placeholder="99000"
          :error="formErrors.price"
          required
        />
        <BaseInput
          v-model="form.stock"
          label="Stok"
          type="number"
          placeholder="10"
          :error="formErrors.stock"
          required
        />
      </div>
      <div class="flex flex-col gap-1.5">
        <label class="text-sm font-medium text-neutral-700">Kategori <span class="text-red-500">*</span></label>
        <select
          v-model="form.category_id"
          class="input-field py-3 text-sm"
          :class="{ 'border-red-400': formErrors.category_id }"
        >
          <option value="">Pilih kategori</option>
          <option v-for="cat in categories" :key="cat.id" :value="String(cat.id)">{{ cat.name }}</option>
        </select>
        <p v-if="formErrors.category_id" class="text-xs text-red-500">{{ formErrors.category_id }}</p>
      </div>
      <BaseInput
        v-model="form.image_url"
        label="URL Gambar"
        placeholder="https://..."
        :error="formErrors.image_url"
      />
      <div class="flex flex-col gap-1.5">
        <label class="text-sm font-medium text-neutral-700">Deskripsi</label>
        <textarea
          v-model="form.description"
          rows="3"
          placeholder="Deskripsi produk..."
          class="input-field resize-none"
        />
      </div>
      <div v-if="formError" class="text-sm text-red-500 bg-red-50 px-4 py-2 rounded-lg">{{ formError }}</div>
    </form>
    <template #footer>
      <div class="flex justify-end gap-3">
        <BaseButton variant="ghost" @click="showModal = false">Batal</BaseButton>
        <BaseButton variant="primary" :loading="saving" @click="handleSave">
          {{ editingProduct ? 'Simpan Perubahan' : 'Tambah Produk' }}
        </BaseButton>
      </div>
    </template>
  </BaseModal>

  <!-- Delete confirm modal -->
  <BaseModal v-model="showDeleteModal" title="Hapus Produk?" size="sm">
    <p class="text-sm text-neutral-600">
      Kamu akan menghapus produk <strong>{{ deletingProduct?.name }}</strong>. Tindakan ini tidak dapat diurungkan.
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
import { Plus, Search, Pencil, Trash2, ImageOff } from 'lucide-vue-next'
import AdminSidebar    from '@/components/layout/AdminSidebar.vue'
import BaseButton      from '@/components/ui/BaseButton.vue'
import BaseInput       from '@/components/ui/BaseInput.vue'
import BaseModal       from '@/components/ui/BaseModal.vue'
import EmptyState      from '@/components/ui/EmptyState.vue'
import adminService    from '@/services/admin.service'
import productService  from '@/services/product.service'
import categoryService from '@/services/category.service'
import { useToast }    from '@/composables/useToast'

const toast = useToast()

const products   = ref([])
const categories = ref([])
const loading    = ref(false)
const searchQuery     = ref('')
const showModal       = ref(false)
const showDeleteModal = ref(false)
const editingProduct  = ref(null)
const deletingProduct = ref(null)
const saving   = ref(false)
const deleting = ref(false)
const formError = ref('')

const form = reactive({
  name: '', price: '', stock: '', category_id: '', image_url: '', description: '',
})

const formErrors = reactive({
  name: '', price: '', stock: '', category_id: '', image_url: '',
})

const filteredProducts = computed(() => {
  if (!searchQuery.value) return products.value
  const q = searchQuery.value.toLowerCase()
  return products.value.filter(
    (p) => p.name.toLowerCase().includes(q) || (p.category || '').toLowerCase().includes(q)
  )
})

async function fetchProducts() {
  loading.value = true
  try {
    const firstRes = await productService.getAll({ limit: 100, page: 1 })
    const total    = firstRes.meta?.total ?? firstRes.data?.length ?? 0
    let   all      = firstRes.data || []

    const totalPages = Math.ceil(total / 100)
    if (totalPages > 1) {
      const requests = []
      for (let p = 2; p <= totalPages; p++) {
        requests.push(productService.getAll({ limit: 100, page: p }))
      }
      const results = await Promise.all(requests)
      results.forEach((r) => { all = all.concat(r.data || []) })
    }
    products.value = all
  } finally {
    loading.value = false
  }
}

async function fetchCategories() {
  try {
    const res = await categoryService.getAll()
    categories.value = res.data || []
  } catch {
    categories.value = []
  }
}

function openCreateModal() {
  editingProduct.value = null
  Object.assign(form, { name: '', price: '', stock: '', category_id: '', image_url: '', description: '' })
  Object.keys(formErrors).forEach((k) => (formErrors[k] = ''))
  formError.value = ''
  showModal.value = true
}

function openEditModal(product) {
  editingProduct.value = product
  Object.assign(form, {
    name:        product.name,
    price:       String(product.price),
    stock:       String(product.stock),
    category_id: product.category_id ? String(product.category_id) : '',
    image_url:   product.image_url || '',
    description: product.description || '',
  })
  Object.keys(formErrors).forEach((k) => (formErrors[k] = ''))
  formError.value = ''
  showModal.value = true
}

function validateForm() {
  Object.keys(formErrors).forEach((k) => (formErrors[k] = ''))
  let valid = true
  if (!form.name.trim())                           { formErrors.name        = 'Nama wajib diisi.';   valid = false }
  if (!form.price || Number(form.price) <= 0)      { formErrors.price       = 'Harga tidak valid.';  valid = false }
  if (form.stock === '' || Number(form.stock) < 0) { formErrors.stock       = 'Stok tidak valid.';   valid = false }
  if (!form.category_id)                           { formErrors.category_id = 'Pilih kategori.';     valid = false }
  
  if (form.image_url.trim()) {
    const lowerURL = form.image_url.trim().toLowerCase()
    if (!lowerURL.startsWith('http://') && !lowerURL.startsWith('https://')) {
      formErrors.image_url = 'URL Gambar harus diawali dengan http:// atau https://';
      valid = false
    }
  }
  return valid
}

async function handleSave() {
  if (!validateForm()) return
  saving.value    = true
  formError.value = ''
  const payload = {
    name:        form.name.trim(),
    price:       Number(form.price),
    stock:       Number(form.stock),
    category_id: Number(form.category_id),
    image_url:   form.image_url.trim() || '',
    description: form.description.trim() || '',
  }
  try {
    if (editingProduct.value) {
      await adminService.updateProduct(editingProduct.value.id, payload)
      toast.success('Produk berhasil diperbarui.')
    } else {
      await adminService.createProduct(payload)
      toast.success('Produk berhasil ditambahkan.')
    }
    showModal.value = false
    await fetchProducts()
  } catch (err) {
    formError.value = err.response?.data?.message || 'Gagal menyimpan produk.'
  } finally {
    saving.value = false
  }
}

function confirmDelete(product) {
  deletingProduct.value = product
  showDeleteModal.value = true
}

async function handleDelete() {
  deleting.value = true
  try {
    await adminService.deleteProduct(deletingProduct.value.id)
    toast.success('Produk berhasil dihapus.')
    showDeleteModal.value = false
    await fetchProducts()
  } catch {
    toast.error('Gagal menghapus produk.')
  } finally {
    deleting.value = false
  }
}

function formatPrice(price) {
  return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', maximumFractionDigits: 0 }).format(price)
}

onMounted(() => {
  fetchProducts()
  fetchCategories()
})
</script>
