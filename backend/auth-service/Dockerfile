FROM golang:1.20-alpine AS builder`

# Set environment variables
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

# Create a directory for the app
WORKDIR /app

# Copy go.mod and go.sum (to handle go dependencies) to the app directory
COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o auth-service main.go

# Step 2: Build a small image to run the app
FROM alpine:latest

# Set up certificates (required for HTTPS connections)
RUN apk --no-cache add ca-certificates

# Create a directory for the app
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/auth-service .

# Expose the port
EXPOSE 8080

# Start the app
CMD ["./auth-service"]




