package controllers

import (
	"talktalk/models"
)

// ControllerInterface ...
type ControllerInterface interface {
	SignUp(dto *models.UserDTO) error
	ActivateUser(email, code string) error
}
