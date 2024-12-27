# MyApp - Go Project

This is a simple Go application that demonstrates the use of modular architecture with controllers, services, repositories, and utilities.

## Features

- **Gin Framework** for routing and HTTP request handling.
- Modular architecture with clean separation of concerns.
- Custom logger that writes logs to both the console and a file (`logger.log`).
- A simple example with an `Album` resource, including CRD operations.
- Example of trait-like behavior using Go's struct embedding.
- That App just use data from Array at repository not database. Next i will create use database

## Project Structure

```
myapp/
├── app/
│   ├── controllers/        # Controller for handling requests
│   │   └── albumController.go
│   ├── repositories/       # Data management (could be a database or in-memory)
│   │   └── albumRepository.go
│   ├── services/           # Business logic layer
│   │   └── albumService.go
├── config/                 # Configuration and environment variables
│   └── config.go
├── helpers/                # Helper functions for responses and utilities
│   └── responseHelper.go
├── routes/                 # Define and initialize routes
│   ├── indexRoute.go
│   └── albumRoute.go
├── utils/                  # Utility functions (e.g., logging)
│   └── logger.go
├── main.go                 # Entry point for the application
```

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Awasefra/golang-gin-restful-api-album.git
   ```

2. **Install dependencies**:

   You need to install [Go](https://golang.org/doc/install) if it's not installed already.

   ```bash
   go mod get
   ```

3. **Set up environment variables**:

   Create a `.env` file in the root directory:

   ```
   PORT=8080
   ```

4. **Run the application**:

   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`.

## Endpoints
 u can test api using postman collection, just import: ALBUMREST.postman_collection.json
- `GET /api/albums`: Get a list of all albums.
- `GET /api/albums/:id`: Get a single album by ID.
- `POST /api/albums`: Create a new album.

to cek health ( app is run or no):  `GET /health`

## Logging

The application logs both to the console and to a file called `logger.log` in the root directory. The logger outputs timestamps and log levels for easier debugging.

## Contributing

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes.
4. Push to your fork and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
