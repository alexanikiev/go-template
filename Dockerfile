# --- Stage 1: Build the Go application ---
FROM golang:1.24-alpine AS builder

# Set working directory at the module root (same as where go.mod is located)
WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the entire project source code
COPY . .

# Build the solver binary from cmd/solver/
RUN go build -o solver ./cmd/solver

# --- Stage 2: Run the application ---
FROM alpine:latest

# Set working directory in the container
WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/solver .

# Set permissions (optional but recommended)
RUN chmod +x solver

# Run the solver application
CMD ["/app/solver"]
