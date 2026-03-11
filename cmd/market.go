package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/victorzakharyev/polymarket_cli/client"
)

var (
	marketID string
)

// marketCmd represents the market command
var marketCmd = &cobra.Command{
	Use:   "market",
	Short: "Get details for a specific market",
	Long:  `Fetches current probabilities and details for a specific market using its condition ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		pmClient := client.NewClient()

		result, err := pmClient.GetMarket(marketID)
		if err != nil {
			printError(err)
			return
		}

		if result == nil {
			printError(fmt.Errorf("market not found for condition ID: %s", marketID))
			return
		}

		printSuccess(result)
	},
}

func init() {
	rootCmd.AddCommand(marketCmd)

	marketCmd.Flags().StringVar(&marketID, "id", "", "Condition ID or Market ID (required)")
	marketCmd.MarkFlagRequired("id")
}
