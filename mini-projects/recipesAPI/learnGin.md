# Learning Gin Framework - Concepts Review

## 1. Anonymous Functions as Route Handlers

### Syntax
```go
engine.GET("/", func(c *gin.Context) {
    c.String(200, "Hello World")
})
```

**Key Points:**
- `func(c *gin.Context)` is an anonymous function (no name)
- Takes parameter `c` – pointer to gin.Context
- Gin automatically calls this when matching HTTP request arrives
- Callback pattern – you provide code for Gin to execute

---

## 2. Understanding gin.Context

**What is it?**
Struct containing:
- HTTP request details (headers, body, method)
- HTTP response writer
- Route parameters, query strings, cookies, form values, JSON data

**Why a Pointer?**
```go
func(c *gin.Context)  // Pointer
```
- Changes affect actual HTTP response
- More efficient (no copying large struct)

---

## 3. Named vs Anonymous Functions

**Named:**
```go
func handleGetUser(c *gin.Context) { ... }
engine.GET("/users/:id", handleGetUser)
```

**Anonymous:**
```go
engine.GET("/users/:id", func(c *gin.Context) { ... })
```

Both work identically – anonymous is just inline.

---

## 4. HTTP Methods in Gin

```go
engine.GET("/path", handler)
engine.POST("/path", handler)
engine.PUT("/path", handler)
engine.DELETE("/path", handler)
engine.PATCH("/path", handler)
engine.HEAD("/path", handler)
engine.Handle("METHOD", "/path", handler)  // Generic
```

---

## 5. Route Parameters

```go
engine.GET("/users/:id", func(c *gin.Context) {
    userID := c.Param("id")
    c.String(200, "User: " + userID)
})
```

Usage: `curl localhost:8087/users/123` → userID = "123"

---

## 6. Query Parameters

```go
engine.GET("/search", func(c *gin.Context) {
    name := c.Query("name")
    c.String(200, "Searching for: " + name)
})
```

Usage: `curl "localhost:8087/search?name=John"` → name = "John"

---

## 7. Sending Responses

```go
c.String(200, "Hello World")           // Text response
c.JSON(200, gin.H{"key": "value"})     // JSON response
c.HTML(200, "template.html", data)     // HTML response
c.Status(404)                          // Just status code
```

---

## 8. curl Flags

| Flag | Purpose | Example |
|------|---------|---------|
| `-X` | HTTP method | `curl -X POST localhost:8087/` |
| `-I` | Headers only (HEAD) | `curl -I localhost:8087/` |
| `-d` | Request body | `curl -d '{"name":"John"}'` |
| `-H` | Custom header | `curl -H "Content-Type: application/json"` |

**Default:** `curl localhost:8087/` = GET request

---

## 9. Gin Engine Setup

```go
engine := gin.New()           // Create engine
engine.GET("/", handler)      // Define routes
engine.Run(":8087")           // Start on port 8087
```

---

## 10. Request Flow

```
HTTP Request → Gin matches route → Creates gin.Context 
→ Calls handler function → Handler uses c.Param(), c.Query(), c.JSON() 
→ Response sent to client
```

---

## Quick Access Methods

```go
c.Param("id")                  // Route parameter
c.Query("name")                // Query string
c.GetHeader("Content-Type")    // HTTP header
c.PostForm("username")         // Form data
```

---

## Key Concepts Summary

- **Anonymous functions** = inline handlers passed to Gin
- **gin.Context** = request/response data container
- **Pointers** = efficient, changes affect original
- **Callbacks** = Gin calls your function when request matches
- **Route parameters** = `:id` in path
- **Query parameters** = `?name=value` in URL
- **HTTP methods** = GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS

## 12. gin.New() vs gin.Default()

### gin.New()
Creates a barebones Gin engine with **zero middlewares** attached.
- No request logging
- No crash recovery (a panic will crash the entire server)
- Use when you want complete control (e.g., attaching your own custom logger).

### gin.Default()
Creates a Gin engine pre-configured with two essential middlewares:
1. **Logger**: Logs incoming HTTP requests to the console.
2. **Recovery**: Catches panics, prevents server crashes, and returns a 500 status code.

Internally, `gin.Default()` is just:
```go
engine := gin.New()
engine.Use(gin.Logger(), gin.Recovery())
```

---

## 13. Swagger Integration in Gin

Swagger provides interactive API documentation. In Go, it is typically generated from comments.

### 1. Required Packages
```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

### 2. General API Annotations
Placed directly above the `main()` function:
```go
// @title My API
// @version 1.0
// @host localhost:8088
// @BasePath /
func main() { ... }
```

### 3. Route Annotations
Placed directly above handler functions.
**Important:** Path parameters use `{id}`, not `:id` in Swagger comments!

```go
// @Summary Get Recipe by ID
// @Tags recipes
// @Param id path string true "Recipe ID"
// @Success 200 {object} Recipe
// @Router /recipe/{id} [get]
func getRecipeById(c *gin.Context) { ... }
```

### 4. Serving the Docs in Gin
Add the Swagger wildcard route to your engine:
```go
import (
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "your-module-name/docs" // Triggers init() to load swagger.json
)

