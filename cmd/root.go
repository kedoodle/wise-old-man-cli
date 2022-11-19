package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wise-old-man-cli",
	Short: "Talk-to Wise Old Man",
	Long: `Talk-to Wise Old Man
Command Dionysius to do your bidding.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
