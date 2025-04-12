# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Create a bin directory to place the compiled binary
RUN mkdir -p /app/bin

# Build the binary
RUN go build -o /app/bin/event-service main.go

# Confirm it was built correctly
RUN echo "[Builder] Contents of /app/bin:" && ls -lh /app/bin

# Stage 2: Runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/bin/event-service .

# Confirm binary is here
RUN echo "[Runtime] Contents of /root:" && ls -lh /root

EXPOSE 8081

CMD ["./event-service"]
