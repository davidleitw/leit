package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var doc = strings.Join([]string{
	"leit is a tool can help you manage your work in command line.",
}, "\n")

var rootCmd = &cobra.Command{
	Use:   "leit",
	Short: "A calendar application in command line.",
	Long:  doc,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
