<template>
  <div class="container-page py-8">
    <h1 class="font-heading font-bold text-2xl text-neutral-800 mb-6">Checkout</h1>

    <!-- Empty cart guard -->
    <div v-if="!cartStore.totalItems && !loading">
      <EmptyState title="Keranjang kosong" description="Tidak ada item untuk di-checkout.">
        <RouterLink to="/products"><BaseButton variant="primary">Belanja Dulu</BaseButton></RouterLink>
      </EmptyState>
    </div>

    <div v-else class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Form alamat -->
      <div class="lg:col-span-2 space-y-4">
        <div class="card p-6">
          <h2 class="font-heading font-bold text-lg mb-4 flex items-center gap-2">
            <MapPin :size="20" class="text-primary-500" />
            Alamat Pengiriman
          </h2>

          <form @submit.prevent="handleCheckout" novalidate>
            <div class="space-y-4">
              <BaseInput
                v-model="address"
                label="Alamat Lengkap"
                type="text"
                placeholder="Jl. Contoh No. 1, RT 01/RW 01, Kelurahan, Kecamatan, Kota"
                :error="addressError"
                required
              />

              <div v-if="errorMsg" class="bg-red-50 border border-red-200 rounded-lg px-4 py-3 text-sm text-red-600">
                {{ errorMsg }}
              </div>

              <BaseButton type="submit" variant="primary" size="lg" block :loading="loading">
                <ShoppingBag :size="18" />
                Konfirmasi Pesanan
              </BaseButton>
            </div>
          </form>
        </div>
      </div>

      <!-- Order summary -->
      <div class="card p-5 h-fit sticky top-24">
        <h2 class="font-heading font-bold text-lg mb-4">Ringkasan Pesanan</h2>

        <!-- Items -->
        <div class="space-y-3 mb-4 max-h-64 overflow-y-auto pr-1">
          <div
            v-for="item in cartStore.items"
            :key="item.id"
            class="flex items-center gap-3"
          >
            <img
              v-if="item.product?.image_url"
              :src="item.product.image_url"
              :alt="item.product?.name"
              class="w-12 h-12 object-cover rounded-lg flex-shrink-0"
            />
            <div class="flex-1 min-w-0">
              <p class="text-xs text-neutral-700 font-medium truncate">{{ item.product?.name }}</p>
              <p class="text-xs text-neutral-400">x{{ item.quantity }}</p>
            </div>
            <span class="text-xs font-bold text-neutral-800 flex-shrink-0">
              {{ formatPrice(item.product?.price * item.quantity) }}
            </span>
          </div>
        </div>

        <div class="border-t border-neutral-100 pt-4 space-y-2 text-sm">
          <div class="flex justify-between text-neutral-600">
            <span>Subtotal</span>
            <span>{{ formatPrice(cartStore.totalPrice) }}</span>
          </div>
          <div class="flex justify-between text-neutral-600">
            <span>Ongkir</span>
            <span class="text-green-600 font-medium">GRATIS</span>
          </div>
          <div class="flex justify-between font-bold text-base text-neutral-800 pt-1 border-t border-neutral-100">
            <span>Total</span>
            <span class="text-primary-600">{{ formatPrice(cartStore.totalPrice) }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { MapPin, ShoppingBag } from 'lucide-vue-next'
import { useCartStore }    from '@/stores/cart.store'
import orderService        from '@/services/order.service'
import BaseInput           from '@/components/ui/BaseInput.vue'
import BaseButton          from '@/components/ui/BaseButton.vue'
import EmptyState          from '@/components/ui/EmptyState.vue'
import { useToast }        from '@/composables/useToast'

const router    = useRouter()
const cartStore = useCartStore()
const toast     = useToast()

const address      = ref('')
const addressError = ref('')
const errorMsg     = ref('')
const loading      = ref(false)

async function handleCheckout() {
  addressError.value = ''
  errorMsg.value     = ''

  if (!address.value.trim()) {
    addressError.value = 'Alamat pengiriman wajib diisi.'
    return
  }
  if (address.value.trim().length < 10) {
    addressError.value = 'Alamat terlalu singkat, mohon isi dengan lengkap.'
    return
  }

  loading.value = true
  try {
    const res = await orderService.checkout(address.value.trim())
    await cartStore.clearCart()
    toast.success('Pesanan berhasil dibuat!')
    router.push({ name: 'OrderSuccess', query: { orderId: res.data?.id } })
  } catch (err) {
    errorMsg.value = err.response?.data?.message || 'Checkout gagal. Coba lagi.'
  } finally {
    loading.value = false
  }
}

function formatPrice(price = 0) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency', currency: 'IDR', maximumFractionDigits: 0,
  }).format(price)
}
</script>
