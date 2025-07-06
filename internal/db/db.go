package db

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func Connect(ctx context.Context, envPath string) (*pgx.Conn, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		logrus.Warn("No .env file found or failed to load .env")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, ErrNoDatabaseURL
	}

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

var ErrNoDatabaseURL = &NoDatabaseURLError{}

type NoDatabaseURLError struct{}

func (e *NoDatabaseURLError) Error() string {
	return "DATABASE_URL is not set in environment"
}
