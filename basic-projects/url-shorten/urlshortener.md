# URL Shortener - HTTP Package Learning Notes

## Project Overview
### Add a URL
```
curl -X POST http://localhost:30805/shorten \
-H "Content-Type: application/json" \
-d '{"url": "https://github.com/ramk1t10/golearn10"}'
```
```
{"url":"https://github.com/ramk1t10/golearn10","short_url":"http://localhost:30805/sh/Ff0Bcv"}
```
### Redirect to a URL
curl -v http://localhost:30805/sh/aBc123

## Core HTTP Concepts

### Server Setup
```go
http.ListenAndServe(":30805", nil)  // Start server on port
http.HandleFunc("/path", handler)    // Register route
```

### Handler Function
```go
func handler(w http.ResponseWriter, r *http.Request)
```
- `r` = **Request** (incoming data)
- `w` = **ResponseWriter** (outgoing data)

---

## http.Request (r) - Reading Input

### Method
```go
r.Method  // "GET", "POST", "PUT", etc.
```

### URL Path
```go
r.URL.Path                    // "/sh/aBc123"
r.URL.Path[len("/sh/"):]      // Extract "aBc123"
```

### Body
```go
json.NewDecoder(r.Body).Decode(&req)  // Parse JSON from request
```

### Headers
```go
r.Header["Content-Type"]  // Access headers
```

---

## http.ResponseWriter (w) - Writing Output

### Set Headers
```go
w.Header().Set("Content-Type", "application/json")
```

### Write Status Code
```go
w.WriteHeader(http.StatusOK)  // 200
```

### Write Body
```go
json.NewEncoder(w).Encode(data)  // Send JSON response
```

### Helper Functions
```go
http.Error(w, "message", http.StatusBadRequest)        // Send error
http.Redirect(w, r, url, http.StatusFound)             // Send redirect (302)
```

---

## Request Flow

```
Client Request
    ↓
http.ListenAndServe receives
    ↓
Routes to handler via http.HandleFunc
    ↓
Handler function called with (w, r)
    ↓
Read from r (method, path, body)
    ↓
Process logic
    ↓
Write to w (headers, status, body)
    ↓
Response sent to client
```

---

## Status Codes Used

- `200` - OK
- `302` - Found (redirect)
- `400` - Bad Request
- `404` - Not Found
- `405` - Method Not Allowed

---

## Key Patterns

### Validate Method
```go
if r.Method != http.MethodPost {
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    return
}
```

### Parse JSON Request
```go
var req URLData
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
}
```

### Send JSON Response
```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
```

### Redirect
```go
http.Redirect(w, r, longURL, http.StatusFound)
```

---

## Project Structure

```
main.go         → Server startup
shortener.go    → Handlers + logic
urlstore.json   → Persistence
```

---

## Testing

```bash
# Shorten URL
curl -X POST http://localhost:30805/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com"}'

# Test redirect
curl -v http://localhost:30805/sh/aBc123
```
Test HTTP handlers without starting a real server.

### Key Functions

#### 1. Create Mock Request
```go
import "net/http/httptest"

req := httptest.NewRequest("POST", "/path", body)
req.Header.Set("Content-Type", "application/json")
```

**Parameters:**
- Method: `"GET"`, `"POST"`, etc.
- URL: `"/path"`
- Body: `strings.NewReader(jsonString)` or `nil`

#### 2. Record Response
```go
w := httptest.NewRecorder()
handler(w, req)
```

**Access response:**
- `w.Code` - Status code (200, 404, etc.)
- `w.Body.String()` - Response body
- `w.Header()` - Response headers

---
---

## Key Learnings

- **Handler signature**: `func(w http.ResponseWriter, r *http.Request)`
- **Read from r**: Method, URL, Body, Headers
- **Write to w**: Headers, Status, Body
- **Order matters**: Set headers before writing body
- **JSON encoding**: Use `json.Encoder/Decoder` with streams
- **Error handling**: Use `http.Error()` helper
- **Redirects**: Use `http.Redirect()` with 302 status
