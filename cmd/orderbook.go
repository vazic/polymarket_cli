package cmd

import (
	"github.com/spf13/cobra"
	"github.com/victorzakharyev/polymarket_cli/client"
)

var (
	orderbookTokenID string
)

// orderbookCmd represents the orderbook command
var orderbookCmd = &cobra.Command{
	Use:   "orderbook",
	Short: "Get the orderbook for a specific token",
	Long:  `Fetches the CLOB API order book for a specific outcome token ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		pmClient := client.NewClient()

		result, err := pmClient.GetOrderbook(orderbookTokenID)
		if err != nil {
			printError(err)
			return
		}

		printSuccess(result)
	},
}

func init() {
	rootCmd.AddCommand(orderbookCmd)

	orderbookCmd.Flags().StringVar(&orderbookTokenID, "token-id", "", "The outcome token ID (required)")
	orderbookCmd.MarkFlagRequired("token-id")
}
