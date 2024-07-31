package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ethsh",
	Short: "Ethers shell is a cli tool for interacting with smart contracts",
	Long:  "Ethers shell is a cli tool for interacting with smart contracts on evm compatible networks - read, write, call etc.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Ethers Shell '%s'\n", err)
		os.Exit(1)
	}
}
