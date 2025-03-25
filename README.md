# Go Inventory Management System 🚀

Blog : https://rohan-bhujbal.hashnode.dev/building-a-rest-api-with-go-from-zero-to-hero

A robust RESTful API service for managing product inventory, built with Go. Features CRUD operations, stock management, and pagination support.

## Tech Stack
- Go 1.22
- Gin Web Framework
- GORM
- PostgreSQL

## Features
- RESTful API endpoints
- Product CRUD operations
- Stock management
- Pagination
- Middleware for logging and error handling
- Structured logging

## Getting Started

### Prerequisites
- Go 1.22+
- Docker (for PostgreSQL)

### Local Setup

1. Clone the repository:
```bash
git clone https://github.com/rohan03122001/inventory-management-system
cd inventory-management-system
```

2. Start PostgreSQL using Docker:
```bash
docker run --name inv-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=inventory -p 5434:5432 -d postgres:14
```

3. Create `.env` file:
```env
DB_HOST=localhost
DB_PORT=5434
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=inventory
```

4. Run the application:
```bash
go mod download
go run cmd/api/main.go
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/health` | Health check |
| GET | `/api/v1/products` | List products (with pagination) |
| GET | `/api/v1/products/:id` | Get product by ID |
| POST | `/api/v1/products` | Create product |
| PUT | `/api/v1/products/:id` | Update product |
| DELETE | `/api/v1/products/:id` | Delete product |
| PATCH | `/api/v1/products/:id/stock` | Update stock |

### Example Requests

```bash
// Create Product
POST http://localhost:8080/api/v1/products
Content-Type: application/json
```
```json
{
    "name": "Playstation",
    "description": "Console",
    "price": 588.99,
    "quantity": 90,
    "sku": "IPH13-005"
}
```
```bash
//Get Products with Pagination
GET http://localhost:8080/api/v1/products?page=1&page_size=10
```
```bash
// Get Single Product
GET http://localhost:8080/api/v1/products/1
```
```bash
// Delete Product
DELETE http://localhost:8080/api/v1/products/1
```
```bash
// Update Stock
PATCH http://localhost:8080/api/v1/products/1/stock
Content-Type: application/json
```
```json
{
    "quantity": -5
}
```

## Project Structure
```
inventory-system/
├── cmd/
│   └── api/               # Application entry points
├── internal/
│   ├── config/           # Configuration management
│   ├── models/           # Data models
│   ├── repository/       # Database operations
│   ├── service/          # Business logic
│   ├── handler/          # HTTP handlers
│   └── middleware/       # Custom middleware
├── pkg/                  # Public libraries
│   └── database/         # Database utilities
└── tests/               # Integration tests
```

## TODO
- [ ] Authentication & Authorization
- [ ] Rate limiting
- [ ] Caching layer
- [ ] API documentation
- [ ] Unit tests
- [ ] Integration tests

## Running Tests
```bash
go test ./...
```

## Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
