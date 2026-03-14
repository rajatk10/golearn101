# REST API - User Management System

A production-ready REST API built with Go for managing user data with role-based access control. The API uses BoltDB for persistent file-based storage and provides comprehensive CRUD operations with full HTTP method support.

## Project Overview

This REST API implements a user management system with the following features:

- **Complete CRUD Operations** – Create, Read, Update, Delete users
- **Partial Updates** – PATCH endpoint for selective field updates
- **HTTP Methods** – GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS
- **Persistent Storage** – BoltDB for reliable file-based data persistence
- **Comprehensive Testing** – Unit tests and performance benchmarks
- **Error Handling** – Proper HTTP status codes and JSON error responses

## Technology Stack

- **Language** – Go 1.26
- **Database** – BoltDB (via Storm ORM)
- **Testing** – Go's built-in testing framework
- **Port** – localhost:8086

## Project Structure

```
├── README.md                          # This file
├── go.mod                             # Go module definition
├── data/                              # Database files (auto-created)
│   └── users.db                       # Production database
└── goHttpRestApi/
    ├── main.go                        # Server entry point
    ├── user/
    │   ├── user.go                    # User model and database operations
    │   └── user_test.go               # Unit tests and benchmarks
    └── handlers/
        ├── rootHandler.go             # Root endpoint handler
        ├── usersRouter.go             # Request routing logic
        ├── usersHandler.go            # CRUD operation handlers
        ├── responses.go               # Response formatting utilities
        ├── usersHandler_test.go       # Handler unit tests
```

## API Endpoints

### User Collection

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users` | List all users |
| POST | `/users` | Create a new user |
| HEAD | `/users` | Check if users exist (no body) |
| OPTIONS | `/users` | Get allowed methods |

### Individual User

| GET, HEAD, POST, PUT, PATCH, DELETE | `/users/{id}` |

## User Model

```json
{
  "id": "507f1f77bcf86cd799439011",
  "name": "John Doe",
  "role": "Engineer",
  "location": "San Francisco",
  "yearsOfExperience": 5,
  "individualContributor": true,
  "managerial": false,
  "executive": false
}
```

### Fields
- Required
  - `name` – User's full name (required)
  - `role` – User's role/position (required) 
- Optional Fields
  - `location` – User's location
  - `yearsOfExperience` – Years of work experience
  - `individualContributor` – IC flag
  - `managerial` – Manager flag
  - `executive` – Executive flag

## Getting Started
- Start the server `go run goHttpRestApi/main.go`
- 
### Example Requests
Either use **curl** or **postman**

**Create a user**
```bash
curl -X POST http://localhost:8086/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Wick",
    "role": "Engineer",
    "location": "Hyderabad",
    "yearsOfExperience": 6,
    "individualContributor": true
  }'
```

**Get all users**
```bash
curl http://localhost:8086/users
```

**Get user by ID**
```bash
curl http://localhost:8086/users/507f1f77bcf86cd799439011
```

**Update user (full replacement)**
```bash
curl -X PUT http://localhost:8086/users/507f1f77bcf86cd799439011 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Rajat Kumar",
    "role": "Senior Engineer",
    "location": "Bangalore",
    "yearsOfExperience": 7,
    "individualContributor": true
  }'
```

**Partial update (PATCH)**
```bash
curl -X PATCH http://localhost:8086/users/507f1f77bcf86cd799439011 \
  -H "Content-Type: application/json" \
  -d '{
    "yearsOfExperience": 7,
    "location": "Bangalore"
  }'
```

**Delete user**
```bash
curl -X DELETE http://localhost:8086/users/507f1f77bcf86cd799439011
```

## Testing

### Run Unit Tests

```bash
# Test user package
go test ./reiley/user -v

# Test handlers package
go test ./reiley/handlers -v

# Test all packages
go test ./...
```

### Run Benchmarks

```bash
# Run all benchmarks
go test -bench .user

# Run specific benchmark
go test -bench BenchmarkCreate .user

# Run with verbose output
go test -bench .user -v

# Running benchmark easiest cd to user directory and run
cd user
go test -bench . -v 

```

## Performance Benchmarks

Benchmarks were run on **Apple M Series** with BoltDB file-based storage.

### Results Summary

| Operation | Iterations | Time/Op | Notes |
|-----------|-----------|---------|-------|
| **BenchmarkCRUD** | 39 | 27.95ms | Full cycle: Create → Read → Update → Read → Delete |
| **BenchmarkCreate** | 121 | 9.16ms | Single user creation with disk I/O |
| **BenchmarkRead** | 3242 | 0.46ms | Single user lookup (20x faster than writes) |
| **BenchmarkDelete** | 123 | 9.05ms | Single user deletion with disk I/O |
| **BenchmarkUpdate** | 132 | 9.09ms | Single user update with disk I/O |

### Performance Analysis

- **Write Operations** (~9ms) – Create, Update, Delete all involve disk I/O to BoltDB
- **Read Operations** (~0.46ms) – Fast in-memory lookups with Storm indexing
- **Combined Operations** (~28ms) – CRUD cycle takes ~3x a single write operation

### Key Insights

1. **Write Bottleneck** – All mutations take ~9ms due to BoltDB disk I/O
2. **Read Efficiency** – 0.46ms demonstrates excellent indexing performance
3. **Predictable Performance** – Consistent across all write operations
4. **Scalability** – For higher throughput, consider Redis or batch operations

## Database

### Storage

- **Type** – BoltDB (embedded key-value store)
- **Location** – `data/users.db` (auto-created)
- **Format** – Binary (not human-readable)
- **Transactions** – ACID-compliant with full transaction support

### Test Database

- **Location** – `data/users_test.db` (isolated from production)
- **Lifecycle** – Created fresh for each test run, cleaned up after completion
- **Purpose** – Ensures test isolation and prevents data corruption

## Error Handling

The API returns appropriate HTTP status codes:

| Status | Meaning |
|--------|---------|
| 200 | Success (GET, PUT, PATCH, DELETE) |
| 201 | Created (POST) |
| 400 | Bad Request (invalid JSON, missing required fields) |
| 404 | Not Found (user ID doesn't exist) |
| 405 | Method Not Allowed (unsupported HTTP method) |
| 500 | Internal Server Error (database issues) |

## Improvements / Enhancements
### For learning, for production usecase needs following

- [ ] Role-based filtering endpoints (`/role/{type}`)
- [ ] User search and filtering
- [ ] Pagination support
- [ ] Authentication and authorization
- [ ] API documentation (Swagger/OpenAPI)
- [ ] Redis caching layer
- [ ] Database migration tools
- [ ] Logging and monitoring

## Development Notes

### Go Version Management

This project uses Go 1.26. If you need to switch versions:

```bash
gvm use go1.26
```

### Database Configuration

To use a different database path, modify the `dbPath` constant in `reiley/user/user.go`:

```go
const (
    dbPath = "data/users.db"  // Change this path
)
```

### Adding New Fields to User

1. Add field to `User` struct in `user/user.go`
2. Add JSON tag for marshaling
3. Update validation logic if needed
4. Update PATCH handler allowed fields in `handlers/usersHandler.go`
5. Add tests for new field

## License

This project is part of a learning exercise for Go web development.


