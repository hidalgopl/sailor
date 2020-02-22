package cmd

import (
	"github.com/hidalgopl/sailor/internal/config"
	"github.com/hidalgopl/sailor/internal/feedback"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(feedbackCmd)
}

var feedbackCmd = &cobra.Command{
	Use:   "feedback",
	Short: "Asks you 5 simple questions and send your feedback to us.",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.GetConf()
		feedProc := feedback.NewFeedbackProcessor(conf.Username, conf.AccessKey)
		err := feedProc.Process()
		if err != nil {
			logrus.Error(err)
		}
	},
}
