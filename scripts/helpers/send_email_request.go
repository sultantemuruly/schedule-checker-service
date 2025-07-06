package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sultantemuruly/schedule-checker-service/internal/db"
	"net/http"
	"time"
)

func SendEmailRequest(email db.ScheduledEmail) error {
	url := "https://www.silhai.com/api/send"
	payload := map[string]string {
		"sender": email.Sender,
		"recipient": email.Recipient,
		"subject": email.Subject,
		"content": email.Content,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send email: %s", resp.Status)
	}

	return nil
}	