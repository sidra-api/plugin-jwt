# Plugin JWT for Sidra Api

This repository contains a plugin for Sidra Api that verifies JWT (JSON Web Tokens) in HTTP requests. The plugin uses HMAC signing for token validation and can be easily integrated into the Sidra Api.

---

## **Table of Contents**
- [Features](#features)
- [Environment Variables](#environment-variables)
- [Installation and Usage](#installation-and-usage)
  - [Build and Run with Docker](#build-and-run-with-docker)
- [Plugin Workflow](#plugin-workflow)
- [Endpoints and Responses](#endpoints-and-responses)
- [Generate JWT Token](#generate-jwt-token)
- [Testing](#testing)

---

## **Features**
- Verifies JWT tokens passed in the `Authorization` header.
- Extracts claims like `iat`, `exp`, `sub`, and `username` from the token.
- Supports dynamic configuration via environment variables.

---

## **Environment Variables**
This plugin supports the following environment variable:

| Variable         | Default Value        | Description                           |
|------------------|----------------------|---------------------------------------|
| `JWT_SECRET_KEY` | `default-secret-key` | The secret key used to validate JWTs. |

To override, set the `JWT_SECRET_KEY` variable in your runtime environment.

---

## **Installation and Usage**

### Build and Run with Docker

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/plugin-jwt.git
   cd plugin-jwt
   ```

2. **Build the Docker image:**
   ```bash
   docker build -t plugin-jwt .
   ```

3. **Run the Docker container:**
   ```bash
   docker run -e JWT_SECRET_KEY="your-secret-key" -p 8080:8080 plugin-jwt
   ```

The plugin will start and listen for incoming requests on port 8080.

---

## **Plugin Workflow**
1. The plugin extracts the `Authorization` header from incoming requests.
2. It validates the JWT token:
   - Ensures the token is signed using the `HS256` algorithm.
   - Verifies the token signature using the configured secret key.
3. If the token is valid, it extracts claims such as `iat`, `exp`, `username`, and includes them in the response headers.
4. If the token is invalid or missing, it returns a `401 Unauthorized` response.

---

## **Endpoints and Responses**
### Example Request:
```http
GET /api/v1/resource HTTP/1.1
Host: localhost:8080
Authorization: Bearer <your-token>
```

### Responses:
- **200 OK:**
  The token is valid, and claims are included in the response headers.
  ```json
  {
    "message": "Request authorized"
  }
  ```
  
  **Headers:**
  ```
  iat: <issued-at>
  exp: <expiry>
  sub: <subject>
  username: <username>
  ```

- **401 Unauthorized:**
  The token is invalid or missing.
  ```json
  {
    "error": "Unauthorized"
  }
  ```

---

## **Generate JWT Token**
Use the `generate/main.go` file to create a JWT token for testing purposes.

### Steps:
1. Run the `generate/main.go` script:
   ```bash
   go run generate/main.go
   ```

2. A token will be printed to the console:
   ```
   Generated JWT Token: Bearer <your-token>
   ```

Copy this token and use it in the `Authorization` header of your requests.

---

## **Testing**
To test the plugin:

1. Use a tool like Postman or cURL to send requests to the plugin.
2. Set the `Authorization` header with a valid JWT token.

### Example with cURL:
```bash
curl -X GET http://localhost:8080/api/v1/resource \
  -H "Authorization: Bearer <your-token>"
```

### Example with Postman:
- Create a new GET request.
- Set the URL to `http://localhost:8080/api/v1/resource`.
- Add the `Authorization` header with `Bearer <your-token>`.
- Send the request.

Ensure that the response contains the expected claims in the headers if the token is valid.
