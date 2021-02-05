package usecase

import (
	"talktalk/authentication/handler/dto"
	"talktalk/authentication/repository"
	"talktalk/entities"
	"time"
)

// NewAuthenticationUsecase ...
func NewAuthenticationUsecase(ui repository.UserRepositoryInterface, aci repository.ActivationCodeRepositoryInterface) Interface {
	return &authenticationUsecase{
		userRepository:           ui,
		activationCodeRepository: aci,
	}
}

type authenticationUsecase struct {
	userRepository           repository.UserRepositoryInterface
	activationCodeRepository repository.ActivationCodeRepositoryInterface
}

func (au *authenticationUsecase) RegisterNewUser(d *dto.UserDTO) (*entities.User, error) {
	_, err := au.userRepository.GetUserByEmail(d.Email)
	if err != nil {
		return nil, err
	}

	var data *entities.User
	au.userRepository.Transaction(func() error {
		model := entities.User{
			Email:     d.Email,
			Password:  d.Password,
			Status:    entities.UserNotActivated,
			CreatedAt: time.Now().UTC(),
		}
		data, err = au.userRepository.Create(&model)

		return err
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}
