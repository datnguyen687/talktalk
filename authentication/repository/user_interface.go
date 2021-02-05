package repository

import "talktalk/entities"

// userReader ...
type userReader interface {
	GetUserByEmail(email string) (*entities.User, error)
}

// UserWriter ...
type userWriter interface {
	Create(model *entities.User) (*entities.User, error)
	Update(model *entities.User) (*entities.User, error)
	Delete(email string) error
	Transaction(func() error) error
}

// UserRepositoryInterface ...
type UserRepositoryInterface interface {
	userReader
	userWriter
}
