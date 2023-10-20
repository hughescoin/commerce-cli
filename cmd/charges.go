package cli

import (
	"fmt"

	"github.com/coinbase-samples/commerce-sdk-go"

	"github.com/spf13/cobra"
)

var setPriceValue string
var setChargeId string

var chargesCmd = &cobra.Command{
	Use:   "charges",
	Short: "Interact with the charges endpoint",
	Long:  `Interact with the Coinbase Commerce charges endpoint to create and view charges.`,
	Run: func(cmd *cobra.Command, args []string) {

		if client == nil {
			fmt.Println("Please initiatie the client first using `./commerce init`")
			return
		}

		if setPriceValue != "" && setChargeId != "" {
			fmt.Println("Cannot have both a PriceValue and ChargeId")
			return
		}

		if setPriceValue == "" && setChargeId == "" {
			fmt.Println("Please provide either --setprice or --get flag.")
		}

		if setPriceValue != "" {
			chargeReq := commerce.ChargeRequest{
				PricingType: "fixed_price",
				LocalPrice: commerce.LocalPrice{
					Amount:   setPriceValue,
					Currency: "USD",
				},
			}
			resp, err := client.CreateCharge(&chargeReq)
			if err != nil {
				fmt.Printf("Error creating charge: %s ", err)
			}
			fmt.Printf("Charge created successfully! \n %v", resp.Data)

		}

		if setChargeId != "" {
			charge, err := client.GetCharge(setChargeId)
			if err != nil {
				fmt.Printf("Error obtaining charge: %s \n Error: %s", setChargeId, err)
			}

			fmt.Printf("Charge details: %+v \n", charge)
		}

	},
}

func init() {
	rootCmd.AddCommand(chargesCmd)
	chargesCmd.Flags().StringVarP(&setPriceValue, "setPrice", "p", "", "Set the price for the charge")
	chargesCmd.Flags().StringVarP(&setChargeId, "get", "g", "", "Retrieve a charge by its code")

}
