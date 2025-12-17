# Reasoning & Design Decisions

## Overview
The project is built using Go and Fiber for high performance and ease of use. SQLC is chosen for type-safe database interaction without the runtime overhead of reflection-based ORMs.

## Folder Structure
- `cmd/server`: Entry point.
- `internal`: Application code not meant for import by other projects.
    - `handler`: HTTP transport layer.
    - `service`: Business logic (e.g., age calculation).
    - `repository`: Data access wrapper around SQLC.
    - `models`: API specific structs.
- `pkg`: Publicly reusable code (e.g., `age` utility).
- `db`: Database related files (migrations, queries, generated code).

## DB Decisions
- **SQLC**: Generates Go code from SQL, ensuring queries are syntactically correct and type-safe at compile time.
- **Postgres**: robust relational database.
- **Migrations**: Plain SQL files for version control of schema.

## API Design
- **REST**: Standard resource-based routing.
- **Validation**: `go-playground/validator` ensures data integrity at the entry point.
- **Response**: Uniform JSON structure. Errors are returned as JSON objects.

## Logging & Middleware
- **Zap**: Structured, fast logging.
- **RequestID**: Tracing requests across logs.
- **Duration**: Logged for performance monitoring.

## Tests
- `pkg/util`: Unit tests for age calculation covering edge cases (leap years, current day birthdays).
- `go test ./...` runs all tests.

## Potential Improvements
- **Integration Tests**: Add E2E tests using a real DB container.
- **Configuration**: Use a more robust config library (e.g., Viper) for complex setups.
- **Pagination**: List users endpoint could support pagination for large datasets.
