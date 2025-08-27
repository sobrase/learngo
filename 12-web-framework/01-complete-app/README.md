# Exercise 13: Complete Web Application

## Objective
Build a complete web application that combines all the concepts learned throughout the curriculum.

## Problem
Create a RESTful web application for a simple task management system with database integration, authentication, and real-time updates.

## Requirements
1. **HTTP Server**: Use `net/http` for routing and request handling
2. **Database**: Use `database/sql` with SQLite for data persistence
3. **JSON**: Use `encoding/json` for API communication
4. **Concurrency**: Use goroutines and channels for background tasks
5. **Testing**: Write comprehensive unit tests
6. **Error Handling**: Implement proper error handling throughout
7. **Configuration**: Use environment variables and config files
8. **Logging**: Implement structured logging
9. **Middleware**: Create authentication and logging middleware
10. **File I/O**: Handle file uploads and static file serving

## Features to Implement
1. **User Management**:
   - User registration and authentication
   - JWT token-based authentication
   - Password hashing with bcrypt

2. **Task Management**:
   - CRUD operations for tasks
   - Task status tracking (pending, in-progress, completed)
   - Task categories and priorities
   - Due date management

3. **Real-time Features**:
   - WebSocket connections for real-time updates
   - Task assignment notifications
   - Live task status updates

4. **API Endpoints**:
   ```
   POST   /api/auth/register    - User registration
   POST   /api/auth/login       - User login
   GET    /api/tasks            - List all tasks
   POST   /api/tasks            - Create new task
   GET    /api/tasks/{id}       - Get specific task
   PUT    /api/tasks/{id}       - Update task
   DELETE /api/tasks/{id}       - Delete task
   GET    /api/users/{id}/tasks - Get user's tasks
   WS     /ws                   - WebSocket connection
   ```

5. **Database Schema**:
   ```sql
   CREATE TABLE users (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
       username TEXT UNIQUE NOT NULL,
       email TEXT UNIQUE NOT NULL,
       password_hash TEXT NOT NULL,
       created_at DATETIME DEFAULT CURRENT_TIMESTAMP
   );

   CREATE TABLE tasks (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
       title TEXT NOT NULL,
       description TEXT,
       status TEXT DEFAULT 'pending',
       priority TEXT DEFAULT 'medium',
       category TEXT,
       due_date DATETIME,
       user_id INTEGER,
       assigned_to INTEGER,
       created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
       updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
       FOREIGN KEY (user_id) REFERENCES users(id),
       FOREIGN KEY (assigned_to) REFERENCES users(id)
   );
   ```

## Expected Output
```
Starting Task Management System...
Database initialized successfully
Server starting on :8080

API Documentation:
- POST /api/auth/register - Register new user
- POST /api/auth/login - User login
- GET /api/tasks - List tasks (requires auth)
- POST /api/tasks - Create task (requires auth)
- GET /api/tasks/{id} - Get task details
- PUT /api/tasks/{id} - Update task
- DELETE /api/tasks/{id} - Delete task
- GET /api/users/{id}/tasks - Get user tasks
- WS /ws - WebSocket for real-time updates

Example API Usage:

1. Register a user:
POST /api/auth/register
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "secure_password"
}

2. Login:
POST /api/auth/login
{
  "username": "john_doe",
  "password": "secure_password"
}

3. Create a task:
POST /api/tasks
Authorization: Bearer <jwt_token>
{
  "title": "Learn Go Programming",
  "description": "Complete the Go learning curriculum",
  "priority": "high",
  "category": "learning",
  "due_date": "2024-02-01T00:00:00Z"
}

4. Get all tasks:
GET /api/tasks
Authorization: Bearer <jwt_token>

Response:
{
  "tasks": [
    {
      "id": 1,
      "title": "Learn Go Programming",
      "description": "Complete the Go learning curriculum",
      "status": "pending",
      "priority": "high",
      "category": "learning",
      "due_date": "2024-02-01T00:00:00Z",
      "user_id": 1,
      "created_at": "2024-01-01T12:00:00Z"
    }
  ],
  "total": 1
}
```

## Project Structure
```
12-web-framework/01-complete-app/
├── main.go              # Application entry point
├── config/
│   └── config.go        # Configuration management
├── models/
│   ├── user.go          # User model and methods
│   └── task.go          # Task model and methods
├── handlers/
│   ├── auth.go          # Authentication handlers
│   ├── task.go          # Task handlers
│   └── websocket.go     # WebSocket handlers
├── middleware/
│   ├── auth.go          # Authentication middleware
│   └── logging.go       # Logging middleware
├── database/
│   └── db.go            # Database connection and queries
├── utils/
│   ├── jwt.go           # JWT utilities
│   └── password.go      # Password hashing utilities
├── tests/
│   ├── auth_test.go     # Authentication tests
│   └── task_test.go     # Task management tests
└── static/
    └── index.html       # Simple web interface
```

## Hints
- Use `gorilla/mux` for advanced routing
- Use `golang-jwt/jwt` for JWT token handling
- Use `golang.org/x/crypto/bcrypt` for password hashing
- Use `github.com/gorilla/websocket` for WebSocket support
- Use `github.com/mattn/go-sqlite3` for SQLite driver
- Implement proper error handling with custom error types
- Use context for request cancellation and timeouts
- Implement rate limiting for API endpoints
- Use environment variables for configuration
- Write comprehensive tests with table-driven tests
