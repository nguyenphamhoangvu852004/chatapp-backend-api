package utils

import (
	"chapapp-backend-api/global"
	"fmt"
	"net/smtp"
	"strings"

	"go.uber.org/zap"
)

const (
	Host          = "smtp.gmail.com"
	Port          = 587
	User          = "nguyenphamhoangvu852004@gmail.com"
	Password      = "ubzr bngs eyrc btrk"
	AdminReceiver = "nguyenphamhoangvu852004@gmail.com"
)

type EmailAdress struct {
	Address string
	Name    string
}
type Mail struct {
	From    EmailAdress
	To      []string
	Body    string
	Subject string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
func SendTextEmailOTP(to []string, from string, body string) error {
	contentEmail := Mail{
		EmailAdress{
			Address: from, Name: "test"},
		to,
		body,
		"OTP",
	}

	messageMail := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", User, Password, Host)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", Host, Port), auth, from, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Send mail fail", zap.Error(err))
		return err
	}

	return nil
}
