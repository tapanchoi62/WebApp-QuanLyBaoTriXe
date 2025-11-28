# Build stage
FROM golang:1.21-bullseye AS builder

WORKDIR /app

# Set GOPROXY để tránh mạng bị chặn
ENV GOPROXY=https://proxy.golang.org,direct

# Copy go.mod và go.sum
COPY go.mod go.sum ./

# Copy toàn bộ source
COPY . .

# Build binary
RUN go build -o main .

# Run stage
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
