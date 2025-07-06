package scheduled_email_tracker

import (
	"time"

	"gorm.io/gorm"

	"github.com/sirupsen/logrus"

	"github.com/sultantemuruly/schedule-checker-service/internal/db"
	"github.com/sultantemuruly/schedule-checker-service/scripts/helpers"
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
	result := gormDB.Find(&emails)
	if result.Error != nil {
		logrus.Errorf("Failed to query scheduled_emails: %v", result.Error)
		return
	}
	logrus.Infof("Found %d scheduled emails", len(emails))
	for _, email := range emails {
		scheduledDate, err := convert_to_local.ConvertToLocalTimezone(email.ScheduledDate, email.Timezone)
		if err != nil {
			logrus.Errorf("Failed to convert scheduled date to local timezone: %v", err)
			continue
		}

		logrus.Infof("sender: %v, scheduledDate: %v, now: %v", email.Sender, scheduledDate, time.Now())
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