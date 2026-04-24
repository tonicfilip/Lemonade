# Lemonade API

A modern Go API boilerplate using the Gin framework.

## Features

- ✅ Gin web framework
- ✅ Modular project structure
- ✅ CORS middleware
- ✅ Environment configuration
- ✅ Custom logging
- ✅ Docker support
- ✅ Docker Compose for development
- ✅ Makefile for common tasks
- ✅ Example RESTful API endpoints

## Project Structure

```
lemonade/
├── config/           # Configuration management
├── handlers/         # HTTP request handlers
├── middleware/       # Middleware functions
├── routes/          # Route definitions
├── utils/           # Utility functions
├── main.go          # Application entry point
├── go.mod           # Go module definition
├── Dockerfile       # Docker configuration
├── docker-compose.yml # Docker Compose configuration
├── Makefile         # Build tasks
└── README.md        # This file
```

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose (optional)

## Getting Started

### 1. Clone or Initialize the Project

```bash
cd /Users/filip/Documents/projects/golang_learning/Lemonade
```

### 2. Download Dependencies

```bash
go mod download
go mod tidy
```

### 3. Set Up Environment Variables

```bash
cp .env.example .env
```

Edit `.env` as needed:

```env
PORT=8080
ENV=development
GIN_MODE=debug
```

### 4. Run the Application

#### Option A: Direct Run

```bash
go run main.go
```

#### Option B: Build and Run

```bash
make build
./lemonade
```

#### Option C: Using Docker Compose

```bash
docker-compose up
```

## API Endpoints

### Health Check

```bash
GET /health
```

Response:

```json
{
  "status": "ok",
  "code": 200
}
```

### Users Endpoints

#### Get all users

```bash
GET /api/v1/users
```

#### Get a specific user

```bash
GET /api/v1/users/:id
```

#### Create a new user

```bash
POST /api/v1/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com"
}
```

#### Update a user

```bash
PUT /api/v1/users/:id
Content-Type: application/json

{
  "name": "Jane Doe",
  "email": "jane@example.com"
}
```

#### Delete a user

```bash
DELETE /api/v1/users/:id
```

## Available Make Commands

```bash
make build     # Build the application
make run       # Run the application
make test      # Run tests
make clean     # Clean build artifacts
make fmt       # Format code
make lint      # Run linter (requires golangci-lint)
make deps      # Download dependencies
make dev       # Run in development mode
```

## Development

### Running Tests

```bash
go test -v ./...
```

### Code Formatting

```bash
make fmt
```

### Running with Live Reload

Install `air` for live reloading:

```bash
go install github.com/cosmtrek/air@latest
```

Then run:

```bash
air
```

## Docker Deployment

### Build Docker Image

```bash
docker build -t lemonade:latest .
```

### Run Docker Container

```bash
docker run -p 8080:8080 -e PORT=8080 lemonade:latest
```

### Using Docker Compose

```bash
docker-compose up -d
docker-compose down
```

## Project Customization

### Adding New Endpoints

1. Create a handler in `handlers/` directory
2. Add routes in `routes/routes.go`

Example:

```go
// handlers/product.go
func GetProducts(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"products": []string{}})
}

// In routes/routes.go
productGroup := router.Group("/api/v1/products")
{
    productGroup.GET("", handlers.GetProducts)
}
```

### Adding Middleware

Create middleware in `middleware/` directory and apply it in `main.go`:

```go
router.Use(middleware.YourMiddleware())
```

## Dependencies

- `gin-gonic/gin` - Web framework
- `joho/godotenv` - Environment variable loading

## License

MIT

## Contributing

Feel free to submit issues and enhancement requests!
