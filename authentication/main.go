package main

import (
	"fmt"
	"os"
	"talktalk/authentication/handler"
	"talktalk/authentication/usecase"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	log.SetReportCaller(true)
}

func main() {
	dbUsername := os.Getenv("MYSQL_USERNAME")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbAddress := viper.GetString(`db.address`)
	dbDatabase := viper.GetString(`db.database`)

	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local`, dbUsername, dbPassword, dbAddress, dbDatabase)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	uc, err := usecase.NewAuthenticationUsecase(db)
	if err != nil {
		panic(err)
	}

	cfg := handler.Config{
		Port: viper.GetInt(`port`),
	}
	h := handler.NewHTTPAuthentication(cfg, uc)

	log.Error(h.Run())
}
