package repository

import (
	"talktalk/entities"

	"gorm.io/gorm"
)

// NewUserRepository ...
func NewUserRepository(db *gorm.DB) entities.UserInterface {
	ur := &userRepository{
		db: db,
	}

	return ur
}

type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) Transaction(f func() error) error {
	db := ur.db.Begin()

	err := f()

	if err != nil {
		db.Rollback()
		return err
	}

	if err = db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}

func (ur *userRepository) Create(model *entities.User) (*entities.User, error) {
	db := ur.db

	db = db.Model(&entities.User{}).Create(model)

	if db.Error != nil {
		return nil, db.Error
	}
	return model, nil
}

func (ur *userRepository) Update(model *entities.User) (*entities.User, error) {
	db := ur.db

	db = db.Model(&entities.User{}).Updates(model)

	if db.Error != nil {
		return nil, db.Error
	}
	return model, nil
}

func (ur *userRepository) Delete(email string) error {
	db := ur.db.Begin()

	db = db.Model(&entities.User{}).Where(`email=?`, email).Delete(&entities.User{})

	return db.Error
}
