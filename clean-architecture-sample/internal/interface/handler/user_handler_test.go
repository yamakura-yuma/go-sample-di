package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"clean-architecture-sample/internal/infrastructure"
	"clean-architecture-sample/internal/usecase"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	repo := infrastructure.NewFileUserRepository("../../../data/test_users_handler.json")
	uc := usecase.NewUserUseCase(repo)
	h := NewUserHandler(uc)
	r := gin.Default()
	r.POST("/users", h.CreateUserHandler)
	r.GET("/users/:id", h.GetUserHandler)
	return r
}

func TestCreateUserHandler(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{"id": "1", "name": "GinTest", "email": "gin@example.com"})
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != 201 {
		t.Errorf("expected 201, got %d", w.Code)
	}
}
