package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Connect(ctx context.Context) (*pgx.Conn, error) {
	// Only load .env in local/dev, not in production
	if _, err := os.Stat(".env"); err == nil {
		_ = godotenv.Load(".env")
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
