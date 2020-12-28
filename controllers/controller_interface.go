package controllers

import (
	"talktalk/models"
)

// ControllerInterface ...
type ControllerInterface interface {
	SignUp(dto *models.UserDTO) error
	ActivateUser(email, code string) error
	ResendCode(email string) (string, error)
}
