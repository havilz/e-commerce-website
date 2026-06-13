<template>
  <footer class="bg-neutral-800 text-neutral-300 mt-auto">
    <div class="container-page py-10">
      <div class="grid grid-cols-1 sm:grid-cols-3 gap-8">
        <!-- Brand -->
        <div>
          <div class="flex items-center gap-2 mb-3">
            <span class="font-heading font-extrabold text-xl text-white">ShopKu</span>
          </div>
          <p class="text-sm text-neutral-400 leading-relaxed">
            Belanja online mudah, cepat, dan menyenangkan. Temukan jutaan produk pilihan.
          </p>
        </div>

        <!-- Links -->
        <div>
          <h3 class="font-semibold text-white mb-3 text-sm uppercase tracking-wider">Layanan</h3>
          <ul class="space-y-2 text-sm">
            <li><RouterLink to="/products" class="hover:text-primary-400 transition-colors">Semua Produk</RouterLink></li>
            <li><RouterLink to="/orders" class="hover:text-primary-400 transition-colors">Pesanan Saya</RouterLink></li>
            <li><RouterLink to="/cart" class="hover:text-primary-400 transition-colors">Keranjang</RouterLink></li>
          </ul>
        </div>

        <!-- Contact -->
        <div>
          <h3 class="font-semibold text-white mb-3 text-sm uppercase tracking-wider">Hubungi Kami</h3>
          <ul class="space-y-2 text-sm text-neutral-400">
            <li>support@shopku.id</li>
            <li>0800-SHOPKU</li>
            <li>Senin–Jumat, 08.00–17.00</li>
          </ul>
        </div>
      </div>

      <div class="border-t border-neutral-700 mt-8 pt-6 text-center text-xs text-neutral-500">
        © {{ year }} ShopKu. Dibuat dengan Vue 3 + Go.
      </div>
    </div>
  </footer>
</template>

<script setup>
const year = new Date().getFullYear()
</script>
