package cmd

import (
	"fmt"
	"github.com/hidalgopl/sailor/internal/config"
	"github.com/spf13/cobra"
)



func init() {
	rootCmd.AddCommand(configCmd)
}


var configCmd = &cobra.Command{
	Use: "config",
	Short: "Print the loaded config",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.GetConf()
		fmt.Println(conf.PrettyPrint())
	},
}
