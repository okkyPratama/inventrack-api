# InvenTrack API

InvenTrack (Inventory Tracking) API adalah sistem manajemen pengelolaan inventory yang dibangun menggunakan Go. Sistem ini menyediakan API RESTful untuk mengelola kategori, produk, pemasok, gudang, dan transaksi.

## Fitur

- Autentikasi pengguna menggunakan JWT
- Manajemen kategori produk
- Manajemen produk
- Manajemen pemasok
- Manajemen gudang
- Pencatatan transaksi (masuk/keluar) dengan pembaruan stok otomatis
- Database PostgreSQL dengan migrasi otomatis

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

## Struktur Proyek

```
inventrack-api/
├── auth/
│   └── auth.go
├── config/
│   └── .env
├── controllers/
│   ├── category.go
│   ├── product.go
│   ├── supplier.go
│   ├── transaction.go
│   ├── user.go
│   └── warehouse.go
├── database/
│   ├── database.go
│   └── migrations/
│       ├── 1_initiate_master_table.sql
│       └── 2_initiate_transaction_table.sql
├── middleware/
│   └── jwt.go
├── repository/
│   ├── category.go
│   ├── product.go
│   ├── supplier.go
│   ├── transaction.go
│   └── warehouse.go
├── structs/
│   ├── category.go
│   ├── product.go
│   ├── supplier.go
│   ├── transaction.go
│   ├── user.go
│   └── warehouse.go
├── go.mod
├── go.sum
└── main.go
```
## ERD Database
![alt text](https://github.com/okkyPratama/inventrack-api/blob/main/inventrack.png?raw=true)

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

### Warehouse

- `GET /api/warehouses` : Mendapatkan semua gudang
- `POST /api/warehouses` : Membuat gudang baru
- `GET /api/warehouses/{id}` : Mendapatkan detail gudang
- `PUT /api/warehouses/{id}` : Memperbarui gudang
- `DELETE /api/warehouses/{id}` : Menghapus gudang

### Transaksi

- `POST /api/transactions` : Membuat transaksi baru
- `GET /api/products/{id}/transactions` : Melihat riwayat transaksi untuk produk tertentu

## Penggunaan

1. Daftarkan pengguna baru menggunakan endpoint `/api/users/register`
2. Login menggunakan endpoint `/api/users/login` untuk mendapatkan token JWT
3. Gunakan token JWT di header `Authorization` untuk mengakses endpoint yang dilindungi
4. Kelola kategori, produk, pemasok, dan gudang menggunakan endpoint yang sesuai
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
