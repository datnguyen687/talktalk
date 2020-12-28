package mysql

import (
	"talktalk/models"

	"gorm.io/gorm"
)

func (service *sqlService) FilterUser(filter *models.UserFilter) ([]models.User, error) {
	db := service.db.Model(&models.User{})

	if filter.Email != nil {
		db = db.Where(`email=?`, *filter.Email)
	}

	u := []models.User{}

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (service *sqlService) InsertUser(model *models.User) error {
	// db := service.db.Begin()
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		db.Rollback()
	// 	}
	// }()

	// if err := db.Error; err != nil {
	// 	return err
	// }

	// if r := db.Model(&models.User{}).Create(model); r.Error != nil {
	// 	db.Rollback()
	// 	return r.Error
	// }

	// return db.Commit().Error

	return service.db.Transaction(func(tx *gorm.DB) error {
		if err := service.db.Model(&models.User{}).Create(model).Error; err != nil {
			return err
		}

		return nil
	})
}

func (service *sqlService) GetUser(email string) (*models.User, error) {
	db := service.db.Model(&models.User{}).Where(`email=?`, email)

	u := &models.User{}

	if err := db.First(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (service *sqlService) UpdateUser(model *models.User) error {
	db := service.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if r := db.Model(&models.User{}).Where(`email=?`, model.Email).Updates(model); r.Error != nil {
		db.Rollback()
		return r.Error
	}

	return db.Commit().Error
}
