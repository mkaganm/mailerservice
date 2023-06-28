package mail

import (
	"log"
	"mailerservice/pkg/config"
	"net/smtp"
	"strconv"
)

type Mail struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func SendMail(mail Mail) error {

	smtpHost := config.EnvConfigs.SmtpHost
	smtpPort := config.EnvConfigs.SmtpPort
	username := config.EnvConfigs.MailerSenderAddress
	password := config.EnvConfigs.MailerSenderPassword
	auth := smtp.PlainAuth("", username, password, smtpHost)

	message := []byte("Subject: " + mail.Subject + "\r\n" + "\r\n" + mail.Body)

	err := smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, mail.From, mail.To, message)
	if err != nil {
		log.Default().Println("Error while sending email: " + err.Error())
		return err
	}

	log.Default().Println("Email sent successfully!")
	return nil

}
