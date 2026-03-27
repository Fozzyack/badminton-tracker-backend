# Backend

Go backend for the badminton tracker application.

## Requirements

- Go 1.26+
- Docker and Docker Compose

## Environment

Create a `backend/.env` file with at least:

```env
PORT=8080
DATABASE_URL=postgres://postrgres:postgres@localhost:5433/badminton?sslmode=disable
ENV=development
```

## Run Postgres

From the repo root:

```bash
docker compose -f backend/docker-compose.yml up -d
```

## Run the API

From `backend/`:

```bash
go run .
```

The server defaults to port `8080` if `PORT` is not set.

## Migrations

Migrations are embedded and run automatically when the application starts.

Current tables:

- `users`
- `sessions`
