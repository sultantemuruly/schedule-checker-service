package scheduled_email_tracker

import (
	"time"

	"gorm.io/gorm"

	"github.com/sirupsen/logrus"

	"github.com/sultantemuruly/schedule-checker-service/internal/db"
	// "github.com/sultantemuruly/schedule-checker-service/scripts/helpers"
)

func waitUntilNext5MinuteMark() {
	now := time.Now()
	minutes := now.Minute()
	seconds := now.Second()

	waitMinutes := 5 - (minutes % 5)
	if waitMinutes == 0 && seconds == 0 {
		return // Already at a 5-minute mark
	}

	// If we're already at a 5-minute mark but not at 0 seconds, wait 5 minutes
	if waitMinutes == 0 {
		waitMinutes = 5
	}
	waitDuration := time.Duration(waitMinutes)*time.Minute - time.Duration(seconds)*time.Second - time.Duration(now.Nanosecond())*time.Nanosecond
	logrus.Infof("Waiting %v until next 5-minute mark...", waitDuration)
	time.Sleep(waitDuration)
}

func logScheduledEmails(gormDB *gorm.DB) {
	var emails []db.ScheduledEmail
	result := gormDB.Where("status = ?", "pending").Find(&emails)
	if result.Error != nil {
		logrus.Errorf("Failed to query scheduled_emails: %v", result.Error)
		return
	}
	logrus.Infof("Found %d pending scheduled emails", len(emails))
	nowUTC := time.Now().UTC()

	count_to_send := 0
	for _, email := range emails {
		if !email.ScheduledDate.After(nowUTC) { // scheduledDate <= nowUTC
			logrus.Infof("UserID: %s, ScheduledDate: %s, Now(UTC): %s",
				email.UserID,
				email.ScheduledDate.UTC().Format(time.RFC3339),
				nowUTC.Format(time.RFC3339),
			)
			count_to_send++
		}
	}

	if count_to_send == 0 {
		logrus.Infof("No emails to send")
		return
	}
}

func TrackScheduledEmails(gormDB *gorm.DB) {
	waitUntilNext5MinuteMark()

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		logScheduledEmails(gormDB)
		<-ticker.C
	}
} 