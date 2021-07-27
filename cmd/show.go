package cmd

import (
	"fmt"

	"github.com/RunningIkkyu/leetcode-com-cli/pkg"
	"github.com/spf13/cobra"
)

var idOrTitle string
var showLanguage string

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&idOrTitle, "id", "i", "", "show question with specific id or title.")
	showCmd.Flags().StringVarP(&showLanguage, "language", "l", "", "set question language. zh/en")
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display question details",
	Long:  `Display question details. With -g/-l/-x, the code template would be auto generated for you.`,
	Run: func(cmd *cobra.Command, args []string) {
		showLanguage = pkg.StandardedLanguage(showLanguage)
		var titleSlug string
		switch len(args) {
		case 0:
			m := pkg.GetTodayQuestionInfo()
			titleSlug = m["titleSlug"]
		case 1:
			titleSlug = args[0]
		default:
			fmt.Println("You should provide exactly one question id/title.")
			return
		}
		pkg.PrettyPrintQuestionByTitleSlug(titleSlug, showLanguage)
	},
}
