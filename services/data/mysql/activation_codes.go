package mysql

import "talktalk/models"

func (sql *sqlService) InsertActivationCode(model *models.ActivationCode) error {
	db := sql.db.Model(&models.ActivationCode{}).Begin()

	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	db = sql.db.Create(model)
	if db.Error != nil {
		db.Rollback()
		return db.Error
	}

	return db.Commit().Error
}

func (sql *sqlService) GetActivationCode(email string) (*models.ActivationCode, error) {
	db := sql.db.Model(&models.ActivationCode{}).Where(`user_email=?`, email)

	ac := &models.ActivationCode{}
	if err := db.First(ac).Error; err != nil {
		return nil, err
	}

	u, err := sql.GetUser(email)
	if err != nil {
		return nil, err
	}
	ac.User = *u

	return ac, nil
}

func (sql *sqlService) DeleteActivationCode(email, code string) error {
	db := sql.db.Model(&models.ActivationCode{}).Begin()

	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	db = sql.db.Where(`user_email=?`, email).Where(`code=?`, code).Delete(&models.ActivationCode{})
	if db.Error != nil {
		db.Rollback()
		return db.Error
	}

	return db.Commit().Error
}
