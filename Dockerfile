# Build stage
FROM golang:1.24.4-alpine AS builder

# Ensure static build
RUN apk add --no-cache build-base

WORKDIR /app
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server main.go

# Final minimal image
FROM scratch

WORKDIR /app

# Copy binary only
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
