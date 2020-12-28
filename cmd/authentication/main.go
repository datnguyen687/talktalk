package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"talktalk/handlers/authentication"
	mysqlDS "talktalk/services/data/mysql"
	sendGrid "talktalk/services/email/sendgrid"
	"time"

	"github.com/spf13/viper"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()
	if configPath == nil || *configPath == "" {
		flag.Usage()
		return
	}

	viper.SetConfigFile(*configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	cfg := authentication.ServerConfig{
		Port: viper.GetInt(`port`),
		SQL: mysqlDS.Config{
			DbName:   viper.GetString(`mysql.db_name`),
			Host:     viper.GetString(`mysql.host`),
			Password: viper.GetString(`mysql.password`),
			Port:     viper.GetInt(`mysql.port`),
			Username: viper.GetString(`mysql.user_name`),
		},
		Email: sendGrid.Config{
			APIKey: os.Getenv(`sendgrid_api_key`),
			Email:  os.Getenv(`email`),
		},
	}
	server, err := authentication.NewAuthorizationServer(cfg)

	if err != nil {
		log.Println(err)
		return
	}

	log.Fatalln(server.Run())
}
