# Book Inventory - Gin Learning Project

Project ini adalah aplikasi sederhana untuk belajar framework [Gin](https://gin-gonic.com/) di Go. Aplikasi ini mengelola data buku dengan fitur login berbasis JWT, daftar buku, detail buku, tambah buku, ubah buku, dan hapus buku.

## Fitur

- Login menggunakan username dan password sederhana
- Autentikasi request memakai JWT
- Menampilkan daftar buku
- Menampilkan detail buku
- Menambah buku baru
- Mengubah data buku
- Menghapus buku
- Auto migrate dan seed data awal saat aplikasi pertama kali dijalankan

## Tech Stack

- Go
- Gin Web Framework
- GORM
- PostgreSQL
- JWT
- Bootstrap untuk tampilan frontend

## Struktur Project

- `main.go` - entry point aplikasi
- `auth/` - handler login dan generate token JWT
- `middleware/` - middleware validasi JWT
- `app/` - handler untuk CRUD buku
- `db/` - koneksi database, migrasi, dan seeder
- `models/` - definisi model data
- `templates/` - file HTML untuk tampilan

## Prasyarat

- Go 1.26 atau lebih baru
- PostgreSQL

## Konfigurasi

Buat file `.env` di root project dan isi:

```env
POSTGRES_URL=postgres://username:password@localhost:5432/book_inventory?sslmode=disable
```

Sesuaikan `username`, `password`, host, port, dan nama database dengan konfigurasi lokal kamu.

## Cara Menjalankan

1. Pastikan PostgreSQL sudah aktif.
2. Siapkan file `.env`.
3. Jalankan aplikasi:

```bash
go run main.go
```

4. Buka browser ke:

```text
http://localhost:8080
```

## Login Default

Gunakan kredensial berikut:

- Username: `admin`
- Password: `123`

## Endpoint Utama

- `GET /` - redirect ke login
- `GET /login` - halaman login
- `POST /login` - proses login dan generate JWT
- `GET /books` - daftar buku
- `GET /book/:id` - detail buku
- `GET /addBook` - form tambah buku
- `POST /book` - simpan buku baru
- `GET /updateBook/:id` - form edit buku
- `POST /updateBook/:id` - update buku
- `POST /deleteBook/:id` - hapus buku

## Catatan

- Data buku awal akan di-seed otomatis jika tabel masih kosong.
- Banyak route buku menggunakan query parameter `auth` untuk membawa token JWT.
- Project ini cocok sebagai latihan memahami routing, middleware, template rendering, dan integrasi database di Gin.
