package cmd

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

var playersCmd = &cobra.Command{
	Use:   "players <username> [username]...",
	Short: "Update players by username",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("updating players: %+q\n", args)
		client := http.Client{}
		for _, username := range args {
			url := url.URL{
				Scheme: "https",
				Host:   "api.wiseoldman.net",
				Path:   fmt.Sprintf("/v2/players/%s", username),
			}
			_, err := client.Post(url.String(), "application/json", nil)
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
