package mysql

import (
	"fmt"
	"os"
	"talktalk/internal/database"
	"talktalk/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gormSQL struct {
	db *gorm.DB
}

// NewGormSQL ...
func NewGormSQL() database.Interface {
	return &gormSQL{}
}

func (gsql *gormSQL) Init(cfg interface{}) error {
	addr := os.Getenv("TALKTALK_MYSQL_HOST")
	user := os.Getenv("TALKTALK_MYSQL_USER")
	password := os.Getenv("TALKTALK_MYSQL_PASSWORD")
	dbName := os.Getenv("TALKTALK_MYSQL_DB_NAME")
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		user,
		password,
		addr,
		dbName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	if err = db.AutoMigrate(
		&models.UserStatus{},
		&models.User{},
	); err != nil {
		return err
	}

	gsql.db = db

	return nil
}
