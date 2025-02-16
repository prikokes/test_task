package integrated_tests

import (
	"avito_internship_task/internal/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"runtime"
	"testing"
)

func setupTestDB(t *testing.T) (*gorm.DB, error) {
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "../..")
	envPath := filepath.Join(projectRoot, ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		t.Fatalf("Failed to load .env: %v", err)
	}

	db := utils.InitDB()

	if _, err := db.DB().Exec("TRUNCATE users, merch, user_merch, transactions RESTART IDENTITY CASCADE"); err != nil {
		log.Printf("Ошибка при очистке БД: %v", err)
	}

	return db, nil
}
