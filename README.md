# go_jobs
## About
This project is a study case of building an API, using Golang, for job opportunities and openings; The API is powered by Go-Gin as a router, GoORM for database communication, SQLite as the database, and Swagger for API documentation and testing. The project follows a modern package structure to keep the codebase organized and maintainable.

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [SQLite](https://www.sqlite.org/index.html) as the database
- [Swagger](https://swagger.io/) for API documentation and testing

## Instalation
To use this project, you need to follow these steps:
1. Clone the repository: `git clone https://github.com/ThiagoVicentini/go_jobs.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main <port>` (if port is not specified, it'll be set to $PORT, which is 8080 by default) 

## Usage

After the API is running, you can use the Swagger UI to interact with the endpoints for searching, creating, editing, and deleting job opportunities. The API can be accessed at `http://localhost:$PORT/swagger/index.html`.

Default $PORT if not provided=8080.