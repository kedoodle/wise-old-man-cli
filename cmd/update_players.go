package cmd

import (
	"fmt"
	"log"

	wise "github.com/kedoodle/wise-old-man"
	"github.com/spf13/cobra"
)

var playersCmd = &cobra.Command{
	Use:   "players <username> [username]...",
	Short: "Update players by username",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("updating players: %+q\n", args)
		wise := wise.New()
		ctx := cmd.Context()
		for _, username := range args {
			_, err := wise.UpdatePlayer(ctx, username)
			if err != nil {
				log.Printf("error updating player %q: %v", username, err)
			}
			log.Printf("updated player %q", username)
		}
	},
}

func init() {
	updateCmd.AddCommand(playersCmd)
}
