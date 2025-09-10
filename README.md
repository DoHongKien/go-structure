# Go Structure

A well-structured Go application demonstrating clean architecture principles, modern development practices, and comprehensive feature implementation for building scalable web services.

## 🚀 Features

- **RESTful API Server** - Built with Gin framework
- **Database Integration** - MySQL with GORM ORM
- **Caching** - Redis for high-performance caching
- **Authentication** - JWT-based authentication system
- **API Documentation** - Swagger/OpenAPI integration
- **Logging** - Structured logging with Zap and log rotation
- **Rate Limiting** - Request rate limiting for API protection
- **Dependency Injection** - Google Wire for clean dependency management
- **Configuration Management** - Viper for flexible configuration
- **Containerization** - Docker support for easy deployment
- **Testing** - Comprehensive test suite with benchmarks
- **Tutorials** - Educational Go examples and best practices

## 🛠 Technology Stack

### Core Technologies
- **Go** 1.24+ - Programming language
- **Gin** - HTTP web framework
- **GORM** - ORM library for database operations
- **MySQL** - Primary database
- **Redis** - Caching and session storage

### Development Tools
- **Google Wire** - Dependency injection
- **Viper** - Configuration management
- **Zap** - Structured logging
- **Swagger** - API documentation
- **Docker** - Containerization
- **Lumberjack** - Log rotation

### Additional Libraries
- **JWT** - Authentication tokens
- **Ulule Limiter** - Rate limiting
- **UUID** - Unique identifier generation

## 📁 Project Structure

```
├── cmd/                    # Application entry points
│   ├── cli/               # CLI application
│   └── server/            # HTTP server application
├── config/                # Configuration files
│   ├── local.yaml         # Local development config
│   └── production.yaml    # Production config
├── internal/              # Private application code
│   ├── controller/        # HTTP controllers
│   ├── initialize/        # Application initialization
│   ├── middlewares/       # HTTP middlewares
│   ├── model/            # Data models
│   ├── repo/             # Repository layer
│   ├── routers/          # Route definitions
│   ├── service/          # Business logic
│   ├── utils/            # Internal utilities
│   └── wire/             # Dependency injection setup
├── pkg/                  # Public packages
│   ├── logger/           # Logging utilities
│   ├── response/         # HTTP response helpers
│   ├── setting/          # Configuration settings
│   └── utils/            # Shared utilities
├── storages/             # Storage and logs
├── tests/                # Test files
├── tutorials/            # Educational examples
├── Dockerfile           # Docker configuration
├── Makefile            # Build automation

└── go.mod              # Go module definition
```

## 🚦 Prerequisites

- **Go** 1.24 or higher
- **MySQL** 5.7+ or 8.0+
- **Redis** 6.0+
- **Docker** (optional, for containerized deployment)

## 📦 Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/DoHongKien/go-structure.git
   cd go-structure
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Setup configuration**
   ```bash
   # Copy and modify the configuration file
   cp config/local.yaml config/local.yaml.local
   # Edit the configuration according to your environment
   ```

4. **Setup database**
   ```bash
   # Create MySQL database
   mysql -u root -p -e "CREATE DATABASE ecommerce;"
   ```

5. **Setup Redis**
   ```bash
   # Start Redis server
   redis-server
   ```

## 🏃‍♂️ Running the Application

### Development Mode

```bash
# Run the server
make dev

# Or directly with go run
go run ./cmd/server
```

The server will start on `http://localhost:9999`

### Using Docker

```bash
# Build Docker image
docker build -t go-structure .

# Run with Docker
docker run -p 9999:9999 \
  -e MYSQL_HOST=your_mysql_host \
  -e MYSQL_USER=your_mysql_user \
  -e MYSQL_PASSWORD=your_mysql_password \
  -e MYSQL_DB=ecommerce \
  go-structure
```

### CLI Application

The project includes a CLI application for configuration testing:

```bash
# Run CLI application to test configuration loading
go run ./cmd/cli/viper
```

This will display the current configuration settings from the config files.

## 📚 API Documentation

The API documentation is automatically generated using Swagger and available at:

```
http://localhost:9999/swagger/index.html
```

To regenerate the documentation:

```bash
make swag
```

## 🔧 Development

### Generate Wire Dependencies

```bash
make wire
```

### Available Make Commands

```bash
make dev         # Run development server
make swag        # Generate Swagger documentation  

make wire        # Generate dependency injection code
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific test files
go test ./tests/...

# Run benchmarks
go test -bench=. ./tests/
```

## 🎯 Key Components

### Models
- **Customer** - User management
- **Order** - Order processing
- **OrderDetail** - Order line items

### API Endpoints
- Authentication endpoints
- Customer management
- Order processing
- User operations

### Middleware
- Authentication middleware
- Rate limiting
- CORS handling
- Request logging

## 🎓 Learning Resources

The project includes comprehensive tutorials in the `/tutorials` directory:

- **Goroutines** - Concurrent programming examples
- **API Calls** - HTTP client implementations
- **File I/O** - File read/write operations
- **JSON Processing** - Marshal/unmarshal examples

## 🔒 Environment Variables

Key environment variables for configuration:

```bash
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=yourpassword
MYSQL_DB=ecommerce
REDIS_HOST=localhost
REDIS_PORT=6379
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and conventions
- Write tests for new functionality
- Update documentation for API changes
- Ensure all tests pass before submitting PR
- Use meaningful commit messages

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **DoHongKien** - *Initial work* - [DoHongKien](https://github.com/DoHongKien)

## 🙏 Acknowledgments

- Clean Architecture principles
- Go community best practices
- Modern web development patterns
- Open source contributors

## 📞 Support

For support and questions:

- Create an issue in the repository
- Check the documentation
- Review the tutorial examples

---

*This project serves as a reference implementation for Go web applications with clean architecture, comprehensive features, and modern development practices.*