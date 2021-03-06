package repository

import (
	"talktalk/entities"

	"gorm.io/gorm"
)

// NewActivationCodeRepository ...
func NewActivationCodeRepository(db *gorm.DB) (entities.ActivationCodeInterface, error) {
	if err := db.AutoMigrate(&entities.ActivationCode{}); err != nil {
		return nil, err
	}

	acr := &activationCodeRepository{
		db: db,
	}

	return acr, nil
}

type activationCodeRepository struct {
	db *gorm.DB
}

func (acr *activationCodeRepository) Create(model *entities.ActivationCode) (*entities.ActivationCode, error) {
	db := acr.db

	db = db.Model(&entities.ActivationCode{}).Create(model)

	if db.Error != nil {
		return nil, db.Error
	}
	return model, nil
}

func (acr *activationCodeRepository) Update(model *entities.ActivationCode) (*entities.ActivationCode, error) {
	db := acr.db

	db = db.Model(&entities.ActivationCode{}).Updates(model)

	if db.Error != nil {
		return nil, db.Error
	}
	return model, nil
}

func (acr *activationCodeRepository) Delete(id int) error {
	db := acr.db.Begin()

	db = db.Model(&entities.ActivationCode{}).Where(`id=?`, id).Delete(&entities.ActivationCode{})

	return db.Error
}
