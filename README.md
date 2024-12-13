# Plugin JWT

## Deskripsi
Plugin JWT adalah sebuah middleware untuk memverifikasi token JWT (JSON Web Token) pada Sidra Api. Plugin ini memeriksa validitas token, klaim, dan memberikan respons yang sesuai berdasarkan hasil verifikasi.

## Fitur Utama
- Verifikasi JWT menggunakan metode HMAC.
- Mendukung penggunaan **Environment Variables** untuk konfigurasi secret key.
- Mengembalikan klaim token dalam header respons.

---

## Instalasi
### Prasyarat
- Go versi 1.23 atau lebih tinggi.
- Docker (opsional untuk membangun image).

### Langkah Instalasi
1. Clone repository:
   ```bash
   git clone <repository-url>
   cd plugin-jwt
   ```
2. Jalankan perintah berikut untuk mengunduh dependencies:
   ```bash
   go mod tidy
   ```

3. Build plugin:
   ```bash
   go build -o plugin-jwt main.go
   ```

---

## Konfigurasi Environment Variables
Plugin menggunakan **JWT_SECRET_KEY** untuk menyimpan secret key. Jika variabel ini tidak disetel, plugin akan menggunakan nilai default `default-secret-key`.

Set environment variable dengan perintah berikut:
```bash
export JWT_SECRET_KEY="your-secret-key"
```

---

## Cara Menjalankan
### Menggunakan Go Binary
1. Jalankan binary secara langsung:
   ```bash
   ./plugin-jwt
   ```

### Menggunakan Docker
1. Bangun image Docker:
   ```bash
   docker build -t plugin-jwt .
   ```
2. Jalankan container:
   ```bash
   docker run -e JWT_SECRET_KEY="your-secret-key" -p 8080:8080 plugin-jwt
   ```

---

## Workflow
1. Klien mengirim request dengan header **Authorization** yang berisi token JWT dalam format `Bearer <token>`.
2. Plugin memverifikasi token dengan langkah berikut:
   - Memastikan token menggunakan metode signing yang valid (HMAC).
   - Mengecek klaim token (iat, exp, sub, username).
   - Menambahkan klaim valid ke header respons.
3. Jika token valid, plugin mengembalikan status 200 dan klaim.
4. Jika token tidak valid, plugin mengembalikan status 401.

---

## Pengujian
### 1. Menghasilkan Token JWT
Gunakan script di folder `generate` untuk membuat token JWT:
```bash
cd generate
export JWT_SECRET_KEY="your-secret-key"
go run main.go
```

### 2. Menggunakan Postman
1. Setel endpoint URL:
   ```
   http://localhost:8080
   ```
2. Tambahkan header:
   ```
   Authorization: Bearer <token>
   ```
3. Kirim request:
   - Jika token valid, respons akan berupa status 200 dengan klaim dalam header.
   - Jika token tidak valid, respons akan berupa status 401.

---

## Contoh Output
### Respons Berhasil
```json
{
  "status_code": 200,
  "headers": {
    "iat": "1697029200",
    "exp": "1697032800",
    "sub": "foo",
    "username": "foo"
  }
}
```

### Respons Gagal
```json
{
  "status_code": 401,
  "body": "Unauthorized"
}
```

---

## Catatan Tambahan
- Pastikan secret key sama antara generator token dan plugin JWT.
- Token memiliki masa berlaku yang diatur melalui klaim `exp`. Plugin akan menolak token yang kedaluwarsa.
