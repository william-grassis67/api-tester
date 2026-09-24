# API Tester CLI

A lightweight command-line HTTP client built in Go for developers who need to test and inspect APIs directly from the terminal.

API Tester CLI provides a simple interface for sending HTTP requests without relying on a graphical API client.

## Features

* HTTP methods: GET, POST, PUT, PATCH and DELETE
* JSON request bodies
* Custom HTTP headers
* Query parameters
* HTTP status inspection
* Response headers
* Response body output
* Request execution time
* Lightweight and cross-platform
* Command-line focused workflow

## Installation

### Requirements

* Go 1.26 or later

### Clone the repository

```bash
git clone https://github.com/william-grassis67/api-tester.git
cd api-tester
```

### Build

```bash
go build -o apitest
```

Run the executable:

```bash
./apitest
```

### Install globally

On Linux:

```bash
go build -o apitest
mv apitest ~/bin/
```

If `~/bin` is in your `PATH`, you can use:

```bash
apitest
```

from any directory.

## Usage

### GET

```bash
apitest GET https://api.example.com/users
```

### POST

```bash
apitest POST https://api.example.com/users \
  --body '{"name":"William","email":"william@example.com"}'
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
api-tester/
├── help/
├── methods/
├── main.go
├── go.mod
├── apitest
└── README.md
```

The project structure may evolve as new features are implemented.

## Technology

API Tester CLI is built with:

* Go
* Go Standard Library
* `net/http`
* `encoding/json`

The project intentionally keeps external dependencies to a minimum.

## Development

Run the project directly:

```bash
go run .
```

Run tests:

```bash
go test ./...
```

Build the executable:

```bash
go build -o apitest
```

## Roadmap

### Core

* [x] Project initialization
* [x] GET requests
* [x] POST requests
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
* [ ] Improved terminal output
* [ ] Interactive mode

### Advanced

* [ ] OpenAPI import
* [ ] API test assertions
* [ ] Automated test execution
* [ ] CI/CD integration
* [ ] Export request results
* [ ] Parallel requests

## Design Goals

### Simple

The goal is to make API requests from the terminal without requiring a complicated command syntax.

### Fast

The application is designed to have low overhead and quick startup times.

### Portable

API Tester CLI is written in Go and can be compiled as a standalone executable for different operating systems.

### Developer-focused

The project focuses on practical features that developers commonly need when developing and testing APIs.

## Contributing

Contributions are welcome.

To contribute:

1. Fork the repository.
2. Create a feature branch.

```bash
git checkout -b feature/my-feature
```

3. Implement your changes.
4. Test the project.

```bash
go test ./...
```

5. Commit your changes.

```bash
git commit -m "feat: add my feature"
```

6. Push your branch.

```bash
git push origin feature/my-feature
```

7. Open a Pull Request.

## License

This project is currently under development.

The license will be defined before the first stable release.

## Author

Developed by **William Gabriel Roque de Assis**.

GitHub:
https://github.com/william-grassis67

## Status

API Tester CLI is currently in active development.

The command syntax and available features may change before the first stable release.
