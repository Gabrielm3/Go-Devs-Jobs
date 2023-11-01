# Go Devs Jobs

Go Devs Jobs is an employment opportunities API developed in Golang. This API employs Go-Gin as its router, SQLite as its database, and GoORM for database interaction. It also includes Swagger for documentation and testing, and boasts a well-organized package structure.

## Project Structure

The project is structured as follows:

- `config`: Contains project configuration.
- `docs`: Swagger documentation.
- `handler`: API controllers.
- `router`: Router configuration.
- `schemas`: Data structure definitions.
- `go.mod` and `go.sum`: Files for managing Go dependencies.
- `main.go`: The entry point of the application.
- `makefile`: File for automating common tasks.

## Prerequisites

Before getting started, ensure that you have Go installed on your system.
Additionally, configure SQLite for data storage.

## How to Set Up and Run

Follow these steps to set up and run the API:

1. Clone the repository: `git clone https://github.com/Gabrielm3/Go-Devs-Jobs.git`
2. Navigate to the project folder: `cd Go-Devs-Jobs`
3. Configure SQLite: Verify that SQLite is correctly configured. Check the settings in `config/config.go` and `config/sqlite.go`.
4. Compile the code: `go build`
5. Execute the application: `./Go-Devs-Jobs`

The API will run at `http://localhost:8080`.

## API Documentation

API documentation is available via Swagger. You can access it at `http://localhost:8080/swagger/index.html`.

## API Endpoints

- `GET /api/v1/opening`: Retrieve a job opening.
- `POST /api/v1/opening`: Create a new job opening.
- `DELETE /api/v1/opening`: Delete a job opening.
- `PUT /api/v1/opening`: Update a job opening.
- `GET /api/v1/openings`: List all job openings.

## Contributing

Contributions are welcome! If you encounter a bug or have an improvement, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. Refer to the LICENSE file for details.
