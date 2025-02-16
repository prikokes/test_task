package integrated_tests

import (
	"avito_internship_task/internal/handler"
	"avito_internship_task/internal/models"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetMerchList(t *testing.T) {
	db, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	testMerch := []models.Merch{
		{MerchName: "Футболка", Price: 100},
		{MerchName: "Кепка", Price: 50},
	}
	for _, m := range testMerch {
		db.Create(&m)
	}

	handler := handler.CreateMerchHandler(db)
	req := httptest.NewRequest("GET", "/merch", nil)
	w := httptest.NewRecorder()

	handler.GetMerchList(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.MerchListResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Len(t, response.Items, 2)
}

func TestBuyMerch(t *testing.T) {
	db, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	// Создаем тестового пользователя
	user := models.User{
		Username:     "testuser",
		HashPassword: "hash123",
		Balance:      1000,
	}
	db.Create(&user)

	// Создаем тестовый мерч
	merch := models.Merch{
		MerchName: "Футболка",
		Price:     100,
	}
	db.Create(&merch)

	handler := handler.CreateMerchHandler(db)
	req := httptest.NewRequest("POST", "/merch/buy/Футболка", nil)
	w := httptest.NewRecorder()

	ctx := context.WithValue(req.Context(), "user_id", user.UserID)
	req = req.WithContext(ctx)

	vars := map[string]string{
		"item": "Футболка",
	}
	req = mux.SetURLVars(req, vars)

	handler.BuyMerch(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedUser models.User
	db.First(&updatedUser, user.UserID)
	assert.Equal(t, int64(900), updatedUser.Balance)

	var userMerch models.UserMerch
	db.Where("user_id = ? AND merch_id = ?", user.UserID, merch.MerchID).First(&userMerch)
	assert.Equal(t, int64(1), userMerch.Quantity)
}

func TestBuyMerchInsufficientFunds(t *testing.T) {
	db, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	user := models.User{
		Username:     "pooruser",
		HashPassword: "hash123",
		Balance:      10,
	}
	db.Create(&user)

	merch := models.Merch{
		MerchName: "Футболка",
		Price:     100,
	}
	db.Create(&merch)

	handler := handler.CreateMerchHandler(db)
	req := httptest.NewRequest("POST", "/merch/buy/Футболка", nil)
	w := httptest.NewRecorder()

	ctx := context.WithValue(req.Context(), "user_id", user.UserID)
	req = req.WithContext(ctx)

	vars := map[string]string{
		"item": "Футболка",
	}
	req = mux.SetURLVars(req, vars)

	handler.BuyMerch(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Не достаточно средств")
}
