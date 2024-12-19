# **JWT Plugin**

## **Description**  
The JWT Plugin acts as middleware to verify JSON Web Tokens (JWT) on Sidra API. It checks the token's validity, claims, and responds accordingly based on verification results.

---

## **Key Features**  
- JWT verification using the HMAC method.  
- Supports **Environment Variables** for configuring the secret key.  
- Returns token claims in the response header.

---

## **Installation**

### **Prerequisites**  
- Go version 1.23 or higher.  
- Docker (optional for building an image).

### **Installation Steps**  
1. Clone the repository:  
   ```bash
   git clone <repository-url>
   cd plugin-jwt
   ```
2. Download dependencies:  
   ```bash
   go mod tidy
   ```
3. Build the plugin:  
   ```bash
   go build -o plugin-jwt main.go
   ```

### **Environment Variables Configuration**

The plugin uses `JWT_SECRET_KEY` to store the secret key. If this variable is not set, the plugin defaults to `default-secret-key`.

Set the environment variable using the following command:  
```bash
export JWT_SECRET_KEY="your-secret-key"
```

### **How to Run**

#### Using Go Binary
Run the binary directly:  
```bash
./plugin-jwt
```

#### Using Docker
Build the Docker image:  
```bash
docker build -t plugin-jwt .
```
Run the container:  
```bash
docker run -e JWT_SECRET_KEY="your-secret-key" -p 8080:8080 plugin-jwt
```

---

## **Workflow**

1. The client sends a request with an Authorization header containing the JWT in the format `Bearer <token>`.
2. The plugin verifies the token using the following steps:
   - Ensures the token uses a valid signing method (HMAC).
   - Checks the token claims (`iat`, `exp`, `sub`, `username`).
   - Adds the verified claims to the response header.
3. If the token is valid:
   - The plugin returns a 200 status with claims.
4. If the token is invalid:
   - The plugin returns a 401 status.

---

## **Testing**

1. **Generate a JWT Token**  
   Use the script in the `generate` folder to create a JWT token:
   ```bash
   cd generate
   export JWT_SECRET_KEY="your-secret-key"
   go run main.go
   ```

2. **Using Postman**  
   - Set the endpoint URL: `http://localhost:8080`
   - Add the following header: `Authorization: Bearer <token>`
   - Send the request:
     - If the token is valid: The response will include a 200 status and claims in the header.
     - If the token is invalid: The response will include a 401 status.

---

## **Sample Outputs**

### Successful Response
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

### Failed Response
```json
{
  "status_code": 401,
  "body": "Unauthorized"
}
```

---

## **Additional Notes**

- Ensure the secret key is consistent between the token generator and the JWT plugin.
- Tokens have an expiration time controlled by the `exp` claim. The plugin rejects expired tokens.
- Logs will provide detailed information about the request and verification process.