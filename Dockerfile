# ---- Build Stage ----
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git (required for go mod if using private repos)
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o schedule-checker ./cmd/schedule-checker/main.go

# ---- Run Stage ----
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/schedule-checker .

COPY .env .env

# Run the binary
CMD ["./schedule-checker"] 