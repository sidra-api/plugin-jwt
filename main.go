package main

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sidra-gateway/go-pdk/server"
)

var secretKey = []byte("secret-key")

func verifyJWT(tokenString string) (bool, error, *jwt.Token) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return false, err, token
	}
	return true, nil, token
}

func handler(r server.Request) server.Response {
	authHeader := r.Headers["Authorization"]

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return server.Response{
			StatusCode: 401,
			Body:       "Unauthorized",
		}
	}

	tokenString := authHeader[len("Bearer "):]

	valid, err, token := verifyJWT(tokenString)

	if !valid || err != nil {
		return server.Response{
			StatusCode: 401,
			Body:       "Unauthorized",
		}
	}

	payloads, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return server.Response{
			StatusCode: 401,
			Body:       "Unauthorized - Invalid token claims",
		}
	}

	headers := convertHeaders(token.Header)
	headers["iat"] = fmt.Sprintf("%v", payloads["iat"])
	headers["exp"] = fmt.Sprintf("%v", payloads["exp"])
	headers["sub"] = fmt.Sprintf("%v", payloads["sub"])

	username, ok := payloads["username"].(string)
    if !ok {
        return server.Response{
            StatusCode: 401,
            Body:       "Unauthorized - Username not found in token",
        }
    }

    headers["username"] = username

	return server.Response{
		StatusCode: 200,
		Headers:    headers,
	}
}

func convertHeaders(headers map[string]interface{}) map[string]string {
	converted := make(map[string]string)
	for key, value := range headers {
		converted[key] = fmt.Sprintf("%v", value)
	}
	return converted
}

func main() {
	server.NewServer("jwt", func(r server.Request) server.Response {
		return handler(r)
	}).Start()
}
