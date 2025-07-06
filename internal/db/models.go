package db

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;column:id"`
	ClerkID   string    `gorm:"column:clerk_id;unique;not null"`
	Email     string    `gorm:"column:email;unique;not null"`
	FirstName string    `gorm:"column:first_name;not null"`
	LastName  string    `gorm:"column:last_name;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()"`
}

type GoogleAccount struct {
	ClerkUserID  string `gorm:"primaryKey;column:clerk_user_id"`
	AccessToken  string `gorm:"column:access_token"`
	RefreshToken string `gorm:"column:refresh_token"`
	ExpiryDate   int64  `gorm:"column:expiry_date"`
}

type ScheduledEmail struct {
	ID            int       `gorm:"primaryKey;column:id"`
	UserID        string    `gorm:"column:user_id"`
	Sender        string    `gorm:"column:sender;not null"`
	Recipient     string    `gorm:"column:recipient;not null"`
	Subject       string    `gorm:"column:subject;not null"`
	Content       string    `gorm:"column:content;not null"`
	ScheduledDate time.Time `gorm:"column:scheduled_date;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:now()"`
	Status        string    `gorm:"column:status;not null;default:pending"`
	Timezone      string    `gorm:"column:timezone"`
}
