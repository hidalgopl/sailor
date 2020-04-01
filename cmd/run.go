package cmd

import (
	"net/http"

	"github.com/hidalgopl/sailor/internal/auth"
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
	Short: "Runs your secureapi test session!",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.GetConf()
		authenticator := auth.Authenticator{
			Username:   conf.Username,
			AccessKey:  conf.AccessKey,
			URL:        "http://localhost:8072/auth",
			HttpClient: &http.Client{},
		}

		isAllowed, userID, err := authenticator.DoAuth()
		if err != nil {
			logrus.Errorf("Something's wrong on our end, apologies: %s", err)
		}
		if isAllowed {
			err := runner.Run(conf, userID)
			if err != nil {
				logrus.Error(err)
			}

		} else {
			logrus.Errorf("Can't authenticate user %s", conf.Username)
		}

	},
}
