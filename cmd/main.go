package main

import (
	"github.com/gofiber/fiber/v2"
	"mailerservice/pkg/config"
	"mailerservice/pkg/controller"
	"mailerservice/pkg/utils"
)

func main() {
	// Init env configs
	config.InitEnvConfigs()

	app := fiber.New(fiber.Config{
		//ReadTimeout:   time.Second * 15,
		//WriteTimeout:  time.Second * 15,
		Concurrency:  10,
		ServerHeader: "user_service_v1",
		AppName:      "user_service_v1",
	})

	// Register routes
	controller.RegisterRoutes(app)

	// Listen on port
	err := app.Listen(config.EnvConfigs.LocalServerPort)
	utils.FatalErr("Error while serving the api", err)
}
