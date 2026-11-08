# go-gin-gorm

REST API CRUD sederhana untuk data **user**, dibuat dengan [Go](https://go.dev/), [Gin](https://github.com/gin-gonic/gin), dan [GORM](https://gorm.io/) dengan database MySQL.

## Fitur

- CRUD data user (Create, Read, Update, Delete)
- Tabel dibuat otomatis lewat `AutoMigrate`
- Soft delete (data tidak benar-benar dihapus dari database)
- Respons dalam format JSON

## Teknologi

| Komponen | Keterangan |
|---|---|
| Bahasa | Go 1.21+ |
| Web framework | Gin |
| ORM | GORM |
| Database | MySQL (contoh memakai XAMPP) |

## Struktur Project

```
go-gin-gorm/
├── config/
│   └── database.go          # koneksi database
├── controller/
│   └── user_controller.go   # handler untuk setiap endpoint
├── model/
│   └── model.go             # struct User
├── router/
│   └── router.go            # daftar route
├── main.go                  # entry point aplikasi
├── go.mod
└── go.sum
```

## Persiapan

### 1. Install Go

Pastikan Go sudah terpasang:

```bash
go version
```

### 2. Jalankan MySQL

Jika memakai XAMPP, start service MySQL dari aplikasi XAMPP, atau lewat terminal:

```bash
sudo /Applications/XAMPP/xamppfiles/bin/mysql.server start
```

### 3. Buat database

```bash
/Applications/XAMPP/xamppfiles/bin/mysql -u root
```

```sql
CREATE DATABASE gin_gorm;
EXIT;
```

### 4. Atur koneksi database

Koneksi diatur di `config/database.go`:

```go
dsn := "root:@tcp(localhost:3306)/gin_gorm?charset=utf8mb4&parseTime=True&loc=Local"
```

Format DSN: `user:password@tcp(host:port)/nama_database?...`

Sesuaikan jika user, password, port, atau nama database kamu berbeda. Default XAMPP memakai user `root` tanpa password.

## Menjalankan Aplikasi

```bash
# clone repository
git clone <url-repository-kamu>
cd go-gin-gorm

# unduh dependency
go mod tidy

# jalankan
go run main.go
```

Server berjalan di `http://localhost:8080`. Tabel `users` akan dibuat otomatis saat aplikasi pertama kali dijalankan.

## Model User

| Field | Tipe | JSON | Keterangan |
|---|---|---|---|
| `ID` | uint | `ID` | primary key, otomatis |
| `CreatedAt` | datetime | `CreatedAt` | otomatis |
| `UpdatedAt` | datetime | `UpdatedAt` | otomatis |
| `DeletedAt` | datetime | `DeletedAt` | untuk soft delete |
| `Name` | string | `name` | nama user |
| `Address` | string | `address` | alamat user |
| `Age` | int | `age` | umur user |

## Endpoint API

| Method | Endpoint | Fungsi |
|---|---|---|
| `POST` | `/users` | Membuat user baru |
| `GET` | `/users` | Mengambil semua user |
| `GET` | `/users/:id` | Mengambil satu user berdasarkan ID |
| `PUT` | `/users/:id` | Mengubah data user |
| `DELETE` | `/users/:id` | Menghapus user |

### Contoh Penggunaan

**Membuat user**

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Budi","address":"Jakarta","age":25}'
```

Respons:

```json
{
  "data": {
    "ID": 1,
    "CreatedAt": "2026-11-08T10:00:00+07:00",
    "UpdatedAt": "2026-11-08T10:00:00+07:00",
    "DeletedAt": null,
    "name": "Budi",
    "address": "Jakarta",
    "age": 25
  }
}
```

**Mengambil semua user**

```bash
curl http://localhost:8080/users
```

**Mengambil satu user**

```bash
curl http://localhost:8080/users/1
```

**Mengubah user**

```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Budi Santoso","address":"Bandung","age":26}'
```

**Menghapus user**

```bash
curl -X DELETE http://localhost:8080/users/1
```

Respons:

```json
{
  "Message": "deleted successfully!"
}
```

### Respons Error

Jika data tidak ditemukan (`GET`, `PUT`, `DELETE` dengan ID yang tidak ada):

```json
{
  "error": "Record not found!"
}
```

Jika body JSON tidak valid saat membuat user, server mengembalikan status `400 Bad Request`.

## Melihat Data di Database

```bash
/Applications/XAMPP/xamppfiles/bin/mysql -u root
```

```sql
USE gin_gorm;
SELECT * FROM users;
```

## Rencana Pengembangan

- [ ] Validasi input (misalnya field `name` wajib diisi)
- [ ] Tambah pagination pada `GET /users`
- [ ] Tambah unit test

## Lisensi

Project ini dibuat untuk keperluan belajar.