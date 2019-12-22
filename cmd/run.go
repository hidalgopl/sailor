package cmd

import (
	"fmt"
	"github.com/hidalgopl/sailor/internal/auth"
	"github.com/hidalgopl/sailor/internal/config"
	"github.com/hidalgopl/sailor/internal/runner"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs your secureapi test session!",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.GetConf()
		fmt.Printf("Running test suite for user %s\n", conf.Username)
		authenticator := auth.Authenticator{
			Username:  conf.Username,
			AccessKey: conf.AccessKey,
			Url:       "http://localhost:8072/auth",
		}
		isAllowed, msg, userId := authenticator.DoAuth()
		if isAllowed {
			runner.Run(conf, userId)

		} else {
			fmt.Println(msg)
			fmt.Printf("Can't authenticate user %s", conf.Username)
		}

	},
}
