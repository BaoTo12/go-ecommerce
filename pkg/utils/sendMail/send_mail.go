package sendmail

import (
	"bytes"
	"fmt"
	"html/template"
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

func SendTemplateMailOTP(to []string, from string, nameTemplate string, dataTemplate map[string]interface{}) error {
	htmlBody, err := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		// If an error occurs during template processing, return the error.
		return err
	}

	// Call a placeholder `send` function to send the email with the generated HTML body.
	// This function would contain your SMTP client or mailer library code.
	return send(to, from, htmlBody)
}

func getMailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)

	t := template.Must(template.New(nameTemplate).ParseFiles("templates-email/" + nameTemplate))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}

	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "Test",
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
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
