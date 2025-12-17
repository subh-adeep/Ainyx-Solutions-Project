# Build Stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# Run Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/db/migrations ./db/migrations


# Install migrate tool (optional, but good for checking migrations inside container if needed)
# For simplicity, we assume migrations are run externally or via entrypoint script.
# But task requirement is just to include Dockerfile.

EXPOSE 8080

CMD ["./main"]
