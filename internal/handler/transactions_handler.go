package handler

import (
	"avito_internship_task/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

type TransactionsHandler struct {
	db models.DB
}

func CreateTransactionsHandler(db models.DB) *TransactionsHandler {
	return &TransactionsHandler{
		db: db,
	}
}

func (th *TransactionsHandler) TransferCoins(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int64)

	var req models.TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неправильный запрос.", http.StatusBadRequest)
		return
	}

	tx := th.db.Begin()

	var fromUser models.User
	if err := tx.Where("user_id = ?", userID).First(&fromUser).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Пользователь не авторизован.", http.StatusUnauthorized)
		return
	}

	var toUser models.User
	if err := tx.Where("username = ?", req.ToUsername).First(&toUser).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Получатель не найден.", http.StatusBadRequest)
		return
	}

	if fromUser.Balance < req.Money {
		tx.Rollback()
		http.Error(w, "Не достаточно средств.", http.StatusBadRequest)
		return
	}

	if err := tx.Model(&fromUser).Update("balance", fromUser.Balance-req.Money).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Ошибка при списании.", http.StatusInternalServerError)
		return
	}

	if err := tx.Model(&toUser).Update("balance", toUser.Balance+req.Money).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Ошибка при зачислении.", http.StatusInternalServerError)
		return
	}

	transaction := models.Transaction{
		FromUserID: userID,
		ToUserID:   toUser.UserID,
		Money:      req.Money,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Ошибка при сохранении транзакции", http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
}

func (th *TransactionsHandler) GetTransactionsInfo(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int64)

	var received []struct {
		Username string
		Money    int64
	}

	if err := th.db.Table("transactions").
		Select("users.username, transactions.money").
		Joins("JOIN users ON users.user_id = transactions.from_user_id").
		Where("transactions.to_user_id = ?", userID).
		Scan(&received).Error; err != nil {
		http.Error(w, "Ошибка при получении входящих транзакций", http.StatusInternalServerError)
		return
	}

	var sent []struct {
		Username string
		Money    int64
	}

	if err := th.db.Table("transactions").
		Select("users.username, transactions.money").
		Joins("JOIN users ON users.user_id = transactions.to_user_id").
		Where("transactions.from_user_id = ?", userID).
		Scan(&sent).Error; err != nil {
		http.Error(w, "Ошибка при получении исходящих транзакций", http.StatusInternalServerError)
		return
	}

	response := models.TransactionsResponse{
		Received: make([]models.TransactionInfo, len(received)),
		Sent:     make([]models.TransactionInfo, len(sent)),
	}

	for i, r := range received {
		response.Received[i] = models.TransactionInfo{
			Username: r.Username,
			Money:    r.Money,
		}
	}

	for i, s := range sent {
		response.Sent[i] = models.TransactionInfo{
			Username: s.Username,
			Money:    s.Money,
		}
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Ошибка при кодировании ответа: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
