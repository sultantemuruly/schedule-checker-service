package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/sultantemuruly/schedule-checker-service/internal/db"
)

func main() {
	ctx := context.Background()
	envPath := ".env"
	conn, err := db.Connect(ctx, envPath)
	if err != nil {
		logrus.Fatalf("Database connection failed: %v", err)
	}
	defer conn.Close(ctx)

	logrus.Info("Successfully connected to the database!")
}
