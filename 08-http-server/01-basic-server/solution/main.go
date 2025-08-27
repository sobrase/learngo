package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Todo struct for the API
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// In-memory storage for todos
var todos = []Todo{
	{ID: 1, Title: "Learn Go", Completed: false, CreatedAt: time.Now()},
	{ID: 2, Title: "Build API", Completed: true, CreatedAt: time.Now()},
}
var nextID = 3

// Middleware for logging
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s %s %s - %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	}
}

// Home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
		<h1>Welcome to Go HTTP Server</h1>
		<p>Available endpoints:</p>
		<ul>
			<li>GET / - This page</li>
			<li>GET /health - Health check</li>
			<li>GET /api/todos - List all todos</li>
			<li>POST /api/todos - Create new todo</li>
			<li>GET /api/todos/{id} - Get specific todo</li>
			<li>PUT /api/todos/{id} - Update todo</li>
			<li>DELETE /api/todos/{id} - Delete todo</li>
		</ul>
	`))
}

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	})
}

// Get all todos
func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"todos": todos,
		"total": len(todos),
	}
	json.NewEncoder(w).Encode(response)
}

// Create new todo
func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newTodo.ID = nextID
	newTodo.CreatedAt = time.Now()
	nextID++

	todos = append(todos, newTodo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// Get specific todo
func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/todos/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

// Update todo
func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/todos/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			updatedTodo.ID = id
			updatedTodo.CreatedAt = todo.CreatedAt
			todos[i] = updatedTodo

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

// Delete todo
func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/todos/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

// Todo router
func todoRouter(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/api/todos":
		getTodosHandler(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/api/todos":
		createTodoHandler(w, r)
	case r.Method == http.MethodGet && len(r.URL.Path) > len("/api/todos/"):
		getTodoHandler(w, r)
	case r.Method == http.MethodPut && len(r.URL.Path) > len("/api/todos/"):
		updateTodoHandler(w, r)
	case r.Method == http.MethodDelete && len(r.URL.Path) > len("/api/todos/"):
		deleteTodoHandler(w, r)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func main() {
	fmt.Println("Starting HTTP server on :8080...")

	// Set up routes with middleware
	http.HandleFunc("/", loggingMiddleware(homeHandler))
	http.HandleFunc("/health", loggingMiddleware(healthHandler))
	http.HandleFunc("/api/todos", loggingMiddleware(todoRouter))
	http.HandleFunc("/api/todos/", loggingMiddleware(todoRouter))

	// Start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
