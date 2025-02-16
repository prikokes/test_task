package models

type TransferRequest struct {
	ToUsername string `json:"to_username"`
	Money      int64  `json:"money"`
}

type LoginRequest struct {
	Username     string `json:"username"`
	HashPassword string `json:"hashPassword"`
}
