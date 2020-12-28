package data

import (
	"talktalk/models"
)

// ServiceInterface ...
type ServiceInterface interface {
	FilterUser(filter *models.UserFilter) ([]models.User, error)
	GetUser(email string) (*models.User, error)
	InsertUser(model *models.User) error
	InsertActivationCode(model *models.ActivationCode) error
	UpdateUser(model *models.User) error

	GetActivationCode(email string) (*models.ActivationCode, error)
	DeleteActivationCode(email, code string) error
}
