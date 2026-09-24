# API Tester CLI

A lightweight command-line HTTP client for developers who want to test, inspect, and debug APIs directly from the terminal.

API Tester CLI is being built with Go, with a focus on simplicity, performance, portability, and a clean developer experience.

## Features

* Send HTTP requests directly from the terminal
* Support for GET, POST, PUT, PATCH, and DELETE
* Custom HTTP headers
* Query parameters
* JSON request bodies
* HTTP status inspection
* Response headers
* Response body formatting
* Request execution time
* Request history
* Environment variables
* Cross-platform executable

## Installation

### From source

Clone the repository:

```bash
git clone https://github.com/william-grassis67/apitest.git
```

Enter the project directory:

```bash
cd apitest
```

Build the application:

```bash
go build -o apitest
```

Run:

```bash
./apitest
```

## Usage

### GET request

```bash
apitest GET https://api.example.com/users
```

### POST request

```bash
apitest POST https://api.example.com/users \
  -b '{"name":"William","email":"william@example.com"}'
```

### Custom headers

```bash
apitest GET https://api.example.com/users \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Query parameters

```bash
apitest GET "https://api.example.com/users?page=1&limit=10"
```

## Example

```text
$ apitest GET https://api.example.com/users

Request
GET https://api.example.com/users

Response
Status: 200 OK
Time: 142ms
Size: 1.4 KB

Body:
[
  {
    "id": 1,
    "name": "William"
  }
]
```

## Project Structure

```text
apitest/
├── cmd/
├── internal/
│   ├── httpclient/
│   └── output/
├── main.go
├── go.mod
├── go.sum
└── README.md
```

The project structure may evolve as new features are introduced.

## Technology

API Tester CLI is built with:

* Go
* Go Standard Library
* net/http
* encoding/json

Additional dependencies may be introduced when they provide clear value to the project.

## Development

Clone the repository:

```bash
git clone https://github.com/william-grassis67/apitest.git
cd apitest
```

Run the project:

```bash
go run .
```

Run tests:

```bash
go test ./...
```

Build:

```bash
go build -o apitest
```

## Roadmap

### Core

* [x] Project initialization
* [ ] GET requests
* [ ] POST requests
* [ ] PUT requests
* [ ] PATCH requests
* [ ] DELETE requests
* [ ] Custom headers
* [ ] Query parameters
* [ ] JSON request body
* [ ] Response formatting
* [ ] Error handling

### Developer Experience

* [ ] Request history
* [ ] Saved requests
* [ ] Environment variables
* [ ] Collections
* [ ] Configuration file
* [ ] Colored terminal output
* [ ] Interactive mode

### Advanced

* [ ] Import OpenAPI specifications
* [ ] API test assertions
* [ ] Automated test execution
* [ ] CI/CD support
* [ ] Export request results
* [ ] Parallel requests

## Design Goals

The project follows a few principles:

### Simple

A developer should be able to execute an HTTP request without learning a complicated command syntax.

### Fast

The tool should have minimal startup time and low resource consumption.

### Portable

The application should work across major operating systems without requiring a runtime environment.

### Developer-focused

Features should solve real problems encountered when developing and testing APIs.

## Contributing

Contributions are welcome.

To contribute:

1. Fork the repository.
2. Create a branch for your change.
3. Implement and test your changes.
4. Commit your changes.
5. Open a Pull Request.

Example:

```bash
git checkout -b feature/request-history
```

Run the test suite before submitting:

```bash
go test ./...
```

## License

This project is currently under development.

The license will be defined before the first stable release.

## Author

Developed by William Gabriel Roque de Assis.

GitHub:

https://github.com/william-grassis67

## Status

API Tester CLI is currently in active development.

The API and command syntax may change before the first stable release.
