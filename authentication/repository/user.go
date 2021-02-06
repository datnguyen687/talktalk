package repository

import (
	"talktalk/entities"

	"gorm.io/gorm"
)

// NewUserRepository ...
func NewUserRepository(db *gorm.DB) (entities.UserInterface, error) {
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		return nil, err
	}

	ur := &userRepository{
		db: db,
	}

	return ur, nil
}

type userRepository struct {
	db *gorm.DB
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

func (ur *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	db := ur.db.Model(&entities.User{})

	data := &entities.User{}
	if db = db.Where(`email=?`, email).First(data); db.Error != nil {
		return nil, db.Error
	}

	return data, nil
}
