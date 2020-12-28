package email

// ServiceInterface ...
type ServiceInterface interface {
	SendActivationCode(email, code string) error
}
