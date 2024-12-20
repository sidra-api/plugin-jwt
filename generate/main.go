package main

import (
	"log"
	"os"
	// "time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte(getEnv("JWT_SECRET_KEY", "default-secret-key"))

// Fungsi getEnv membaca environment variable dengan nilai default jika tidak ditemukan
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	// Membuat token baru dengan metode signing HMAC
	token := jwt.New(jwt.SigningMethodHS256)

	// Menambahkan klaim (claims) ke token
	// claims := token.Claims.(jwt.MapClaims)
	// claims["iat"] = time.Now().Unix()
	// claims["exp"] = time.Now().Add(time.Minute * 100).Unix()
	// claims["username"] = "foo"

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatalf("Error signing token: %v", err)
	}

	// Menambahkan prefix "Bearer " agar token siap digunakan
	tokenString = "Bearer " + tokenString

	log.Println("Generated JWT Token: ", tokenString)
}
