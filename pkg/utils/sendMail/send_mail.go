package sendmail

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/BaoTo12/go-ecommerce/global"
	"go.uber.org/zap"
)

const (
	SMTPHost     = ""
	SMTPPort     = ""
	SMTPUserName = ""
	SMTPPassword = ""
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress `json:"sender"`
	To      []string     `json:"receiver"`
	Subject string       `json:"subject"`
	Body    string       `json:"body"`
}

func BuildMessage(mail Mail) (msg string) {
	msg += fmt.Sprintf("From %s \r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s \r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("Body: %s\r\n", mail.Body)
	return msg
}

func SendTextMailOTP(to []string, from string, otp string) error {
	contentEmail := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "Test",
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp),
	}

	messageMail := BuildMessage(contentEmail)
	// send SMTP
	// smtp.PlainAuth(identity, username, password, host)
	auth := smtp.PlainAuth("", SMTPUserName, SMTPPassword, SMTPHost)
	address := SMTPHost + SMTPPort

	err := smtp.SendMail(address, auth, from, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Email send fail in util package", zap.Error(err))
		return err
	}
	return nil
}
