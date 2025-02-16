package integrated_tests

import (
	"avito_internship_task/internal/handler"
	"avito_internship_task/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestMain(m *testing.M) {
	os.Setenv("JWT_SECRET", "test_secret_key")
	code := m.Run()
	os.Exit(code)
}

func TestLoginNewUser(t *testing.T) {
	db, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	handler := handler.CreateAuthHandler(db)

	loginReq := models.LoginRequest{
		Username:     "newuser",
		HashPassword: "password123",
	}

	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/api/auth", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.Login(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Token   string `json:"token"`
		Balance int64  `json:"balance"`
	}
	err = json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Token)
	assert.Equal(t, int64(1000), response.Balance)

	var count int64
	db.Model(&models.User{}).Where("username = ?", "newuser").Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestLoginExistingUser(t *testing.T) {
	db, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	existingUser := models.User{
		Username:     "existinguser",
		HashPassword: string(hashedPass),
		Balance:      1000,
	}
	db.Create(&existingUser)

	handler := handler.CreateAuthHandler(db)

	loginReq := models.LoginRequest{
		Username:     "existinguser",
		HashPassword: "password123",
	}
	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/auth", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.Login(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Token   string `json:"token"`
		Balance int64  `json:"balance"`
	}
	err = json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Token)
	assert.Equal(t, int64(1000), response.Balance)
}

func TestLoginWrongPassword(t *testing.T) {
	db, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	existingUser := models.User{
		Username:     "existinguser",
		HashPassword: string(hashedPass),
		Balance:      1000,
	}
	db.Create(&existingUser)

	handler := handler.CreateAuthHandler(db)

	loginReq := models.LoginRequest{
		Username:     "existinguser",
		HashPassword: "wrongpassword",
	}
	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/auth", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.Login(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Неверный пароль")
}

func TestLoginUserLimit(t *testing.T) {
	db, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	for i := 0; i < 100000; i++ {
		user := models.User{
			Username:     fmt.Sprintf("user%d", i),
			HashPassword: "hash123",
			Balance:      1000,
		}
		db.Create(&user)
	}

	handler := handler.CreateAuthHandler(db)

	loginReq := models.LoginRequest{
		Username:     "onemore",
		HashPassword: "password123",
	}
	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/auth", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.Login(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "Достигнут лимит пользователей")
}

func TestLoginInvalidJSON(t *testing.T) {
	db, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	handler := handler.CreateAuthHandler(db)

	req := httptest.NewRequest("POST", "/auth", bytes.NewBuffer([]byte("invalid json")))
	w := httptest.NewRecorder()

	handler.Login(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Неправильный запрос")
}
