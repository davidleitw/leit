package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "1.0.0"

var versionCmd *cobra.Command = &cobra.Command{
	Use:   "version",
	Short: "Prints leit's version number.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
