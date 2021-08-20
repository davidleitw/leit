package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd *cobra.Command = &cobra.Command{
	Use:   "create",
	Short: "新增一筆行程",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create function")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
