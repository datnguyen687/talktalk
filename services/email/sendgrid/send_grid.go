package email

import (
	"fmt"
	"talktalk/services/email"

	goSendGrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGrid ...
type SendGrid struct {
	client *goSendGrid.Client
}

// NewEmailService ...
func NewEmailService(cfg Config) email.ServiceInterface {
	return &SendGrid{
		client: goSendGrid.NewSendClient(cfg.APIKey),
	}
}

// SendActivationCode ...
func (sg *SendGrid) SendActivationCode(email, code string) error {
	from := mail.NewEmail("Example User", "test@example.com")
	subject := "TalkTalk activation code"
	to := mail.NewEmail("Example User", email)
	plainTextContent := fmt.Sprintf("your activation code is %s - expired after 15 minutes", code)
	// htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")

	_, err := sg.client.Send(message)

	return err
}
