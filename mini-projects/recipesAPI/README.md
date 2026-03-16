# Recipe Management REST API

A RESTful API built with Go and the Gin framework for managing recipes. This project demonstrates core backend concepts including database persistence, caching, caching, authentication, and structured logging.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete recipes.
- **Database**: MongoDB for persistent storage of recipes and user profiles.
- **Caching**: Redis implementation for faster read operations (e.g., fetching lists of recipes).
- **Authentication**: Custom JWT (JSON Web Token) authentication flow (Signup & Signin).
- **Middleware**: Custom Gin middleware for protecting private routes.
- **Dockerized**: Easy local setup using Docker Compose for MongoDB and Redis.

## Tech Stack

- **Language**: Go (Golang)
- **Web Framework**: [Gin](https://gin-gonic.com/)
- **Database**: MongoDB (via `go.mongodb.org/mongo-driver/v2`)
- **Cache**: Redis (via `github.com/redis/go-redis/v9`)
- **Authentication**: JWT (via `github.com/golang-jwt/jwt/v5`)
- **Containerization**: Docker & Docker Compose

## Prerequisites

- Go (1.20+)
- Docker and Docker Compose (for running MongoDB and Redis locally)

## Getting Started

### 1. Start Infrastructure
Start the MongoDB and Redis containers using the provided Docker Compose file:
```bash
docker-compose up -d
```

### 2. Run the Application
Start the Gin server:
- setup environment variables
  - MONGODB_URI
    - `mongodb://localhost:27017`
  - JWT_SECRET
```bash
go run main.go
```
The server will start on `http://localhost:8088`.

## API Endpoints

### Authentication
- `POST /signup` - Register a new user
- `POST /signin` - Authenticate and receive a JWT

### Recipes (Public)
- `GET /recipes` - List all recipes (Cached via Redis)
- `GET /recipes/search?tag=...` - Search recipes by tags
- `GET /recipes/:id` - Get a specific recipe

### Recipes (Protected - Requires JWT)
*Requires `Authorization: <token>` header.*
- `POST /recipes` - Create a new recipe
- `PUT /recipes/:id` - Update an existing recipe
- `DELETE /recipes/:id` - Delete a recipe

## Documentation
For deeper dives into the concepts learned while building this project, see [`learnGin.md`](./learnGin.md).
