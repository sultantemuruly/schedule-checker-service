package db

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormConnect() (*gorm.DB, error) {
	// Only load .env in local/dev, not in production
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load(".env")
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, ErrNoDatabaseURL
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
