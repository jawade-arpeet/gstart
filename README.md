# gstart

**gstart** is an opinionated project structure (boilerplate) for building REST API backends in Go. It wires up a layered architecture, database and cache clients, configuration management, and local infrastructure via Docker Compose — so new services can start from a consistent, production-leaning foundation instead of an empty `main.go`.

## Features

- **Layered architecture** — clear separation between router, handler, service, repository, and client layers.
- **Gin** HTTP framework with API versioning (`/api/v1/...`).
- **PostgreSQL** client built on `pgx/v5` with connection pooling.
- **Redis** and **Resend** clients for caching and transactional email.
- **Viper**-based config loading from `config.json`, validated at startup.
- **Environment-aware** setup that toggles Gin's release mode.
- **Docker Compose** stack for local Postgres and Redis with Makefile shortcuts.
- A working `/health` endpoint out of the box.

## Project Structure

```
gstart/
├── cmd/
│   └── main.go                  # Application entry point
├── internal/
│   ├── client/                  # External clients (Postgres, Redis, Resend)
│   │   ├── client.go
│   │   ├── postgres_client.go
│   │   ├── redis_client.go
│   │   └── resend_client.go
│   ├── config/                  # Config structs + loading via Viper
│   │   ├── config.go
│   │   ├── client_config.go
│   │   └── server_config.go
│   ├── constants/                # Shared constants (e.g. run environments)
│   │   └── constants.go
│   ├── handler/                  # HTTP handlers (Gin context in, JSON out)
│   │   ├── handler.go
│   │   └── health_handler.go
│   ├── middleware/                # Gin middleware
│   │   └── middleware.go
│   ├── repository/               # Data access layer
│   │   └── repository.go
│   ├── router/                    # Route registration
│   │   ├── router.go
│   │   └── v1/
│   │       ├── v1.go
│   │       └── health_router.go
│   ├── server/                     # Server bootstrapping
│   │   └── server.go
│   └── service/                    # Business logic layer
│       └── service.go
├── config.json                     # App configuration
├── docker-compose.yaml              # Local Postgres + Redis
├── Makefile
├── go.mod
└── go.sum
```

## Prerequisites

Go 1.26.5+ and Docker Compose.

## Getting Started

1. **Clone the repository**
    ```bash
    git clone https://github.com/jawade-arpeet/gstart.git
    cd gstart
    ```
2. **Configure the app** — Update `config.json` with your local settings.
3. **Start local infrastructure** — `make compose-up`
4. **Run the server** — `make run`
5. **Verify it's running** — `curl http://localhost:8080/api/v1/health`

## Configuration

Configuration is loaded from `config.json` and validated at startup. Missing required fields cause the app to fail fast.

```json
{
    "server": {
        "port": 8080,
        "run_env": "dev"
    },
    "client": {
        "postgres": {
            "host": "localhost",
            "port": 5432,
            "username": "username",
            "password": "password",
            "db_name": "gstart-db"
        },
        "redis": {
            "host": "localhost",
            "port": 6379,
            "username": "",
            "password": "password",
            "db_name": "0"
        },
        "resend": {
            "api_key": "re_xxxxxxxxx"
        }
    }
}
```

## Available Commands

| Command             | Description                                       |
| ------------------- | ------------------------------------------------- |
| `make run`          | Run the API server (`go run cmd/main.go`)         |
| `make tidy`         | Run `go mod tidy`                                 |
| `make fmt`          | Format the codebase with `gofmt`                  |
| `make compose-up`   | Start the Postgres and Redis containers           |
| `make compose-down` | Stop and remove the Postgres and Redis containers |

## Contributing

Issues and pull requests are welcome — please open an issue first for structural changes.
