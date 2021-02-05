package repository

import "talktalk/entities"

// activationCodeReader ...
type activationCodeReader interface {
}

// activationCodeWriter ...
type activationCodeWriter interface {
	Create(model *entities.ActivationCode) (*entities.ActivationCode, error)
	Update(model *entities.ActivationCode) (*entities.ActivationCode, error)
	Delete(id int) error
	Transaction(func() error) error
}

// ActivationCodeRepositoryInterface ...
type ActivationCodeRepositoryInterface interface {
	activationCodeReader
	activationCodeWriter
}
