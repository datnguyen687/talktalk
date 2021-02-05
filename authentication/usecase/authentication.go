package usecase

import (
	"talktalk/entities"
)

// NewAuthenticationUsecase ...
func NewAuthenticationUsecase(ui entities.UserInterface, aci entities.ActivationCodeInterface) Interface {
	return &authenticationUsecase{
		userInterface:           ui,
		activationCodeInterface: aci,
	}
}

type authenticationUsecase struct {
	userInterface           entities.UserInterface
	activationCodeInterface entities.ActivationCodeInterface
}
