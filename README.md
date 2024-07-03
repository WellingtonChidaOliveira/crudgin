# Go-Gin Application with PostgreSQL

This project is a Go application using the Gin framework, connected to a PostgreSQL database. It's structured to separate concerns between server logic and database interactions, making it scalable and easy to maintain.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (see `go.mod` for the required version)
- Docker and Docker Compose
- A `.env` file configured with your database credentials

### Installing

1. Clone the repository to your local machine.
2. Ensure you have Docker installed and running.
3. Create a `.env` file in the root directory with the following variables:
    ```
    DB_USERNAME=your_username
    DB_PASSWORD=your_password
    DB_DATABASE=your_database
    ```
4. Build the Go application:
    ```
    make build
    ```
5. Start the PostgreSQL container:
    ```
    make db-up
    ```

### Running the Application

After installation, you can run the application using:

```
make run
```

This will start the Go application located in the `bin/` directory.


### Stopping the Application

To stop the application and the PostgreSQL container, you can use:

```
make db-down
```

## Structure

The project is organized as follows:

- `.env`: Environment variables for database configuration.
- `bin/`: Contains the compiled Go application.
- `docker-compose.yml`: Docker Compose configuration for running PostgreSQL.
- `go.mod` and `go.sum`: Go module files for managing dependencies.
- `internals/`: Contains the core logic of the application, separated into `database/` for database interactions and `server/` for server and routing logic.
- `main.go`: The entry point of the application.
- `Makefile`: Contains commands for building, running, and managing the Docker containers.
- `README.md`: This file.
- `requests.http`: A collection of HTTP requests for testing the API.
- `types/`: Contains Go structs and types used across the application.

## Contributing

Please read `README.md` for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

- **Wellington Chida** - *WC*


