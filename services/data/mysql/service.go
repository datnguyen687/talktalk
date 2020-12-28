package mysql

import (
	"fmt"
	"talktalk/models"
	dataservice "talktalk/services/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// sqlService ...
type sqlService struct {
	db *gorm.DB
}

// NewMySQLDataService ...
func NewMySQLDataService(cfg Config) (dataservice.ServiceInterface, error) {
	service := &sqlService{}

	if err := service.init(cfg); err != nil {
		return nil, err
	}

	return service, nil
}

func (service *sqlService) init(cfg Config) error {
	connectionStr := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local`,
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	var err error
	if service.db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       connectionStr, // data source name
		DefaultStringSize:         256,           // default size for string fields
		DisableDatetimePrecision:  true,          // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,          // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,          // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,         // auto configure based on currently MySQL version
	}), &gorm.Config{}); err != nil {
		return err
	}

	return service.db.AutoMigrate(
		&models.User{},
		&models.UserStatus{},
		&models.ActivationCode{},
	)
}
