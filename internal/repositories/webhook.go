package repositories

import "log"

type WebhookRepository struct {
	dsn string
}

func NewWebhookRepository(dsn string) *WebhookRepository {
	log.Printf("WebhookRepository using DSN: %s", dsn)
	return &WebhookRepository{dsn: dsn}
}

// TODO: implement persistence for webhook logs and retries.
