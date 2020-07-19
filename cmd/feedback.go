package cmd

import (
	"github.com/hidalgopl/sailor/internal/config"
	"github.com/hidalgopl/sailor/internal/feedback"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(feedbackCmd)
}

var feedbackCmd = &cobra.Command{
	Use:   "feedback",
	Short: "Asks you 5 simple questions and send your feedback to us.",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.GetConf()
		buildCfg, err := config.LoadBuildConfig()
		if err != nil {
			logrus.Error(err)
		}
		feedProc := feedback.NewFeedbackProcessor(conf.Username, conf.AccessKey, buildCfg.APIUrl)
		err = feedProc.Process()
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
	},
}
