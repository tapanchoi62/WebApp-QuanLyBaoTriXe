FROM golang:1.25-alpine

# Install needed tools
RUN apk add --no-cache git bash

# Add go/bin to PATH
ENV PATH="/go/bin:${PATH}"

# Install Air (hot reload)
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Copy entire project
COPY . .

# Expose backend port
EXPOSE 8080

# Run Air
CMD ["air", "-c", ".air.toml"]
