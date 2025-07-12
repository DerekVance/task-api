package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-api/models"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", GetTasks)
	router.POST("/tasks", CreateTask)
	router.PUT("/tasks/:id", UpdateTask)
	router.DELETE("/tasks/:id", DeleteTask)
	return router
}

func resetTaskStore() {
	taskMap = make(map[int]models.Task)
	nextID = 1
}

func TestCreateTask(t *testing.T) {
	resetTaskStore()
	router := setupRouter()

	body := models.Task{
		Title:     "Test Task",
		Completed: false,
	}
	jsonValue, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.Code)
	}

	var result models.Task
	json.Unmarshal(resp.Body.Bytes(), &result)

	if result.ID != 1 || result.Title != body.Title {
		t.Errorf("Unexpected task: %+v", result)
	}
}

func TestGetTasks(t *testing.T) {
	resetTaskStore()
	taskMap[1] = models.Task{ID: 1, Title: "Read book", Completed: false}

	router := setupRouter()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}

	var tasks []models.Task
	json.Unmarshal(resp.Body.Bytes(), &tasks)

	if len(tasks) != 1 || tasks[0].Title != "Read book" {
		t.Errorf("Unexpected response: %+v", tasks)
	}
}

func TestUpdateTask(t *testing.T) {
	resetTaskStore()
	taskMap[1] = models.Task{ID: 1, Title: "Old Title", Completed: false}

	router := setupRouter()

	updated := models.Task{Title: "Updated Title", Completed: true}
	body, _ := json.Marshal(updated)

	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}

	var result models.Task
	json.Unmarshal(resp.Body.Bytes(), &result)

	if result.Title != "Updated Title" || !result.Completed {
		t.Errorf("Task not updated correctly: %+v", result)
	}
}

func TestDeleteTask(t *testing.T) {
	resetTaskStore()
	taskMap[1] = models.Task{ID: 1, Title: "To Delete", Completed: false}

	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", resp.Code)
	}

	if _, exists := taskMap[1]; exists {
		t.Error("Task was not deleted")
	}
}

func TestUpdateNonexistentTask(t *testing.T) {
	resetTaskStore()
	router := setupRouter()

	update := models.Task{Title: "Nothing", Completed: true}
	body, _ := json.Marshal(update)

	req, _ := http.NewRequest("PUT", "/tasks/999", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", resp.Code)
	}
}

func TestDeleteNonexistentTask(t *testing.T) {
	resetTaskStore()
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/tasks/123", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", resp.Code)
	}
}
