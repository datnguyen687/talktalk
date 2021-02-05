package usecase

import (
	"talktalk/authentication/handler/dto"
	"talktalk/entities"
)

// Interface ...
type Interface interface {
	RegisterNewUser(*dto.UserDTO) (*entities.User, error)
}
