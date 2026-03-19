# Recipe Management REST API

A RESTful recipe platform built with Go + Gin and a React frontend. The project demonstrates a production-style architecture with persistence, caching, full-text search, JWT/Cognito authentication flow, and structured logging.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete recipes.
- **Database**: MongoDB as source of truth for recipe data.
- **Caching**: Redis for faster read operations (`recipes` and `recipe:<id>` cache keys).
- **Search**: Elasticsearch-backed search endpoint for recipe name/tags.
- **Authentication**: JWT validation with AWS Cognito JWKS.
- **Structured Logging**: Zap logger with console + file output.
- **Frontend UI**: React app for browsing recipes and searching from the UI.
- **Dockerized Infra**: Easy local setup using Docker Compose for DB, Cache, Search and UI.

Given the data is JSON, this maybe primary reason why I have picked above services. 

## Tech Stack

- **Language**: Go (Golang)
- **Web Framework**: [Gin](https://gin-gonic.com/)
- **DB**: MongoDB (via `go.mongodb.org/mongo-driver/v2`)
- **Cache**: Redis (via `github.com/redis/go-redis/v9`)
- **Search**: Elasticsearch (via `github.com/elastic/go-elasticsearch/v9`)
- **Auth**: JWT + AWS Cognito (JWKS based validation)
- **UI**: React (`recipes-web`)
- **Logging**: Zap (`go.uber.org/zap`)
- **Containerization**: Docker & Docker Compose

## Prerequisites

- Go (1.25+)
- Node.js + npm (for React UI)
- Docker and Docker Compose (for running MongoDB and Redis locally)
- Elasticsearch running locally (or remote URI)

## Environment Variables

Set these before running the backend:

```bash
MONGODB_URI=mongodb://localhost:27017
ELASTICSEARCH_URI=http://localhost:9200

AWS_REGION=ap-south-1
AWS_USER_POOL_ID=ap-south-1_XXXXXXXXX
AWS_CLIENT_ID=xxxxxxxxxxxxxxxxxxxxxxxxxx
AWS_ISSUER=https://cognito-idp.ap-south-1.amazonaws.com/ap-south-1_XXXXXXXXX
```

## Getting Started

### 1. Start Infrastructure
Start MongoDB, Elasticsearch, Kibana and Redis containers:
```bash
docker-compose up -d
```
All infra runs on default ports:
- MongoDB: 27017
- Elasticsearch: 9200
- Kibana: 5601
- Redis: 6379

If you have anything configured, just change it in config `docker-compose.yaml`


### 2. Start Elasticsearch
Run Elasticsearch locally (for example using Docker):

```bash
docker run -d --name elasticsearch \
  -p 9200:9200 -p 9300:9300 \
  -e "discovery.type=single-node" \
  -e "xpack.security.enabled=false" \
  docker.elastic.co/elasticsearch/elasticsearch:8.15.1
```

### 3. Run Backend API
```bash
go run main.go
```

Backend runs on `http://localhost:8088`.

### 4. Run React UI
In a separate terminal:

```bash
cd recipes-web
npm start
```

React app runs on `http://localhost:3000` (default port). Change it if you want to.

## API Endpoints

### Recipes (Public)
- `GET /recipes` - List all recipes (Cached via Redis)
- `GET /recipes/search?q=...` - Search recipes by name/tags in Elasticsearch
- `GET /recipes/search?tag=...` - Exact tag filter in Elasticsearch
- `GET /recipe/:id` - Get one recipe by ID

### Recipes (Write APIs)
- `POST /recipe` - Create a new recipe
- `PATCH /recipe/:id` - Update an existing recipe
- `DELETE /recipe/:id` - Delete a recipe

### Auth Middleware
- Cognito JWT middleware is implemented and can protect write routes.
- Test, initialise the middleware by uncommenting the following line in `main.go`:

```go
engine.Use(authHandler.AuthMiddleware(jwks, issuer, clientID))
```

Then call protected APIs with:

```http
Authorization: Bearer <access_token>
```

## Handy Notes

- MongoDB is the source of truth.
- Redis caches list and item reads.
- Elasticsearch is used for search (`/recipes/search`).
- On create/update/delete, cache is invalidated and Elasticsearch index is synced.
- React app consumes backend APIs for browse + search UX.
