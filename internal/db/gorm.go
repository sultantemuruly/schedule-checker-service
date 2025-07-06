package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"github.com/joho/godotenv"
)

func GormConnect(envPath string) (*gorm.DB, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
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