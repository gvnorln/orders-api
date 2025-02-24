# Orders API

Orders API adalah layanan sederhana yang dibuat menggunakan Go untuk mengelola pesanan. API ini mencakup operasi CRUD (Create, Read, Update, Delete) dan dapat dijalankan dalam container Docker.

## 🚀 Fitur
- **CRUD Order**: Tambah, baca, perbarui, dan hapus pesanan.
- **Dockerized**: Dapat dijalankan dalam container Docker.

---
## 📌 Struktur Proyek
```
/orders-api
├── handlers/order_handler.go    # Handler untuk order API
├── models/order.go              # Model data Order
├── routes/routes.go             # Routing API
├── main.go                      # Entry point aplikasi
├── go.mod                       # Dependencies Go
├── Dockerfile                   # Docker setup
└── README.md                    # Dokumentasi proyek
```

---
## 🛠️ Cara Menjalankan

### **1️⃣ Jalankan Secara Lokal**
Pastikan Go terinstal, lalu jalankan:
```sh
go run main.go
```
API akan berjalan di `http://localhost:8080`

### **2️⃣ Jalankan dengan Docker**
```sh
docker build -t orders-api .
docker run -d -p 8080:8080 orders-api
```
Cek API di `http://localhost:8080`

---
## 📦 API Endpoints
| Metode | Endpoint      | Deskripsi        |
|--------|-------------|----------------|
| GET    | `/orders`   | Ambil semua order |
| GET    | `/orders/{id}` | Ambil order by ID |
| POST   | `/orders`   | Buat order baru |
| PUT    | `/orders/{id}` | Update order |
| DELETE | `/orders/{id}` | Hapus order |

---
## 📜 Lisensi
Proyek ini menggunakan lisensi MIT. Silakan gunakan dan modifikasi sesuai kebutuhan!

