package main

import (
	"context"
	"fmt"
	"log"

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

	rows, err := conn.Query(ctx, "SELECT * FROM scheduled_emails")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(values) // Prints all column values for each row
	}
}
