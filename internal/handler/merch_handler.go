package handler

import (
	"avito_internship_task/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type MerchHandler struct {
	db models.DB
}

func CreateMerchHandler(db models.DB) *MerchHandler {
	return &MerchHandler{
		db: db,
	}
}

func (mh *MerchHandler) GetMerchList(w http.ResponseWriter, r *http.Request) {
	var items []models.Merch
	if err := mh.db.Find(&items).Error; err != nil {
		http.Error(w, "Ошибка при получении списка мерча", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(models.MerchListResponse{Items: items}); err != nil {
		http.Error(w, "Ошибка при кодировании ответа", http.StatusInternalServerError)
		return
	}
}

func (mh *MerchHandler) BuyMerch(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int64)
	vars := mux.Vars(r)
	merchName := vars["item"]

	tx := mh.db.Begin()

	var merch models.Merch
	if err := tx.Where("merch_name = ?", merchName).First(&merch).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Товар не найден", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := tx.Where("user_id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
		return
	}

	if user.Balance < int64(merch.Price) {
		tx.Rollback()
		http.Error(w, "Не достаточно средств", http.StatusBadRequest)
		return
	}

	var userMerch models.UserMerch
	err := tx.Where("user_id = ? AND merch_id = ?", userID, merch.MerchID).First(&userMerch).Error

	if err == nil {
		if err := tx.Model(&userMerch).Update("quantity", userMerch.Quantity+1).Error; err != nil {
			tx.Rollback()
			http.Error(w, "Ошибка при обновлении количества", http.StatusInternalServerError)
			return
		}
	} else {
		userMerch = models.UserMerch{
			UserID:   userID,
			MerchID:  merch.MerchID,
			Quantity: 1,
		}
		if err := tx.Create(&userMerch).Error; err != nil {
			tx.Rollback()
			http.Error(w, "Ошибка при создании записи", http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Model(&user).Update("balance", user.Balance-int64(merch.Price)).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Ошибка при списании денег", http.StatusInternalServerError)
		return
	}

	tx.Commit()

	type buyResponse struct {
		Message  string `json:"message"`
		Balance  int64  `json:"balance"`
		Quantity int64  `json:"quantity"`
	}

	response := buyResponse{
		Message:  "Покупка успешно совершена",
		Balance:  user.Balance,
		Quantity: userMerch.Quantity,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Ошибка при кодировании ответа", http.StatusInternalServerError)
		return
	}
}
