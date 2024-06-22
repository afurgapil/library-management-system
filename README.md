## Go Library Management System

This project demonstrates the basic Go project structure and CRUD operations using the Go programming language.

## Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Test](#test)
6. [Contributing](#contributing)
7. [License](#license)

## Introduction

The Go Library Management System is designed to showcase a simple library management system implementing CRUD operations with a structured Go project. This project aims to help developers understand how to structure a Go(Fiber) application and perform basic database operations with Postgresql.

## Features

- Basic CRUD operations for managing library records
- Unit and Integration Tests
- CI/CD using Github Actions
- API documentation with Swagger

## Installation

1. Clone this project: `git clone https://github.com/afurgapil/library-management-system.git`
2. Navigate to the project directory: `cd library-management-system`
3. Create the database table on your PostgreSQL server by running the `migrations` file.
4. Install dependencies: `go mod tidy`
5. Check `.env` file

## Usage

1. Run the server: `go run cmd/library-management-system/main.go`
2. Access the API using an HTTP client. You can check the documentation on swagger.

## Test

To run all tests `go test ./... -v`

## Contributing

If you encounter any issues or have suggestions for improvements, please feel free to contribute. Your feedback is highly appreciated and contributes to the learning experience.

## License

This project is licensed under the [MIT License](LICENSE). For more information, please refer to the LICENSE file.
