# Go Site

A comprehensive, production-ready web application built with Go, delivering fast, secure, and scalable content for modern audiences. This repository includes backend services, static site generation, and deployment workflows.

---

## Table of Contents
1. [Overview](#overview)
2. [Features](#features)
3. [Prerequisites](#prerequisites)
4. [Installation](#installation)
5. [Configuration](#configuration)
6. [Project Structure](#project-structure)
7. [Usage](#usage)
   - [Running Locally](#running-locally)
   - [Building Static Content](#building-static-content)
8. [API Reference](#api-reference)
9. [Routing & Middleware](#routing--middleware)
10. [Deployment](#deployment)
11. [Testing](#testing)
12. [Logging & Monitoring](#logging--monitoring)
13. [Contributing](#contributing)
14. [License](#license)

---

## Overview

`go_site` is a modular web application framework in Go, designed to serve dynamic APIs and static site content. It emphasizes:

- **Speed**: Lightweight HTTP server powered by `net/http` and `fasthttp` benchmarks.
- **Security**: Built-in middleware for CSP, CORS, rate limiting, and input validation.
- **Scalability**: Microservice-ready architecture with clear separation of routers and handlers.
- **Developer Experience**: Auto-reload for development, CLI tools for code generation, and integrated Swagger documentation.

## Features

- 🛣️ **Router**: `chi`-based routing with nested routes and URL params.
- 🔒 **Security Middleware**: CSP, CORS, CSRF, rate limiting.
- ⚙️ **Config Management**: Environment-driven configuration using `viper`.
- 📄 **Static Site Generation**: Build markdown-based pages into HTML.
- 📚 **API Documentation**: Swagger UI served at `/docs`.
- ⏱️ **Graceful Shutdown**: Ensures active connections drain on SIGINT/SIGTERM.
- 📦 **Dockerized**: Ready-to-use `Dockerfile` and `docker-compose.yml`.

## Prerequisites

- Go 1.18+ installed
- `make` for build scripts
- Docker & Docker Compose (for local containers)

## Installation

```bash
# Clone the repository
git clone https://github.com/nicktretyakov/go_site.git
cd go_site

# Install Go dependencies
go mod download
```

## Configuration

Configuration is managed via environment variables or a `.env` file at project root:

| Variable           | Description                         | Default         |
|--------------------|-------------------------------------|-----------------|
| `APP_ENV`          | Application environment (`dev`,`prod`)| `dev`         |
| `SERVER_PORT`      | HTTP server port                    | `8080`          |
| `LOG_LEVEL`        | Logging level (`debug`,`info`,`error`)| `info`       |
| `RATE_LIMIT`       | Requests per minute                 | `100`           |
| `DB_DSN`           | Database connection string          | *required*      |

## Project Structure

```text
go_site/
├── cmd/                 # Entrypoints (e.g., server, migrate)
├── config/              # Configuration loader using viper
├── internal/
│   ├── api/             # API route handlers and controllers
│   ├── middleware/      # HTTP middleware implementations
│   ├── models/          # Database schema and ORM models
│   ├── repository/      # Data access layer (SQL, Redis)
│   ├── services/        # Business logic services
│   └── staticgen/       # Static site generator for markdown pages
├── migrations/          # Database migration files (SQL)
├── docs/                # Swagger definitions (YAML)
├── Makefile             # Build and dev commands
├── Dockerfile           # Docker image build
├── docker-compose.yml   # Service orchestration for local dev
├── .env.example         # Sample environment file
├── go.mod
└── README.md            # This documentation
```

## Usage

### Running Locally

```bash
# Development mode with live reload
make dev

# Production mode
go build -o bin/server ./cmd/server
./bin/server
```

### Building Static Content

```bash
# Generate HTML from markdown
go run ./cmd/staticgen -input content/ -output public/
```

## API Reference

Start the server and visit `http://localhost:8080/docs` to explore the Swagger UI. Endpoints include:

- `GET /health` ― health check
- `GET /api/v1/posts` ― list blog posts
- `GET /api/v1/posts/{id}` ― get post by ID
- `POST /api/v1/auth/login` ― user authentication

## Routing & Middleware

- **Router**: Powered by `github.com/go-chi/chi`.
- **Middleware**: Chained in `cmd/server.go`, including:
  - `middleware.Logger`
  - `middleware.RequestID`
  - `middleware.CORS`
  - `middleware.RateLimiter`
  - `middleware.Recoverer`

## Deployment

```bash
# Build Docker image
docker build -t go_site:latest .

# Run with Docker Compose
docker-compose up -d
```

## Testing

- **Unit Tests**:
  ```bash
  go test ./internal/... -cover
  ```
- **Integration Tests**:
  ```bash
  make test-integration
  ```

## Logging & Monitoring

- **Logging**: Structured JSON logs via `uber/zap`.
- **Metrics**: Prometheus metrics exposed at `/metrics`.
- **Tracing**: OpenTelemetry integration (Jaeger exporter) configured via `config/tracing.yaml`.

## Contributing

Contributions are welcome! Please:
1. Fork the repo
2. Create a feature branch
3. Write tests for new functionality
4. Submit a pull request with clear description

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

