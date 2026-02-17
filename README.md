# Gift Redemption API

REST API for a gift redemption system featuring authentication, RBAC, pagination, and a rating system.

Built with clean architecture and clear separation of concerns (handler/service/repository), this codebase emphasizes clean code and reusable design patterns. It is intentionally easy to review and run: one `make run` brings up the app, runs migrations + seeders, and generates Swagger docs. Unit tests cover core business logic, and the API is documented for quick exploration.

---

## Quick Start

```bash
# 1. Clone repository
git clone <repository-url>
cd gift-redemption

# 2. Setup environment
cp .env.example .env
# Edit .env according to your database configuration

# 3. Install dependencies
go mod tidy

# 4. Run database (via Docker)
docker compose up -d

# 5. Run application (auto run migration + seeder + generate swagger)
make run

# 6. Access Swagger UI
open http://localhost:8080/swagger/index.html
```

---

## 📚 Table of Contents

* Tech Stack
* Architecture
* Features
* Setup
* Testing
* API Documentation
* Bonus Implementation

---

## Tech Stack

### Core

* **Go 1.24** – Programming language
* **Gin** – HTTP web framework
* **GORM** – ORM library
* **PostgreSQL 14** – Database
* **JWT** – Authentication
* **Swagger** – API documentation

### Libraries

```
github.com/gin-gonic/gin v1.10.0
github.com/golang-jwt/jwt/v5 v5.2.1
github.com/joho/godotenv v1.5.1
github.com/stretchr/testify v1.9.0
github.com/swaggo/gin-swagger v1.6.0
golang.org/x/crypto v0.23.0
gorm.io/driver/postgres v1.5.9
gorm.io/gorm v1.25.10
```

---

## Architecture

### Clean Architecture (Layered)

```
┌─────────────────────────────────────────────────┐
│                   Handler                        │  HTTP Layer
│  (Routing, Request/Response, Validation)        │
└─────────────────┬───────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────┐
│                  Service                         │  Business Logic
│  (Use Cases, Business Rules, Transactions)      │
└─────────────────┬───────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────┐
│                Repository                        │  Data Access
│  (Database Queries, CRUD Operations)            │
└─────────────────┬───────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────┐
│                   Model                          │  Domain Entities
│  (Database Schema, Business Entities)           │
└─────────────────────────────────────────────────┘
```

### Project Structure

```
gift-redemption/
├── cmd/server/              # Application entrypoint
│   ├── main.go             # Bootstrap & DI
│   └── router.go           # Route definitions
├── internal/
│   ├── config/             # Configuration loader
│   ├── database/           # DB connection & migration
│   ├── dto/                # Request/Response structs
│   ├── handler/            # HTTP handlers (controllers)
│   ├── middleware/         # Auth, RBAC, etc
│   ├── model/              # Domain entities (GORM models)
│   ├── pkg/                # Shared utilities
│   │   ├── apperror/       # Custom error types
│   │   └── response/       # JSON response wrapper
│   ├── repository/         # Data access layer
│   │   └── mocks/          # Mock repositories for testing
│   └── service/            # Business logic layer
├── migrations/             # SQL migration files
├── seeds/                  # Database seeders
├── docs/                   # Swagger docs + documentation
│   ├── TESTING.md
│   └── *.postman_collection.json
├── docker-compose.yml
├── Makefile
└── README.md
```

### Design Patterns

* Repository Pattern – Abstracts data access
* Dependency Injection – Manual DI in `main.go`
* Interface Segregation – Each service/repository has its own interface
* Factory Pattern – Constructor functions `New...()`

---

## Features

### Core Features

* JWT-based user authentication
* Gift CRUD with pagination & sorting
* Gift redemption with stock validation
* Rating system (1–5) with star rounding
* Role-Based Access Control (Admin/User)
* Soft delete for users & gifts
* Transaction handling for stock deduction

### API Endpoints

| Method | Endpoint            | Auth | Role  | Description            |
| ------ | ------------------- | ---- | ----- | ---------------------- |
| POST   | `/login`            | -    | -     | User login             |
| GET    | `/gifts`            | ✓    | All   | List gifts (paginated) |
| GET    | `/gifts/:id`        | ✓    | All   | Get gift detail        |
| POST   | `/gifts`            | ✓    | Admin | Create gift            |
| PUT    | `/gifts/:id`        | ✓    | Admin | Update gift (full)     |
| PATCH  | `/gifts/:id`        | ✓    | Admin | Update gift (partial)  |
| DELETE | `/gifts/:id`        | ✓    | Admin | Delete gift            |
| POST   | `/gifts/:id/redeem` | ✓    | All   | Redeem gift            |
| POST   | `/gifts/:id/rating` | ✓    | All   | Rate gift              |
| GET    | `/users`            | ✓    | Admin | List users             |
| GET    | `/users/:id`        | ✓    | Admin | Get user detail        |
| POST   | `/users`            | ✓    | Admin | Create user            |
| PUT    | `/users/:id`        | ✓    | Admin | Update user            |
| DELETE | `/users/:id`        | ✓    | Admin | Delete user            |

---

## Setup

### Prerequisites

* Go 1.24+
* PostgreSQL 14+ (or Docker)
* Make (optional, for shortcuts)

### Local Environment Setup

**1. Clone & Install Dependencies**

```bash
git clone <repository-url>
cd gift-redemption
go mod tidy
```

**2. Environment Configuration**

```bash
cp .env.example .env
```

Edit `.env`:

```env
APP_PORT=8080
APP_ENV=development

DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gift_redemption

JWT_SECRET=your-super-secret-key
JWT_EXPIRY_HOURS=24
```

**3. Database Setup**

Option A – Using Docker (Recommended):

```bash
docker compose up -d
```

Option B – Local PostgreSQL:

```bash
createdb gift_redemption
```

**4. Generate Swagger Docs**

```bash
go install github.com/swaggo/swag/cmd/swag@latest
make generate
```

**5. Run Application**

```bash
make run
```

Server runs on `http://localhost:8080`.
Migrations and seeders run automatically at startup.

### Default Credentials

Admin:

```
Email: admin@gift-redemption.com
Password: password123
```

User:

```
Email: john@example.com
Password: password123
```

---

## Testing

### Automated Unit Testing

```bash
make test
make test-coverage
```

Coverage report will be generated as `coverage.html`.

Run a specific test:

```bash
go test -v ./internal/service -run TestAuthService_Login_Success
```

### Test Coverage

* Auth Service – login & JWT
* User Service – CRUD & validation
* Gift Service – CRUD, pagination, star rounding
* Redemption Service – business validation

Total: **16 unit tests** covering critical business logic.

---

## API Documentation

### Swagger UI

http://localhost:8080/swagger/index.html

### Response Format (JSON:API-like)

Success:

```json
{
  "meta": { "code": 200, "status": "success", "message": "success" },
  "data": []
}
```

Error:

```json
{
  "meta": { "code": 404, "status": "error", "message": "not found" },
  "errors": null
}
```

### Pagination & Sorting

```
GET /gifts?page=1&limit=10&sort_by=avg_rating&sort_dir=desc
```

### Star Rating System

Formula:

```
round(avg_rating * 2) / 2
```

---

## Bonus Implementation

### CRUD User

* Password hashing (bcrypt)
* Email uniqueness validation
* Role assignment (admin/user)

### RBAC

* JWT authentication middleware
* Role-based authorization

### Database Optimization

* Indexes for frequently queried columns
* Transactions with `SELECT FOR UPDATE`
* Connection pooling

---

## Makefile Commands

```bash
make run
make test
make test-coverage
make generate
make build
make tidy
make clean
```

---
