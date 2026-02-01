# syntax=docker/dockerfile:1

# ===========================
# builder stage
# ===========================
ARG GO_VERSION=1.25.6
FROM golang:${GO_VERSION}-trixie AS builder

WORKDIR /backend

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,source=.,target=. \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
    go build -ldflags="-s -w" -trimpath -o main .

# ===========================
# local stage
# ===========================
FROM golang:${GO_VERSION}-trixie AS local

WORKDIR /backend

RUN <<EOF
go install github.com/air-verse/air@latest
go install github.com/go-delve/delve/cmd/dlv@latest
go install golang.org/x/tools/gopls@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
EOF

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

EXPOSE 3000

CMD ["air"]

# ===========================
# deploy stage
# ===========================
FROM gcr.io/distroless/static-debian13:nonroot AS deploy

WORKDIR /app

COPY --from=builder --chown=nonroot:nonroot /backend/main /app/main

EXPOSE 80

USER nonroot:nonroot

CMD ["/app/main"]
