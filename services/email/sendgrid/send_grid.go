package email

import (
	"errors"
	"fmt"
	"net/http"
	"talktalk/services/email"

	goSendGrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGrid ...
type SendGrid struct {
	client *goSendGrid.Client
	email  string
}

// NewEmailService ...
func NewEmailService(cfg Config) email.ServiceInterface {
	return &SendGrid{
		client: goSendGrid.NewSendClient(cfg.APIKey),
		email:  cfg.Email,
	}
}

// SendActivationCode ...
func (sg *SendGrid) SendActivationCode(email, code string) error {
	from := mail.NewEmail("Example User", sg.email)
	subject := "TalkTalk activation code"
	to := mail.NewEmail("Example User", email)
	plainTextContent := fmt.Sprintf("your activation code is %s - expired after 15 minutes", code)
	// htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")

	response, err := sg.client.Send(message)

	if err != nil {
		return err
	}
	if response != nil && response.StatusCode != http.StatusOK {
		return errors.New(response.Body)
	}

	return nil
}
