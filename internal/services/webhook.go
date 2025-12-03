package services

import (
	"context"

	"github.com/kodra-pay/webhook-service/internal/repositories"
)

type WebhookService struct {
	repo *repositories.WebhookRepository
}

func NewWebhookService(repo *repositories.WebhookRepository) *WebhookService {
	return &WebhookService{repo: repo}
}

func (s *WebhookService) Receive(_ context.Context, provider string, headers map[string]string, body []byte) map[string]interface{} {
	return map[string]interface{}{
		"provider": provider,
		"status":   "received",
		"headers":  headers,
		"length":   len(body),
	}
}
