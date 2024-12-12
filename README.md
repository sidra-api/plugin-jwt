# Plugin JWT

## Deskripsi
Plugin JWT digunakan untuk memvalidasi JSON Web Token (JWT) pada Sidra Api. Plugin ini memastikan bahwa hanya permintaan dengan token yang valid dan terverifikasi yang dapat diakses, menjaga keamanan layanan backend.

---

## Cara Kerja
1. **Header Authorization**
   - Plugin membaca header `Authorization` pada setiap permintaan.
   - Token JWT harus diawali dengan prefix `Bearer `.

2. **Validasi Token**
   - Token JWT diverifikasi menggunakan library `github.com/golang-jwt/jwt/v4`.
   - Metode signing yang didukung adalah HMAC dengan kunci rahasia `secret-key`.

3. **Claims dan Headers**
   - Setelah token valid, plugin mengambil informasi dari payload dan header token.
   - Informasi seperti `username`, `iat`, dan `exp` akan dimasukkan dalam respons header.

4. **Respon**
   - Jika token valid:
     - Status: `200 OK`
     - Body: Kosong
     - Headers: Berisi informasi dari token seperti `username`, `iat`, `exp`, dll.
   - Jika token tidak valid:
     - Status: `401 Unauthorized`
     - Body: Pesan kesalahan seperti "Unauthorized" atau "Invalid token claims".

---

## Konfigurasi
- **Kunci Rahasia**: `secret-key`
  - Pastikan untuk mengganti kunci rahasia pada file `main.go` agar lebih aman.

---

## Cara Menjalankan
1. Pastikan Anda sudah menginstal **Sidra Api**.
2. Tambahkan plugin ini ke direktori `plugins/jwt/main.go` pada Sidra Api.
3. Kompilasi dan jalankan Sidra Api.
4. Plugin akan otomatis terhubung melalui UNIX socket pada path `/tmp/jwt.sock`.

---

## Pengujian

### Endpoint
- **URL**: Endpoint mana saja yang dikonfigurasi untuk melewati plugin JWT.

### Langkah Pengujian
1. Kirim request dengan header `Authorization` berisi JWT menggunakan Postman:
   ```plaintext
   GET http://localhost:3080/api/v1/resource
   Authorization: Bearer <token-jwt-anda>
   ```
2. Respons yang diharapkan:
   - Jika token valid, Anda akan mendapatkan status `200 OK` dengan header tambahan dari token.
   - Jika token tidak valid, respons akan berisi status `401 Unauthorized` dengan pesan kesalahan.

---

## Catatan Penting
- **Kunci Rahasia**: Jangan gunakan kunci rahasia default (`secret-key`) pada lingkungan produksi.
- **Durasi Token**: Pastikan `exp` (expiration) token dikelola dengan benar untuk mencegah akses tidak sah.

---

## Lisensi
MIT License
