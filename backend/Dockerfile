# syntax=docker/dockerfile:1

# ===========================
# builder stage
# ===========================
ARG GO_VERSION=1.24.4
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /backend

COPY go.* .
RUN go mod download

COPY . .

RUN go build -o main .

# ===========================
# deploy stage
# ===========================
FROM debian:bookworm-slim

COPY --from=builder /backend/main /backend/main

EXPOSE 3000

CMD ["/backend/main"]
