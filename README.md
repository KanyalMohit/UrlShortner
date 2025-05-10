# URL Shortener Service

A simple, efficient URL shortener service built with Go. This service allows you to create short URLs from long ones, manage them, and redirect users to the original URLs.

## Features

- Create short URLs from long URLs
- Redirect to original URLs
- Get information about shortened URLs
- Delete shortened URLs
- SQLite database for storage
- RESTful API
- Secure URL generation
- Environment-based configuration

## Prerequisites

- Go 1.16 or higher
- SQLite3

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd urlshortener
```

2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables (optional):
Create a `.env` file in the root directory:
```env
# Server Configuration
SERVER_PORT=8080
SERVER_HOST=localhost
SERVER_READ_TIMEOUT=10s
SERVER_WRITE_TIMEOUT=10s

# Database Configuration
DB_PATH=./urlshortener.db

# URL Configuration
URL_SHORT_CODE_LENGTH=6
URL_BASE_URL=http://localhost:8080
```

## Running the Application

1. Start the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080` by default.

## API Usage

### 1. Create Short URL
```bash
curl -X POST http://localhost:8080/api/urls \
  -H "Content-Type: application/json" \
  -d '{"original_url": "https://www.example.com/very/long/url"}'
```

### 2. Get URL Information
```bash
curl http://localhost:8080/api/urls/{shortCode}
```

### 3. Delete URL
```bash
curl -X DELETE http://localhost:8080/api/urls/{shortCode}
```

### 4. Redirect to Original URL
Visit in browser:
```
http://localhost:8080/{shortCode}
```

For detailed API documentation, see [API.md](API.md).

## Project Structure

```
urlshortener/
├── api/           # API handlers and routes
├── config/        # Configuration management
├── database/      # Database operations
├── models/        # Data models
├── utils/         # URL generation utilities
├── .env           # Environment variables
├── .gitignore     # Git ignore rules
├── API.md         # API documentation
├── go.mod         # Go module file
├── go.sum         # Go dependencies checksum
└── main.go        # Main application entry point
```

## Key Components

1. **URL Generation**
   - Secure random generation
   - Configurable length
   - URL-safe characters

2. **Database**
   - SQLite for simplicity
   - Indexed for performance
   - Proper schema design

3. **API**
   - RESTful design
   - Proper error handling
   - Input validation

4. **Configuration**
   - Environment-based
   - Type-safe
   - Default values

## Development

### Running Tests
```bash
go test ./...
```

### Building
```bash
go build
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

[Mohit Kanyal]

## Acknowledgments

- [gorilla/mux](https://github.com/gorilla/mux) for routing
- [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) for SQLite support 