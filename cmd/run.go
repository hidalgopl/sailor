package cmd

import (
	"os"

	"github.com/hidalgopl/sailor/internal/config"
	"github.com/hidalgopl/sailor/internal/runner"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs SecureAPI security checks for you!",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.GetConf()
		err := runner.Run(conf)
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}
