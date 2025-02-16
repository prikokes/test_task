package integrated_tests

import (
	"avito_internship_task/internal/handler"
	"avito_internship_task/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransferCoins(t *testing.T) {
	db, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	sender := models.User{
		Username:     "sender",
		HashPassword: "hash123",
		Balance:      1000,
	}
	db.Create(&sender)

	receiver := models.User{
		Username:     "receiver",
		HashPassword: "hash123",
		Balance:      0,
	}
	db.Create(&receiver)

	handler := handler.CreateTransactionsHandler(db)

	transferReq := models.TransferRequest{
		ToUsername: "receiver",
		Money:      500,
	}
	body, _ := json.Marshal(transferReq)
	req := httptest.NewRequest("POST", "/transfer", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	ctx := context.WithValue(req.Context(), "user_id", sender.UserID)
	req = req.WithContext(ctx)

	handler.TransferCoins(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedSender, updatedReceiver models.User
	db.First(&updatedSender, sender.UserID)
	db.First(&updatedReceiver, receiver.UserID)

	assert.Equal(t, int64(500), updatedSender.Balance)
	assert.Equal(t, int64(500), updatedReceiver.Balance)
}

func TestTransferCoinsInsufficientFunds(t *testing.T) {
	db, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	sender := models.User{
		Username:     "poorSender",
		HashPassword: "hash123",
		Balance:      100,
	}
	db.Create(&sender)

	receiver := models.User{
		Username:     "receiver",
		HashPassword: "hash123",
		Balance:      0,
	}
	db.Create(&receiver)

	handler := handler.CreateTransactionsHandler(db)

	transferReq := models.TransferRequest{
		ToUsername: "receiver",
		Money:      500,
	}
	body, _ := json.Marshal(transferReq)
	req := httptest.NewRequest("POST", "/transfer", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	ctx := context.WithValue(req.Context(), "user_id", sender.UserID)
	req = req.WithContext(ctx)

	handler.TransferCoins(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Не достаточно средств")
}
