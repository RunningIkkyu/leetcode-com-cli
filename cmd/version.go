package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of leetcode-com-cli",
	Long:  `All software has versions. This is leetcode's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Leetcode-cli v0.1")
	},
}
