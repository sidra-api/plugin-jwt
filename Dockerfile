# Gunakan base image untuk Golang
FROM golang:1.23 AS builder

# Set working directory di dalam container
WORKDIR /app

# Copy semua file plugin ke container
COPY . .

# Build binary plugin
RUN go mod tidy && go build -o plugin-jwt main.go

# Gunakan image minimal untuk hasil akhir
FROM alpine:latest

# Set Env Vars default (opsional, bisa di-overwrite saat runtime)
ENV JWT_SECRET_KEY=default-secret-key

# Copy binary dari stage builder ke stage ini
COPY --from=builder /app/plugin-jwt /usr/local/bin/plugin-jwt

# Jalankan binary
ENTRYPOINT ["/usr/local/bin/plugin-jwt"]
