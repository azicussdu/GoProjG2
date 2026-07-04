# GoProjG2

This is a Go backend API project.

## Project structure

- `cmd/api/main.go` is the application entry point.
- `internal/config` handles environment and app configuration.
- `internal/model` contains domain models.
- `internal/handler` contains HTTP handlers and response formatting.
- `internal/service` contains business logic.
- `internal/repository` contains data access logic.
- `internal/apperrors` contains custom application errors.
- `docs` contains project documentation.

## Architecture rules

Follow layered architecture:

HTTP handler -> service -> repository -> database

Handlers should not contain business logic.
Services should not know HTTP details.
Repositories should not know HTTP or handler details.
Models should stay simple and reusable.

## Go conventions

- Use idiomatic Go.
- Keep functions small and clear.
- Always run `gofmt`.
- Prefer explicit errors over hidden behavior.
- Return errors instead of panicking.
- Keep packages focused and minimal.
- Avoid circular dependencies.
- Use context where database or external calls are involved.

## Error handling

Use `internal/apperrors` for application-level errors.
Use `internal/handler/error_response.go` for converting errors into HTTP responses.
Do not leak internal database errors directly to API clients.

## Repository layer

Repository interfaces should live in `internal/repository/interfaces.go`.
Database-specific implementations should stay separate, such as:

- `postgres_conn.go`

Do not put business rules in repository files.

## Service layer

Services should validate business rules and call repositories.
Services should return domain models or application errors.
Do not write HTTP responses from services.

## Handler layer

Handlers should:

- parse request input
- call services
- return JSON responses
- map errors using shared error response helpers

Handlers should not directly call repositories.

## Environment

Configuration comes from `.env`.
Do not commit secrets.
When adding new environment variables, document them in the README or docs.