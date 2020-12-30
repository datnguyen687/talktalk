package cmd

import (
	"errors"
	"talktalk/internal/handlers/gateway"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	gatewayCfgFile string

	gatewayCmd = &cobra.Command{
		Use:   "gateway",
		Short: "http gateway",
		RunE: func(cmd *cobra.Command, args []string) error {
			if gatewayCfgFile == "" {
				return errors.New("config is required")
			}

			viper.SetConfigFile(gatewayCfgFile)
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

			return runner.Run()
		},
	}
)

func init() {
	gatewayCmd.Flags().StringVar(&gatewayCfgFile, "config", "", "config file")
	gatewayCmd.MarkFlagRequired("config")
}
