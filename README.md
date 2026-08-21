# URL Shortening Service

A lightweight, RESTful URL shortening service built with Go. This project implements a complete URL shortening API that allows users to create, retrieve, update, and delete shortened URLs.

## 📋 Project Overview

This project is based on the [URL Shortening Service roadmap](https://roadmap.sh/projects/url-shortening-service) from roadmap.sh. It provides a practical implementation of a web service that transforms long URLs into short, easy-to-share links.

### Key Features

- ✅ **Create Shortened URLs** - Convert long URLs into short, random 8-character codes
- ✅ **Retrieve URLs** - Look up URLs by their short code
- ✅ **Update URLs** - Modify the target URL of an existing shortened link
- ✅ **Delete URLs** - Remove shortened URLs from the system
- ✅ **List All URLs** - Retrieve all shortened URLs
- ✅ **Automatic Timestamps** - Track creation and last modification time for each URL
- ✅ **Error Handling** - Comprehensive HTTP status codes and error messages
- ✅ **Request Logging** - Middleware logging for all incoming requests

## 🛠 Technology Stack

- **Language**: Go 1.23.6
- **HTTP Server**: Go's built-in `net/http` package
- **Data Format**: JSON
- **Storage**: In-memory (slice-based)

## 📁 Project Structure

```
url-shortener/
├── go.mod                    # Go module definition
├── url-shortener.go         # HTTP handlers and routing
├── url-repository.go        # Data model and repository functions
└── README.md                # This file
```

### File Descriptions

- **`go.mod`** - Go module file specifying the project name and Go version
- **`url-shortener.go`** - Contains HTTP request handlers, middleware, and server setup
- **`url-repository.go`** - Implements the data model (`ShortURL` struct) and all CRUD operations with error handling

## 🚀 Getting Started

### Prerequisites

- Go 1.23.6 or later installed on your system
- Basic knowledge of REST APIs and HTTP methods

### Installation

1. Clone or navigate to the project directory:
   ```bash
   cd url-shortener
   ```

2. Download dependencies (if any):
   ```bash
   go mod download
   ```

3. Build the project:
   ```bash
   go build
   ```

4. Run the service:
   ```bash
   go run .
   ```

The API server will start and be ready to accept requests.

## 📡 API Endpoints

### Create a Shortened URL
**POST** `/create`

Request body:
```json
{
  "url": "https://example.com/very/long/url/that/needs/shortening"
}
```

Response (201 Created):
```json
{
  "id": "0",
  "longUrl": "https://example.com/very/long/url/that/needs/shortening",
  "shortUrl": "abc12345",
  "createdAt": "2024-01-15T10:30:00Z",
  "updatedAt": "2024-01-15T10:30:00Z"
}
```

### Retrieve a URL
**GET** `/get/{shortUrl}`

Response (200 OK):
```json
{
  "id": "0",
  "longUrl": "https://example.com/very/long/url/that/needs/shortening",
  "shortUrl": "abc12345",
  "createdAt": "2024-01-15T10:30:00Z",
  "updatedAt": "2024-01-15T10:30:00Z"
}
```

### Update a URL
**PUT** `/update/{shortUrl}`

Request body:
```json
{
  "url": "https://example.com/new/url"
}
```

Response (200 OK):
```json
{
  "id": "0",
  "longUrl": "https://example.com/new/url",
  "shortUrl": "abc12345",
  "createdAt": "2024-01-15T10:30:00Z",
  "updatedAt": "2024-01-15T10:35:00Z"
}
```

### Delete a URL
**DELETE** `/delete/{shortUrl}`

Response (204 No Content)

### List All URLs
**GET** `/all`

Response (200 OK):
```json
[
  {
    "id": "0",
    "longUrl": "https://example.com/very/long/url",
    "shortUrl": "abc12345",
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-01-15T10:30:00Z"
  },
  ...
]
```

## 📊 Data Model

### ShortURL Struct

```go
type ShortURL struct {
    Id        string  // Unique identifier (index-based)
    LongUrl   string  // The original long URL
    ShortUrl  string  // Generated 8-character short code
    CreatedAt string  // ISO 8601 timestamp of creation
    UpdatedAt string  // ISO 8601 timestamp of last update
}
```

## ⚙️ Implementation Details

### Short URL Generation

- **Length**: 8 characters
- **Character Set**: Alphanumeric (a-z, A-Z, 0-9)
- **Algorithm**: Cryptographically seeded random string generation
- **Uniqueness**: Automatically handles duplicates by storing all URLs in a slice

### Error Handling

The service returns appropriate HTTP status codes:

| Status Code | Scenario |
|------------|----------|
| 200 OK | Successful GET, PUT operations |
| 201 Created | Successful POST operation |
| 204 No Content | Successful DELETE operation |
| 400 Bad Request | Invalid request body, URL not found |
| 409 Conflict | URL already exists |

### Middleware

- **Logging Middleware**: Logs all incoming HTTP requests with method and path

## 🔄 Current State

- ✅ In-memory storage (data persists only during runtime)
- ✅ JSON-based API
- ✅ Full CRUD operations
- ✅ Request validation and error handling

## 🚧 Future Enhancements

- Persistent storage (Database integration - PostgreSQL, MongoDB, etc.)
- URL validation and normalization
- Custom short code support
- Expiration/TTL for shortened URLs
- Analytics (click tracking, statistics)
- Rate limiting and authentication
- URL preview functionality
- Batch operations
- Redirect functionality (actually redirect to long URL)

## 📝 Example Usage

### Using cURL

```bash
# Create a short URL
curl -X POST http://localhost:8080/create \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.google.com/search?q=golang"}'

# Get a URL
curl http://localhost:8080/get/abc12345

# Update a URL
curl -X PUT http://localhost:8080/update/abc12345 \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.github.com"}'

# Delete a URL
curl -X DELETE http://localhost:8080/delete/abc12345

# List all URLs
curl http://localhost:8080/all
```

## 📚 Related Resources

- [URL Shortening Service - roadmap.sh](https://roadmap.sh/projects/url-shortening-service)
- [Go Documentation](https://golang.org/doc)
- [Go net/http Package](https://golang.org/pkg/net/http)
- [REST API Best Practices](https://restfulapi.net)

## 📄 License

This project is open source and available under the MIT License.

## 🤝 Contributing

Contributions are welcome! Feel free to fork this repository and submit pull requests for any improvements or bug fixes.

## ⚡ Quick Start

```bash
# Build
go build

# Run
./url-shortener

# The server will be available at http://localhost:8080
```

---

Built as a practical implementation of the URL Shortening Service project specifications from [roadmap.sh](https://roadmap.sh/projects/url-shortening-service).
