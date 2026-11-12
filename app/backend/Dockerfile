
# ==========================================
# Stage 1: Builder
# ==========================================
FROM golang:1.26.4-alpine AS builder

WORKDIR /app

# Install dependencies required for building
RUN apk add --no-cache git ca-certificates

# Copy dependency definitions first
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy application source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -mod=readonly \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/api \
    ./cmd/api/


# ==========================================
# Stage 2: Runtime
# ==========================================
FROM alpine:3.22 AS runtime

WORKDIR /app

# Runtime dependencies
RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S app \
    && adduser -S -G app app

# Copy compiled binary
COPY --from=builder /out/api ./api

# Run as non-root user
USER app

EXPOSE 8080

ENTRYPOINT ["./api"]
