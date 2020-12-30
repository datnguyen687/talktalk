package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	authCfgFile string

	authCmd = &cobra.Command{
		Use:   "auth",
		Short: "authorization grpc server",
		RunE: func(cmd *cobra.Command, args []string) error {
			if authCfgFile == "" {
				return errors.New("config is required")
			}

			// viper.SetConfigFile(authCfgFile)
			// if err := viper.ReadInConfig(); err != nil {
			// 	panic(err)
			// }

			// cfg := gateway.Config{
			// 	Debug: viper.GetBool("debug"),
			// 	Port:  viper.GetInt("port"),
			// }

			// runner := gateway.NewHTTPGateWayHandler()
			// if err := runner.Init(&cfg); err != nil {
			// 	panic(err)
			// }

			// return runner.Run()

			return nil
		},
	}
)

func init() {
	authCmd.Flags().StringVar(&authCfgFile, "config", "", "config file")
	authCmd.MarkFlagRequired("config")
}
