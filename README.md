# 📝 Task API (Go + Gin)

A simple RESTful API for managing tasks, built with Go and the Gin web framework. This project demonstrates a clean Go project structure, modular design, and basic CRUD functionality using in-memory storage.

---

## 🚀 Features

- 📄 List all tasks
- ➕ Create new tasks
- ✏️ Update existing tasks
- ❌ Delete tasks
- 🧱 Modular project layout
- 💾 In-memory storage (no database required)

---

## 📂 Project Structure

```yaml
task-api/
├── cmd/
│ └── server/ # Application entry point
│ └── main.go
├── controllers/ # HTTP handlers
│ └── task_controller.go
├── models/ # Data models
│ └── task.go
├── routes/ # Route definitions
│ └── task_routes.go
├── go.mod
└── go.sum
```

---

## 📦 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/DerekVance/task-api.git
cd task-api
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Server

```bash
go run cmd/server/main.go
```

The server will be running at:
➡️ http://localhost:8080

📮 API Endpoints
GET /tasks/
Returns a list of all tasks.

POST /tasks/
Creates a new task.
Request Body:

```json
{
  "title": "Setup Models",
  "completed": false
}
```

PUT /tasks/:id
Updates a task by ID.
Request Body:

```json
{
  "title": "Setup Database",
  "completed": true
}
```

DELETE /tasks/:id
Deletes a task by ID.

📌 Example curl Requests

# Get all tasks

curl http://localhost:8080/tasks

# Create a task

curl -X POST http://localhost:8080/tasks/\
 -H "Content-Type: application/json" \
 -d '{"title":"Learn Go", "completed":false}'

# Update a task

curl -X PUT http://localhost:8080/tasks/1 \
 -H "Content-Type: application/json" \
 -d '{"title":"Updated Task", "completed":true}'

# Delete a task

curl -X DELETE http://localhost:8080/tasks/1

🔮 Future Improvements

🔐 Add JWT authentication

🧪 Write unit and integration tests

🗃️ Add persistent database
