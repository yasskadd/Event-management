# Stage 1: Build the Event Service
FROM golang:1.23-alpine AS builder

# Environment setup
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

# Working directory inside the container
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the application
RUN go build -o event-service main.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Add certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder
COPY --from=builder /app/event-service .

# Expose the service port
EXPOSE 8081

# Run the service
CMD ["./event-service"]
