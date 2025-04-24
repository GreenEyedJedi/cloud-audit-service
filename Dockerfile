# Use offical Go image as the builder for multistage building
FROM golang:1.22.2 AS builder

# Set working directory inside the container where project will live
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o cloud-audit-service ./cmd/cloud-audit-service

# Start a clean, minimal final image that is distroless (only contains binary and necessary sys libs)
# Final image with just enough to run the binary
FROM debian:bullseye-slim

# Install minimal required libs (optional if needed)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy only the binary and config from the builder stage
COPY --from=builder /app/cloud-audit-service /cloud-audit-service
COPY config.json /config.json

# Set the entrypoint
ENTRYPOINT [ "/cloud-audit-service" ]