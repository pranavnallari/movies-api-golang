# Movie REST API

This Go application implements a basic RESTful API for managing a collection of movies. It provides endpoints to perform CRUD operations (Create, Read, Update, Delete) on movie resources.

## Features

- Fetch all movies
- Fetch a movie by ID
- Create a new movie
- Update a movie by ID
- Delete a movie by ID

## Technologies Used

- Go Programming Language
- Gorilla Mux (HTTP router and dispatcher for Go)
- JSON (for data serialization)

## Usage
Use an HTTP client (e.g., cURL, Postman) to send requests to the API endpoints.
## API Endpoints
GET /movies: Fetch all movies.
GET /movies/{id}: Fetch a movie by ID.
POST /movies: Create a new movie.
PUT /movies/{id}: Update a movie by ID.
DELETE /movies/{id}: Delete a movie by ID.