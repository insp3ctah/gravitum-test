package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"gravitum-test/internal/user"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MockService struct{}

func (m *MockService) CreateUser(name, email string) (*user.User, error) {
	return &user.User{Model: gorm.Model{ID: 1}, Name: name, Email: email}, nil
}

func (m *MockService) GetUser(id uint) (*user.User, error) {
	if id == 1 {
		return &user.User{Model: gorm.Model{ID: 1}, Name: "Test User", Email: "test@example.com"}, nil
	}
	return nil, nil
}

func (m *MockService) UpdateUser(id uint, name, email string) (*user.User, error) {
	log.Printf("id is %d", id)
	if id != 1 {
		return nil, errors.New("пользователь не найден") // ✅ Добавляем ошибку
	}
	return &user.User{Model: gorm.Model{ID: 1}, Name: name, Email: email}, nil
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	var mockService user.UserService = &MockService{} // ✅ Интерфейсное приведение
	handler := user.NewHandler(mockService)

	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUser)
	r.PUT("/users/:id", handler.UpdateUser)

	return r
}

func TestCreateUser(t *testing.T) {
	router := setupRouter()

	payload := `{"name": "John Doe", "email": "john@example.com"}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestGetUser(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/users/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var user user.User
	err := json.Unmarshal(resp.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, "Test User", user.Name)
}

func TestUpdateUser(t *testing.T) { // TODO fix this test
	router := setupRouter()

	payload := `{"name": "Updated Name", "email": "updated@example.com"}`
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var updatedUser user.User
	err := json.Unmarshal(resp.Body.Bytes(), &updatedUser)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", updatedUser.Name)
	assert.Equal(t, "updated@example.com", updatedUser.Email)
}
