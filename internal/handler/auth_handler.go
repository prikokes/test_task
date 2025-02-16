package handler

import (
	"avito_internship_task/internal/models"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type AuthHandler struct {
	db models.DB
}

func CreateAuthHandler(db models.DB) *AuthHandler {
	return &AuthHandler{
		db: db,
	}
}

func (authHandler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неправильный запрос", http.StatusBadRequest)
		return
	}

	var user models.User
	var count int64

	if err := authHandler.db.Model(&models.User{}).Count(&count).Error; err != nil {
		http.Error(w, "Ошибка подсчета пользователей", http.StatusInternalServerError)
		return
	}

	if err := authHandler.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if count >= 100000 {
			http.Error(w, "Достигнут лимит пользователей", http.StatusForbidden)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.HashPassword), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Ошибка при хэшировании пароля", http.StatusInternalServerError)
			return
		}

		user = models.User{
			Username:     req.Username,
			HashPassword: string(hashedPassword),
			Balance:      1000,
		}

		if err := authHandler.db.Create(&user).Error; err != nil {
			http.Error(w, "Ошибка при создании пользователя", http.StatusInternalServerError)
			return
		}
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(req.HashPassword)); err != nil {
			http.Error(w, "Неверный пароль", http.StatusUnauthorized)
			return
		}
	}

	type extendedLoginResponse struct {
		Token   string `json:"token"`
		Balance int64  `json:"balance"`
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.UserID
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		http.Error(w, "Ошибка при создании токена", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(extendedLoginResponse{
		Token:   tokenString,
		Balance: user.Balance,
	}); err != nil {
		http.Error(w, "Ошибка при кодировании ответа", http.StatusInternalServerError)
		return
	}
}
