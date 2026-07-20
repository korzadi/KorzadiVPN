# Build stage
FROM golang:1.26-alpine AS builder
RUN apk add --no-cache git build-base
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o korzadivpn .

# Final stage
FROM alpine:latest
RUN apk add --no-cache wireguard-tools sqlite-libs
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app
COPY --from=builder /app/korzadivpn .
COPY --from=builder /app/.env.example .env

# Exponer puerto
EXPOSE 8080

USER appuser
CMD ["./korzadivpn"]
