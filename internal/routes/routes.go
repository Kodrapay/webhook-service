package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodra-pay/webhook-service/internal/config"
	"github.com/kodra-pay/webhook-service/internal/handlers"
	"github.com/kodra-pay/webhook-service/internal/repositories"
	"github.com/kodra-pay/webhook-service/internal/services"
)

func Register(app *fiber.App, cfg config.Config, repo *repositories.WebhookRepository) {
	health := handlers.NewHealthHandler(cfg.ServiceName)
	health.Register(app)

	svc := services.NewWebhookService(repo)
	h := handlers.NewWebhookHandler(svc)
	app.Post("/webhooks/providers/:provider", h.Receive)
}
