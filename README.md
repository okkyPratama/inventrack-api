# InvenTrack API

InvenTrack (Inventory Tracking) API adalah sistem manajemen pengelolaan persediaan toko sederhana yang dibangun menggunakan Go. Sistem ini menyediakan API RESTful untuk mengelola kategori, produk, pemasok, dan transaksi.

## Fitur

- Autentikasi pengguna menggunakan JWT
- CRUD kategori
- CRUD produk
- CRUD Supplier
- Pencatatan transaksi (masuk/keluar) dengan pembaruan stok otomatis di produk

## Resources
- [Slide ppt](https://drive.google.com/file/d/1aEIQVi9-FHsiFpJL8aQ8udFuAhyUMGE7/view?usp=sharing)
- [Video penjelasan](https://drive.google.com/file/d/18at3YK26UPNEUEzT37D-ggzGtCyeGuMY/view?usp=sharing)
- [Postman Collection](https://drive.google.com/file/d/1X9RWmoSAOKbmcR4khFSneOO0gt6GGLNi/view?usp=sharing)

## ERD Database
![alt text](https://github.com/okkyPratama/inventrack-api/blob/main/inventrack.png?raw=true)

## Dependencies

- Go 1.15+
- PostgreSQL
- Godotenv
- Gorilla Mux
- JWT-Go
- Lib/pq
- Sql-migrate

## Instalasi

1. Klon repositori ini
2. Jalankan `go mod tidy` untuk menginstal dependensi
3. Jalankan aplikasi dengan `go run main.go`

## Endpoint API

### Autentikasi

- `POST /api/users/register` : Mendaftarkan pengguna baru
- `POST /api/users/login` : Login pengguna

### Kategori

- `GET /api/categories` : Mendapatkan semua kategori
- `POST /api/categories` : Membuat kategori baru
- `GET /api/categories/{id}` : Mendapatkan detail kategori
- `PUT /api/categories/{id}` : Memperbarui kategori
- `DELETE /api/categories/{id}` : Menghapus kategori

### Produk

- `GET /api/products` : Mendapatkan semua produk
- `POST /api/products` : Membuat produk baru
- `GET /api/products/{id}` : Mendapatkan detail produk
- `PUT /api/products/{id}` : Memperbarui produk
- `DELETE /api/products/{id}` : Menghapus produk

### Supplier

- `GET /api/suppliers` : Mendapatkan semua supplier
- `POST /api/suppliers` : Membuat supplier baru
- `GET /api/suppliers/{id}` : Mendapatkan detail supplier
- `PUT /api/suppliers/{id}` : Memperbarui supplier
- `DELETE /api/suppliers/{id}` : Menghapus supplier

### Transaksi

- `POST /api/transactions` : Membuat transaksi baru
- `GET /api/products/{id}/transactions` : Melihat riwayat transaksi untuk produk tertentu

## Penggunaan

1. Daftarkan pengguna baru menggunakan endpoint `/api/users/register`
2. Login menggunakan endpoint `/api/users/login` untuk mendapatkan token JWT
3. Gunakan token JWT di header `Authorization` untuk mengakses endpoint yang dilindungi
4. Kelola kategori, produk,dan supplier menggunakan endpoint yang sesuai
5. Catat transaksi masuk/keluar menggunakan endpoint transaksi

## Keamanan

- Semua endpoint (kecuali register dan login) dilindungi dengan middleware JWT
- Kata sandi di-hash sebelum disimpan di database
- Penggunaan prepared statements untuk mencegah SQL injection

## Deployment Link
- [https://inventrack-api.up.railway.app](https://inventrack-api.up.railway.app)

Untuk mengakses API, gunakan URL di atas sebagai base URL. Misalnya:
- Register: POST https://inventrack-api.up.railway.app/api/users/register
- Login: POST https://inventrack-api.up.railway.app/api/users/login
