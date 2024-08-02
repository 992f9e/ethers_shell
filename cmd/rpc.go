package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addFlag string
var listFlag bool

var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "Manage RPC endpoints",
	Long:  `Add or list RPC endpoints for your CLI application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if addFlag != "" {
			addRPC(addFlag)
		} else if listFlag {
			listRPCs()
		} else {
			fmt.Println("Use --add to add an RPC or --list to list all RPCs")
		}
	},
}

func init() {
	rootCmd.AddCommand(rpcCmd)

	rpcCmd.Flags().StringVarP(&addFlag, "add", "a", "", "Add a new RPC endpoint")
	rpcCmd.Flags().BoolVarP(&listFlag, "list", "l", false, "List all RPC endpoints")
}

func addRPC(rpc string) {
	rpcs := viper.GetStringSlice("rpcs")
	rpcs = append(rpcs, rpc)
	viper.Set("rpcs", rpcs)
	if err := viper.WriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.SafeWriteConfig(); err != nil {
				fmt.Println("Error creating config file:", err)
			}
		} else {
			fmt.Println("Error saving config:", err)
		}
	}
	fmt.Printf("RPC endpoint '%s' added.\n", rpc)
}

func listRPCs() {
	rpcs := viper.GetStringSlice("rpcs")
	if len(rpcs) == 0 {
		fmt.Println("No RPC endpoints found.")
		return
	}
	fmt.Println("RPC endpoints:")
	for _, rpc := range rpcs {
		fmt.Println(rpc)
	}
}
