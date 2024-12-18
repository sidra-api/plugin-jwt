package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sidra-gateway/go-pdk/server"
)

var secretKey = []byte(getEnv("JWT_SECRET_KEY", "default-secret-key"))

// Fungsi getEnv membaca environment variable dengan nilai default jika tidak ditemukan
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Fungsi verifyJWT untuk memverifikasi token JWT
func verifyJWT(tokenString string) (bool, *jwt.Token, error) {
	// Parsing token menggunakan metode verifikasi HMAC dan secretKey
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("ERROR: Invalid signing method")
			return nil, logError("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		log.Printf("ERROR: Token parsing failed: %v\n", err)
		return false, token, err
	}

	if !token.Valid {
		log.Println("ERROR: Token is invalid")
	}
	return token.Valid, token, err
}

// Fungsi handler adalah entry point untuk memproses request di Sidra Api
func handler(r server.Request) server.Response {
	authHeader := r.Headers["Authorization"]

	if !strings.HasPrefix(authHeader, "Bearer ") {
		log.Println("INFO: Missing or invalid Authorization header")
		return server.Response{
			StatusCode: 401,
			Body:       "Unauthorized",
		}
	}

	tokenString := authHeader[len("Bearer "):]

	valid, token, err := verifyJWT(tokenString)

	if !valid || err != nil {
		log.Println("INFO: Unauthorized access attempt")
		return server.Response{
			StatusCode: 401,
			Body:       "Unauthorized",
		}
	}

	payloads, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("ERROR: Invalid token claims")
		return server.Response{
			StatusCode: 401,
			Body:       "Unauthorized - Invalid token claims",
		}
	}

	headers := convertHeaders(token.Header)
	headers["iat"] = logValue(payloads["iat"])
	headers["exp"] = logValue(payloads["exp"])
	headers["sub"] = logValue(payloads["sub"])

	username, ok := payloads["username"].(string)
    if !ok {
		log.Println("ERROR: Username not found in token claims")
        return server.Response{
            StatusCode: 401,
            Body:       "Unauthorized - Username not found in token",
        }
    }

    headers["username"] = username
	log.Printf("INFO: Successful authentication for username: %s\n", username)

	return server.Response{
		StatusCode: 200,
		Headers:    headers,
	}
}

// Fungsi logValue untuk mencatat nilai dari klaim token
func logValue(value interface{}) string {
	log.Printf("INFO: Claim value: %v\n", value)
	return fmt.Sprintf("%v", value)
}

// Fungsi logError untuk mencatat error dan mengembalikannya
func logError(message string) error {
	err := fmt.Errorf("%s", message)
	log.Println("ERROR:", err)
	return err
}

// Fungsi convertHeaders mengubah header token (interface{}) menjadi map[string]string
func convertHeaders(headers map[string]interface{}) map[string]string {
	converted := make(map[string]string)
	for key, value := range headers {
		converted[key] = fmt.Sprintf("%v", value)
	}
	return converted
}

func main() {
	log.Println("INFO: Starting JWT server")
	server.NewServer("jwt", func(r server.Request) server.Response {
		return handler(r)
	}).Start()
}
