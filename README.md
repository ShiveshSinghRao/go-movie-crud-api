
# Go Movie CRUD API

This is a simple CRUD (Create, Read, Update, Delete) API built using Go (Golang) and the Gorilla Mux router.

## Description

The Go Movie CRUD API provides endpoints to manage a collection of movies. It allows you to perform the following operations:

- Retrieve a list of all movies
- Retrieve details of a specific movie by ID
- Create a new movie
- Update an existing movie by ID
- Delete a movie by ID

## Installation

To run the API locally, follow these steps:

1. Make sure you have Go installed on your system. If not, download and install it from the official website: [https://golang.org/](https://golang.org/)

2. Set up your Go workspace:

```bash
mkdir -p ~/go/src
cd ~/go/src
```

3. Clone the repository:

```bash
git clone https://github.com/ShiveshSinghRao/go-movie-crud-api.git
cd go-movie-crud-api
```

4. Initialize the Go module and install the dependencies:

```bash
go mod init go-movie-crud
go mod tidy
go get -u github.com/gorilla/mux
```

5. Build and run the API:

```bash
go run main.go
```

The API will start running at `http://localhost:8091`.

## API Endpoints

- **GET /movies**: Retrieve a list of all movies.
- **GET /movies/{id}**: Retrieve details of a specific movie by ID.
- **POST /movies**: Create a new movie. (Requires JSON payload)
- **PUT /movies/{id}**: Update an existing movie by ID. (Requires JSON payload)
- **DELETE /movies/{id}**: Delete a movie by ID.

## Request and Response Format

The API expects and returns JSON data. Here's an example of a movie object:

```json
{
    "id": "1",
    "isbn": "4354353",
    "title": "Movie One",
    "director": {
        "firstname": "John",
        "lastname": "Wick"
    }
}
```

## Usage

You can use tools like Postman to interact with the API endpoints. Make sure to set the appropriate HTTP method and request body (if required) for each endpoint.

