package cmd

import (
	"fmt"
	"github.com/hidalgopl/sailor/internal/config"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "sailor",
	Short: "Sailor worker svc",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	usageMsg := "config file (default is " + config.SECUREAPI_FILE + ")"
	rootCmd.PersistentFlags().StringVar(
		&configFile, "config", "", usageMsg)
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName(config.SECUREAPI_FILENAME)
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/sailor")
		viper.AddConfigPath("$HOME/.sailor")
	}
	viper.SetEnvPrefix("secureapi")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config: %v\n", err)
		os.Exit(1)
	}
}
