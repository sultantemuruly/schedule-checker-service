package main

import (
	"github.com/sirupsen/logrus"
	"github.com/sultantemuruly/schedule-checker-service/internal/db"
	scheduled_email_tracker "github.com/sultantemuruly/schedule-checker-service/internal/tracker"
)

func main() {
	gormDB, err := db.GormConnect()
	if err != nil {
		logrus.Fatalf("Failed to connect with GORM: %v", err)
	}

	scheduled_email_tracker.TrackScheduledEmails(gormDB)
}
