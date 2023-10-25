package cli

import (
	"fmt"

	"github.com/hughescoin/commerce-cli/sdk"
	"github.com/spf13/cobra"
)

var setEventId string
var getAll bool

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Interact with the events endpoint",
	Long:  "Interact with the Coinbase Commerce events endpoint to view event details.",
	Run: func(cmd *cobra.Command, args []string) {

		if sdk.Client == nil {
			fmt.Printf("Please initiate client using `init apiKey <ApiKeyEnvironmentVariable>`")
			return
		}

		if setEventId != "" {

			event, err := sdk.Client.ShowEvent(setEventId)
			if err != nil {
				fmt.Printf("Error retrieving event %s \n Error: %s", setEventId, err)
			}
			fmt.Println(event)
			return
		}

		if getAll {
			allEvents, err := sdk.Client.ListEvents()
			if err != nil {
				fmt.Printf("Error retrieving events %s", err)
				return
			}

			fmt.Println(allEvents)
			return
		}

		fmt.Print("Please provide an eventId to retrieve: `events --get <eventId>`")

	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
	eventsCmd.Flags().StringVarP(&setEventId, "get", "g", "", "Retrieves an event by its id")
	eventsCmd.Flags().BoolVar(&getAll, "all", false, "Retrieve all events")

}
