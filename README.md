
Kodu kopyala
# User Service - Go + HTTP API

This project is a simple user management microservice written in Go. It allows creating, updating, deleting, and listing users with pagination and filtering options. The service exposes a RESTful HTTP API and uses in-memory data storage.

## Features
- **Add a New User**
- **Update an Existing User**
- **Delete a User**
- **List Users** with pagination and filtering based on fields like `first_name`, `last_name`, `email`, `nickname`, `country`.

## Technologies Used
- **Go**: The primary language used for this project.
- **Gorilla Mux**: For HTTP request routing.
- **bcrypt**: For password hashing and verification.
- **Docker**: For containerization of the application.

## Installation

### Prerequisites
- [Go](https://golang.org/doc/install) 1.18 or above
- [Docker](https://docs.docker.com/get-docker/) (if using Docker)
  
### Clone the repository
```bash
git clone https://github.com/yourusername/user-service.git
cd user-service
```

### Build the Application
```bash
make build
```

This will compile the Go code and generate the user-service binary.

### Run the Application

You can run the application directly:

```bash
./user-service
```

By default, the server will run on http://localhost:8080

## API Endpoints
1. Create new user
- URL: POST /users
- Body:
```json
{
  "first_name": "Alice",
  "last_name": "Bob",
  "nickname": "AB123",
  "email": "alice@bob.com",
  "password": "supersecurepassword",
  "country": "UK"
}
```
- Response:
```json
{
  "id": "generated-uuid",
  "first_name": "Alice",
  "last_name": "Bob",
  "nickname": "AB123",
  "email": "alice@bob.com",
  "country": "UK"
}
```

2. Update user
- URL: PUT /users/{id}
- Body:
```json
{
  "first_name": "AliceUpdated",
  "last_name": "BobUpdated",
  "nickname": "AB123Updated",
  "email": "alice.updated@bob.com",
  "country": "US"
}
```
- Responser: Updated user data

3. Delete user
- URL: DELETE /users/{id}

4. List users with pagination and filtering
- URL: GET /users
- Optional Query Parameters:
* first_name
* last_name
* nickname
* email
* country
* limit (default:10)
* offset (default: 0)
- Response: A list of users matching the filter criteria

## Run with Docker

You can also buld and run the service with Docker

### Build the Docker Image
```bash
make docker-build
```

### Run the Docker Conatiner
```bash
make run
```

This will run the service inside a Docker container and expose it on port 8080. 

You can access the API at http://localhost:8080

## Testing

To run the tests, use the following command:
```bash
make test
```

This will execute the Go unit tests that are written for the service, including service, respository, handler and utility tests.

## Licence

This project is licensed under the MIT Licence.