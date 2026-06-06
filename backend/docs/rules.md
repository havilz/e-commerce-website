# Rules

## General Rules

* Backend di develope menggunakan native Go tanpa Framework, tapi menggunakan GORM untuk ORM.
* Sebelum pengerjaan task akan di buat ke dalam folder backend/docs, dengan nama file task.md.
* Semua proses pengerjaan harus sesuai dengan task.md yang sudah di buat.
* Setelah selesai mengerjakan task, update kembali task.md dengan status "Done".
* Hanya berikan source dan command sebagai output tanpa eksekusi langsung.

## GORM Rules

* Semua query harus menggunakan GORM.
* Jangan melakukan query langsung ke database.

## Struct Rules

* Semua Struct Model database yang dibuat harus menggunakan GORM.
* DTO (Data Transfer Object) dikecualikan dari aturan ini karena DTO hanya digunakan untuk memetakan request/response payload (JSON binding).
* Jangan membuat Struct Model yang tidak sesuai dengan query GORM.

## Module Rules

* Setiap module harus memiliki Handler, Service, Repository, dan DTO.
