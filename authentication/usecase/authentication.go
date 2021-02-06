package usecase

import (
	"errors"
	"talktalk/authentication/handler/dto"
	"talktalk/authentication/repository"
	"talktalk/entities"
	"talktalk/utils"
	"time"

	"gorm.io/gorm"
)

// NewAuthenticationUsecase ...
func NewAuthenticationUsecase(db *gorm.DB) (Interface, error) {
	acr, err := repository.NewActivationCodeRepository(db)
	if err != nil {
		return nil, err
	}
	ur, err := repository.NewUserRepository(db)
	if err != nil {
		return nil, err
	}
	return &authenticationUsecase{
		db:                 db,
		activationCodeRepo: acr,
		userRepo:           ur,
	}, nil
}

type authenticationUsecase struct {
	activationCodeRepo entities.ActivationCodeInterface
	userRepo           entities.UserInterface
	db                 *gorm.DB
}

func (au *authenticationUsecase) RegisterNewUser(d *dto.UserDTO) (*entities.User, error) {
	if found, _ := au.userRepo.GetUserByEmail(d.Email); found != nil {
		return nil, errors.New("user already registered")
	}
	db := au.db.Begin()

	userModel := entities.User{
		Status:    entities.UserNotActivated,
		Email:     d.Email,
		Password:  d.Password,
		CreatedAt: time.Now().UTC(),
	}
	user, err := au.userRepo.Create(&userModel)
	if err != nil {
		db.Rollback()
		return nil, err
	}

	code := utils.GenerateActivationCode(entities.ActivationCodeLength)
	now := time.Now()
	expiredAt := now.Add(time.Second * time.Duration(entities.ActivationCodeLifeSpaceInSec))
	activationCodeModel := entities.ActivationCode{
		Code:      code,
		UserID:    user.ID,
		CreatedAt: now,
		ExpiredAt: expiredAt,
	}
	_, err = au.activationCodeRepo.Create(&activationCodeModel)
	if err != nil {
		db.Rollback()
		return nil, err
	}

	if err = db.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}
