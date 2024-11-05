package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("secret-key")

func main() {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 100).Unix()
	claims["username"] = "foo"

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatalf("Error signing token: %v", err)
	}

	

	tokenString = "Bearer " + tokenString

	fmt.Println("Generated JWT Token: ", tokenString)
}
