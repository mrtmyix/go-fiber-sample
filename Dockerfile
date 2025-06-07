# syntax=docker/dockerfile:1

# ===========================
# builder stage
# ===========================
ARG GO_VERSION=1.24.4
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .

RUN go build -o main .

# ===========================
# deploy stage
# ===========================
FROM debian:bookworm-slim

COPY --from=builder /app/main /app/main

CMD ["/app/main"]
