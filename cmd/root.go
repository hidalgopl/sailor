package cmd

import (
	"fmt"
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
	rootCmd.PersistentFlags().StringVar(
		&configFile, "config", "", "config file (default is .secureapi.yml)")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName(".secureapi")
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
