# User Management API

A RESTful API in Go for managing users with dynamic age calculation.

## Tech Stack
- **Language**: Go
- **Framework**: Fiber (v2)
- **Database**: PostgreSQL
- **ORM/Access**: SQLC
- **Logging**: Uber Zap
- **Validation**: go-playground/validator

## Prerequisites
- Go 1.25+
- Docker & Docker Compose
- Make (optional)

## Getting Started

### 1. Run with Docker (Recommended)
This will start both the API and Postgres.
```bash
docker-compose up -d --build
```
API will be available at `http://localhost:8080`.

### 2. Run Locally
1. Start Postgres (or use docker for db only):
   ```bash
   docker-compose up -d db
   ```
2. Configure environment:
   ```bash
   cp .env.example .env 
   # Set DATABASE_URL=postgres://user:password@localhost:5432/max_db?sslmode=disable
   ```
3. Run migrations:
   (You can use `migrate` CLI or just execute SQL files in `db/migrations`)
   
4. Run App:
   ```bash
   go mod tidy
   go run cmd/server/main.go
   ```

### 3. Running Tests
```bash
go test ./...
```

## API Endpoints

- `POST /users`: Create user `{ "name": "Alice", "dob": "1990-01-01" }`
- `GET /users/:id`: Get user
- `PUT /users/:id`: Update user
- `DELETE /users/:id`: Delete user
- `GET /users`: List users (supports pagination: `?page=1&limit=10`)

## Implementation Details
- **Architecture**: Layered (Handler -> Service -> Repository -> DB)
- **Age Calculation**: Done dynamically in Go Service layer using `pkg/util`.
- **SQLC**: Database code generated from SQL queries.
