package cmd

import (
	"fmt"

	"github.com/RunningIkkyu/leetcode-com-cli/pkg"
	"github.com/spf13/cobra"
)

var idOrTitle string

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&idOrTitle, "id", "i", "", "show question with specific id or title.")
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display question details",
	Long:  `Display question details. With -g/-l/-x, the code template would be auto generated for you.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Leetcode-cli v0.1")
		fmt.Println(args)
		if len(args) != 1 {
			fmt.Println("You should provide exactly one question id/title.")
		}
		res := pkg.GetQuestionDetail(args[0])
		if res == "" {
			fmt.Printf("No such question named: '%s'\n", args[0])
		}
		res = pkg.GetPrettyText(res)
		fmt.Println(res)
		fmt.Println("[debug] the title is :", idOrTitle)
		fmt.Println(cmd.Flags())
	},
}
