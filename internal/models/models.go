package models

import (
	"time"
)

type User struct {
	UserID       int64  `gorm:"column:user_id;primary_key;autoIncrement"`
	Username     string `gorm:"column:username;size:50;not null"`
	HashPassword string `gorm:"column:hash_password;size:50;not null"`
	Balance      int64  `gorm:"column:balance;not null"`
}

type Session struct {
	SessionID    int64     `gorm:"column:session_id;primary_key;autoIncrement"`
	UserID       int64     `gorm:"column:user_id;not null"`
	SessionToken string    `gorm:"column:session_token;size:255;not null"`
	ExpiresAt    time.Time `gorm:"column:expires_at;not null"`
}

type Transaction struct {
	TransactionID int64     `gorm:"column:transaction_id;primary_key;autoIncrement"`
	FromUserID    int64     `gorm:"column:from_user_id;not null"`
	ToUserID      int64     `gorm:"column:to_user_id;not null"`
	Money         int64     `gorm:"column:money;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;not null"`
}

type Merch struct {
	MerchID   int64  `gorm:"column:merch_id;primary_key;autoIncrement"`
	MerchName string `gorm:"column:merch_name;size:20;not null"`
	Price     int64  `gorm:"column:price;not null"`
}

type UserMerch struct {
	UserID   int64 `gorm:"column:user_id;primaryKey"`
	MerchID  int64 `gorm:"column:merch_id;primaryKey"`
	Quantity int64 `gorm:"column:quantity;not null"`
}

func (User) TableName() string {
	return "users"
}

func (Session) TableName() string {
	return "session"
}

func (Transaction) TableName() string {
	return "transactions"
}

func (Merch) TableName() string {
	return "merch"
}

func (UserMerch) TableName() string {
	return "user_merch"
}
