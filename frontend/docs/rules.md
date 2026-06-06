# Aturan Pengerjaan Frontend

## Aturan Umum

* Frontend dikembangkan menggunakan Vue 3 + Vite + Tailwind CSS + Pinia + Vue Router.
* Proyek mengikuti pola **Clean Architecture** dengan pemisahan layer: `src/domain`, `src/application`, `src/infrastructure`, `src/presentation`.
* Semua kode harus mengikuti arah dependensi yang benar: layer dalam tidak boleh mengimpor layer luar.
* Sebelum memulai pengerjaan, buat terlebih dahulu file **task.md** di folder `frontend/docs` dengan status `Draft` — ini adalah langkah prioritas utama.
* Semua proses pengerjaan harus sesuai dengan task.md yang sudah dibuat.
* Setelah selesai mengerjakan task, update kembali task.md dengan status `Done`.
* Hanya berikan source code dan command sebagai output tanpa eksekusi langsung.
