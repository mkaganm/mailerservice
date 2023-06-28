package controller

import (
	"github.com/gofiber/fiber/v2"
	"mailerservice/pkg/services"
)

// RegisterRoutes registers all routes for the API
func RegisterRoutes(app *fiber.App) {

	routes := app.Group("/api/v1/mailer")

	routes.Post("/send-mail", services.SendMail)

}
