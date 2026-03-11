package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	jsonOutput bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "polymarket",
	Short: "A CLI tool to interact with Polymarket APIs",
	Long: `Polymarket CLI allows interaction with Polymarket's Gamma and CLOB APIs
from the command line, optimized for JSON output and LLM Agents.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		printError(err)
		os.Exit(1)
	}
}

func init() {
	// Global flag for json output. We make it optional but recommended for agents.
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", true, "Output results in JSON format")
}

// printSuccess handles generic success output, heavily skewed towards JSON
func printSuccess(data interface{}) {
	if jsonOutput {
		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			printError(fmt.Errorf("failed to marshal JSON output: %w", err))
			os.Exit(1)
		}
		fmt.Println(string(b))
	} else {
		// Fallback simple print if they disable --json=false
		fmt.Printf("%+v\n", data)
	}
}

// printError handles uniform error reporting in JSON format if requested
func printError(err error) {
	if jsonOutput {
		errResp := map[string]string{"error": err.Error()}
		b, _ := json.MarshalIndent(errResp, "", "  ")
		fmt.Fprintln(os.Stderr, string(b))
	} else {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}
