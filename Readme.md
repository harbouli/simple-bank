
# Simple Bancking Application with GO


BankApp is a simple banking application developed in Go, designed to demonstrate basic CRUD operations with an emphasis on maintaining ACID properties. It showcases best practices in application development, including unit testing, continuous integration with GitHub Actions, and containerization with Docker for easy database setup and migration.

## Features

- **CRUD Operations**: Manage bank accounts, including creating, reading, updating, and deleting account information.
- **ACID Properties**: Ensures database transactions are processed reliably, maintaining the integrity of the database.
- **Unit Testing**: Comprehensive unit tests to ensure application logic is correct and robust.
- **GitHub Actions**: Automated workflows for continuous integration, running tests, and ensuring code quality.
- **Docker Integration**: Simplifies database creation and migration, ensuring a consistent environment for development and production.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.x or later)
- [Docker](https://www.docker.com/get-started)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Installation

Clone the repository:

```
git clone https://github.com/harbouli/simple-bank.git
```

```
cd simple-bank
```

Install dependencies:

```
go mod tidy
```

### Development Setup

Use the Makefile commands to manage the PostgreSQL database via Docker, handle database migrations, and generate database code.

```
make postgres # Start the PostgreSQL container
make createdb # Create the database
make dropdb # Drop the database
make migrateup # Apply database migrations
make migratedown # Revert database migrations
make sqlc # Generate SQL code
```

## Usage

Provide examples of how to use your application, including any CLI commands or endpoints to test the CRUD operations.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any bugs or feature requests.


