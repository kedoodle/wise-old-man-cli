package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var playersCmd = &cobra.Command{
	Use:   "players",
	Short: "Update players by username",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update players called")
	},
}

func init() {
	updateCmd.AddCommand(playersCmd)
}
