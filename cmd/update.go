package cmd

import "github.com/spf13/cobra"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update players by username, group, or competition",
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
