package dto

// UserDTO ...
type UserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SelfValidate ...
func (UserDTO) SelfValidate() error {
	return nil
}
