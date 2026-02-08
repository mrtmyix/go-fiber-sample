# go-fiber-sample

Go Fiber + PostgreSQL sample application

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### Start development environment

```bash
docker compose up -d
```

### Access

- API: http://localhost:3000
- PostgreSQL: localhost:5432

### Stop environment

```bash
docker compose down
```

## Development

### Enter container

```bash
docker compose exec api bash
```

### Update go.mod

Inside the container:

```bash
go mod tidy
```

Or from host:

```bash
docker compose exec api go mod tidy
```

### Update dependencies to latest versions

Update all dependencies to the latest minor/patch versions:

```bash
# Inside the container
go get -u ./...
go mod tidy
```

Or from host:

```bash
docker compose exec api bash -c "go get -u ./... && go mod tidy"
```

Update specific package:

```bash
# Inside the container
go get -u github.com/gofiber/fiber/v3
go mod tidy
```

Update only patch versions (safer):

```bash
# Inside the container
go get -u=patch ./...
go mod tidy
```

