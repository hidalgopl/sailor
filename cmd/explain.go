package cmd

import (
	"github.com/hidalgopl/sailor/internal/sectests"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(explainCmd)
}

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Creates config template for sailor.",
	ValidArgs: sectests.SEC_TEST_KEYS,
	Run: func(cmd *cobra.Command, args []string) {
		err := cobra.OnlyValidArgs(cmd, args)
		if err != nil {
			logrus.Errorf("%v, Valid args are: %v", err, sectests.SEC_TEST_KEYS)

			os.Exit(1)
		}
		sectests.PrintExplanation(args)
		os.Exit(0)
	},
}
