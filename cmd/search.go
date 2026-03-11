package cmd

import (
	"github.com/spf13/cobra"
	"github.com/victorzakharyev/polymarket_cli/client"
)

var (
	searchQuery string
	searchLimit int
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for markets on Polymarket",
	Long:  `Queries the Gamma API for markets matching the text query.`,
	Run: func(cmd *cobra.Command, args []string) {
		pmClient := client.NewClient()

		results, err := pmClient.SearchMarkets(searchQuery, searchLimit)
		if err != nil {
			printError(err)
			return
		}

		printSuccess(results)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVar(&searchQuery, "query", "", "Text query to search for markets (required)")
	searchCmd.MarkFlagRequired("query")

	searchCmd.Flags().IntVar(&searchLimit, "limit", 10, "Maximum number of results to return")
}
