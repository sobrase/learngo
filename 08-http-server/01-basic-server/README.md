# Exercise 8: HTTP Server

## Objective
Learn about Go's `net/http` package and building web servers.

## Problem
Create a simple HTTP server with different endpoints and request handling.

## Requirements
1. Create a basic HTTP server using `http.ListenAndServe()`
2. Define HTTP handlers for different endpoints:
   - GET requests
   - POST requests
   - JSON responses
3. Use `http.HandleFunc()` and custom handlers
4. Work with request parameters and body
5. Use middleware for logging and authentication
6. Handle different HTTP status codes

## Tasks
1. Create a simple web server with multiple endpoints
2. Implement a REST API for a simple todo list
3. Add middleware for request logging
4. Handle JSON requests and responses
5. Implement basic error handling

## Expected Output
```
Starting HTTP server on :8080...

GET / - Welcome message
GET /health - Health check
GET /api/todos - List all todos
POST /api/todos - Add new todo
GET /api/todos/{id} - Get specific todo
DELETE /api/todos/{id} - Delete todo

Server logs:
2024/01/01 12:00:00 GET / - 200 OK
2024/01/01 12:00:01 POST /api/todos - 201 Created
2024/01/01 12:00:02 GET /api/todos - 200 OK
2024/01/01 12:00:03 GET /api/todos/1 - 200 OK
2024/01/01 12:00:04 DELETE /api/todos/1 - 200 OK
2024/01/01 12:00:05 GET /api/todos/999 - 404 Not Found

API Responses:
GET /api/todos:
{
  "todos": [
    {"id": 1, "title": "Learn Go", "completed": false},
    {"id": 2, "title": "Build API", "completed": true}
  ]
}

POST /api/todos:
{
  "id": 3,
  "title": "New Todo",
  "completed": false
}
```

## Hints
- Use `http.HandleFunc()` to register route handlers
- Use `http.ListenAndServe()` to start the server
- Use `json.Marshal()` and `json.Unmarshal()` for JSON handling
- Use `http.Error()` to send error responses
- Use `http.StatusText()` to get status code descriptions
- Remember to set `Content-Type` header for JSON responses
- Use `gorilla/mux` for more advanced routing (optional)
- Handle panics with `recover()` in handlers
