package services

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"mailerservice/pkg/config"
	"mailerservice/pkg/mail"
)

type sendMailRequest struct {
	To      []string `json:"to" validate:"required"`
	Subject string   `json:"subject" validate:"required"`
	Body    string   `json:"body" validate:"required"`
}

type sendMailResponse struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

// SendMail sends an email to the user
func SendMail(c *fiber.Ctx) error {

	request := c.Request()

	var req sendMailRequest

	err := json.Unmarshal(request.Body(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "bad request",
			"error":  err.Error(),
		})
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "bad request",
			"errors": validationErrors,
		})
	}

	mailInfo := mail.Mail{
		From:    config.EnvConfigs.MailerSenderAddress,
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	}
	err = mail.SendMail(mailInfo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	resp := sendMailResponse{
		Status: "success",
		Data:   map[string]interface{}{},
	}

	resp.Data["infos"] = mailInfo
	return c.Status(fiber.StatusOK).JSON(resp)
}
