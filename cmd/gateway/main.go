package main

import (
	"flag"
	gateway "talktalk/internal/handlers/gateway"

	logrus "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
)

func main() {
	configPath := flag.String("config", "", "path to config")
	flag.Parse()
	if configPath == nil || *configPath == "" {
		flag.Usage()
		return
	}

	viper.SetConfigFile(*configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	cfg := gateway.Config{
		Debug: viper.GetBool("debug"),
		Port:  viper.GetInt("port"),
	}

	runner := gateway.NewHTTPGateWayHandler()
	if err := runner.Init(&cfg); err != nil {
		panic(err)
	}

	logrus.Fatalln(runner.Run())
}