// In main():
engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

### 5. The Golden Rule of Swagger in Go
**Every time you change a `// @...` comment, you MUST run:**
```bash
swag init
```
If you do not run `swag init`, the generated `docs/swagger.json` will not update, and your UI will be out of sync with your code!

---

## 14. JSON Marshalling vs Unmarshalling

In Go, these terms describe converting between Go data structures and JSON strings (bytes), which is crucial for caching in Redis or sending HTTP responses.

### 1. Marshalling (Go Struct → JSON String)
Takes a Go object and converts it into a JSON byte slice.
```go
dbRecipes := []Recipe{{Name: "Pasta"}}
// data becomes: []byte(`[{"name":"Pasta"}]`)
data, err := json.Marshal(dbRecipes) 
```

### 2. Unmarshalling (JSON String → Go Struct)
Takes a JSON string (like one read from Redis) and parses it back into a Go struct.
```go
val := `[{"name":"Pasta"}]`
recipes := make([]Recipe, 0)
// MUST pass a pointer (&recipes) so the function can modify the slice!
err := json.Unmarshal([]byte(val), &recipes)
```

---

## 15. JWT Authentication Flow

JSON Web Tokens (JWT) act like secure "VIP passes" for your API. Instead of sending a username/password with every request, the user logs in once and gets a temporary token.

### The Structure of a JWT
A JWT is a single string made of three parts separated by dots (`xxxxx.yyyyy.zzzzz`):
1. **Header:** Metadata about the token (e.g., the math algorithm used, like `HS256`).
2. **Payload (Claims):** The actual data you want to store (e.g., username, expiration time). *Note: This is just base64 encoded, NOT encrypted. Anyone can read it, so don't put passwords here!*
3. **Signature:** A cryptographic stamp made by hashing the Header + Payload using your Server's Secret Key. This proves the token hasn't been altered.

### How to Create a Token (Sign-In)

1. **Validate Credentials:** Check if the provided username and password are correct (e.g., check against a database).
2. **Create Claims:** "Claims" are the data written on the token (like the username and an expiration time).
   ```go
   claims := &Claims{Username: "admin", StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(30 * time.Minute).Unix()}}
   ```
3. **Initialize the Token:** Create the token structure in memory using a specific math algorithm (like HS256).
   ```go
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   ```
4. **Stamp with Secret Key:** The server signs the token using a secret string. 
   - **What is it?** It is a local string that lives *only* on your server. You can technically use any string you want (e.g., `"my_secret_password"`).
   - **Why is it important?** If someone guesses or steals this string, they can forge their own valid tokens. In production, this should be a long, random, cryptographically secure string loaded from an `.env` file and *never* committed to GitHub.
   ```go
   tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
   ```
   This `tokenString` is what you send back to the user to include in future requests!

---

## 16. HTTPS, SSL Certificates, and Ngrok (Reverse Proxies)

When developing a REST API, you need to secure the traffic using HTTPS (TLS/SSL). 
### Create local certificates
`openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout certs/localhost.key -out certs/localhost.crt`

### Why Local Certificates Say "Not Trusted"
If you generate a certificate locally using `openssl` (a self-signed certificate) and run your Gin server with `engine.RunTLS(...)`, your browser will show a "Not Trusted" warning. This happens because browsers only trust certificates signed by pre-installed, global "Certificate Authorities" (CAs) like Let's Encrypt or DigiCert. Your self-signed cert is technically encrypted, but the browser cannot verify *who* created it.

### How Ngrok Solves This for Local Development
When you run `ngrok http 8088`, Ngrok acts as a reverse proxy:
1. **The Public Leg (Encrypted):** Ngrok generates a public URL (e.g., `https://1234.ngrok-free.app`) with a valid, globally trusted Let's Encrypt certificate. External users hit this URL, and their browser shows a green padlock.
2. **The Tunnel Leg (Encrypted):** Ngrok tunnels the HTTPS traffic securely to your local machine.
3. **The Local Leg (Unencrypted):** The Ngrok agent on your Mac decrypts the traffic and forwards it as standard HTTP to your Go server running on `localhost:8088`.

Because of this, your Go application does not need to handle certificates directly. You just run `engine.Run(":8088")`. External users interact with your API by replacing `http://localhost:8088` with the Ngrok URL in their `curl` commands, Postman, or UI applications.

### Production Deployment Options
In a real production environment (like AWS or DigitalOcean), you have two main options:
1. **Reverse Proxy (Industry Standard):** You place a proxy server (Nginx, Caddy, AWS ALB) in front of your Go app. The proxy handles Let's Encrypt certificates and TLS decryption, then forwards plain HTTP to your internal Go server.
2. **Go Handles TLS Directly:** You configure your Go app to fetch and manage Let's Encrypt certificates itself (e.g., using `golang.org/x/crypto/acme/autocert`), binding directly to port 443.
