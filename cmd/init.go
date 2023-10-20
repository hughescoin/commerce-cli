package cli

import (
	"fmt"
	"net/http"

	"github.com/coinbase-samples/commerce-sdk-go"
	"github.com/spf13/cobra"
)

var client *commerce.Client

func initClient(apiKeyVariable string) (*commerce.Client, error) {

	if client != nil {
		return client, nil
	}

	creds, err := commerce.ReadEnvCredentials(apiKeyVariable)
	if err != nil {
		fmt.Printf("Error reading environmental variable: %s", err)
	}

	client = commerce.NewClient(creds, http.Client{})
	return client, nil
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Commerce SDK Client",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKeyVar, err := cmd.Flags().GetString("apiKey")
		if err != nil {
			return err
		}

		_, err = initClient(apiKeyVar)
		if err != nil {
			return fmt.Errorf("Error initializing client: %s", err)
		}

		fmt.Println("Client initialized successfully using environment variable:", apiKeyVar)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().String("apiKey", "", "Environment variable that holds the API Key for the Commerce SDK client")
	initCmd.MarkFlagRequired("apiKey")
}
