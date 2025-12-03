package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kodra-pay/webhook-service/internal/services"
)

type WebhookHandler struct {
	svc *services.WebhookService
}

func NewWebhookHandler(svc *services.WebhookService) *WebhookHandler {
	return &WebhookHandler{svc: svc}
}

func (h *WebhookHandler) Receive(c *fiber.Ctx) error {
	provider := c.Params("provider")
	body := c.Body()
	headers := map[string]string{}
	c.Request().Header.VisitAll(func(k, v []byte) {
		headers[string(k)] = string(v)
	})
	return c.JSON(h.svc.Receive(c.Context(), provider, headers, body))
}
