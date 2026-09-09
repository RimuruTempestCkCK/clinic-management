# Clinic Management System API

Proyek backend ini dibuat sebagai sarana pembelajaran Golang. Aplikasi ini adalah RESTful API untuk Sistem Manajemen Klinik yang menangani berbagai operasional klinik seperti manajemen pengguna, dokter, pasien, rekam medis, hingga proses pembuatan janji temu (appointments).

## 🚀 Tech Stack

- **Go**: Bahasa pemrograman utama
- **Gin**: Framework web berkinerja tinggi untuk membangun REST API
- **GORM**: ORM (Object Relational Mapping) untuk berinteraksi dengan database
- **PostgreSQL / Supabase**: Sistem manajemen basis data relasional
- **JWT (JSON Web Token)**: Untuk autentikasi dan otorisasi yang aman
- **Validator**: Untuk memvalidasi payload request dari pengguna
- **Swagger**: Untuk dokumentasi API interaktif

## 👥 Role Pengguna

Sistem ini mendukung 4 peran (Role) dengan hak akses yang berbeda:
1. **ADMIN**: Memiliki kontrol penuh atas sistem.
2. **DOCTOR**: Dapat mengelola janji temu, rekam medis, dan resep pasiennya.
3. **STAFF**: Dapat membantu mengelola operasional harian dan jadwal.
4. **PATIENT**: Dapat membuat janji temu dan melihat rekam medis miliknya.

## 📌 Fitur Utama (Dalam Pengembangan)

- **Authentication**: Register, Login, Refresh, Logout
- **Manajemen Pasien & Dokter**: Operasi CRUD
- **Appointments (Janji Temu)**: Alur status janji temu (Pending -> Confirmed -> In Progress -> Completed -> Cancelled)
- **Rekam Medis (Medical Records)**
- **Resep & Obat (Prescriptions & Medicines)**

## 📂 Struktur Direktori

```text
clinic-management-api/
├── cmd/
│   └── server/
│       └── main.go       # Entry point aplikasi
├── internal/
│   ├── config/           # Konfigurasi aplikasi (Database, Env)
│   ├── database/         # Inisialisasi dan koneksi database
│   ├── middleware/       # Middleware untuk Auth dan Role checking
│   └── ...               # Domain modules (user, patient, doctor, dll)
├── docs/                 # Dokumentasi Swagger
├── migrations/           # Skrip migrasi database
└── go.mod                # Dependensi modul Go
```

## 🛠 Cara Menjalankan (Local Development)

1. Clone repositori ini.
2. Salin `.env.example` ke `.env` dan konfigurasikan koneksi database Anda.
3. Jalankan `go mod tidy` untuk mengunduh dependensi.
4. Jalankan aplikasi menggunakan perintah:
   ```bash
   go run cmd/server/main.go
   ```
5. Akses dokumentasi API via Swagger di `http://localhost:8080/swagger/index.html` (Setelah Swagger di-generate).

---
*Dibuat untuk tujuan edukasi dan pembelajaran ekosistem backend dengan Golang.*
