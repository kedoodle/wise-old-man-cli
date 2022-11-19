package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var playersCmd = &cobra.Command{
	Use:   "players <username> [username]...",
	Short: "Update players by username",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("updating players: %+q\n", args)
	},
}

func init() {
	updateCmd.AddCommand(playersCmd)
}
