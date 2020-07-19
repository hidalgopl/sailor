package cmd

import (
	"fmt"
	"github.com/hidalgopl/sailor/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(initConfigCmd)
}

var initConfigCmd = &cobra.Command{
	Use:   "init-config",
	Short: "Creates config template for sailor.",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.CreateConfigTemplate()
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
		cwd, err := os.Getwd()
		fmt.Printf("Config template created: %s/%s", cwd, config.SECUREAPI_FILE)
		os.Exit(0)
	},
}

