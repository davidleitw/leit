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

	var date string
	var title string
	var fulldate string
	var start, end string

	createCmd.Flags().StringVarP(&title, "title", "t", "", "行程標題 [必填]")
	createCmd.Flags().StringVarP(&start, "start", "s", "", "開始時間")
	createCmd.Flags().StringVarP(&end, "end", "e", "", "結束時間")
	createCmd.Flags().StringVarP(&date, "date", "d", "", "縮寫日期, 格式為: \"04-22\"")
	createCmd.Flags().StringVarP(&fulldate, "full_date", "D", "", "完整日期, 格式為: \"2019-06-16\"")

	_ = createCmd.MarkFlagRequired("title")
}
