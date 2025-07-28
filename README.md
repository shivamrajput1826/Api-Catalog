# API Catalog

A Go-based service for managing and cataloging API events, properties, and tracking plans. Built with [Fiber](https://gofiber.io/), [GORM](https://gorm.io/), and [Swagger](https://swagger.io/) for API documentation.

---

## Features

- **Event Management:** Create, update, delete, and list events.
- **Property Management:** Manage properties associated with events.
- **Tracking Plans:** Organize events and properties into tracking plans.
- **Validation:** Request validation using struct tags and custom logic.
- **Transaction Support:** Safe, atomic operations using GORM transactions.
- **Swagger Documentation:** Auto-generated API docs at `/swagger/index.html`.

---

## Project Structure

```
api-catalog/
├── cmd/api/           # Main application entrypoint
│   └── main.go
├── config/            # Configuration loading
├── internal/
│   ├── db/            # Database connection and migration
│   ├── dtos/          # Data transfer objects (request/response)
│   ├── handlers/      # HTTP handlers
│   ├── models/        # Database models and interfaces
│   ├── repositories/  # Data access layer (repositories)
│   ├── routes/        # Route definitions
│   ├── services/      # Business logic
│   └── validation/    # Custom validation logic
├── logger/            # Logging setup
├── middleware/        # Fiber middleware
└── go.mod
```

---

## Getting Started

### Prerequisites

- Go 1.18+
- A running database (e.g., PostgreSQL, MySQL, SQLite)

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/api-catalog.git
   cd api-catalog
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Configure environment variables:**
   - Edit or create a `.env` file or update `config/config.go` as needed.

4. **Run database migrations:**
   Migrations run automatically on startup.

5. **Start the server:**
   ```sh
   go run cmd/api/main.go
   ```

6. **Access Swagger docs:**
   - Visit [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## API Documentation

- **Swagger UI:** `/swagger/index.html`
- **Base Path:** `/`
- **Host:** `localhost:8080`

---

## Development

- **Generate Swagger docs:**
  ```sh
  swag init -g cmd/api/main.go
  ```

- **Run tests:**
  ```sh
  go test ./...
  ```

---

## License

MIT

---



