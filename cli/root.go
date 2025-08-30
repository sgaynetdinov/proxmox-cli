package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "proxmox-cli",
	Short: "A CLI tool for managing Proxmox VE",
	Long:  `A command line interface for interacting with Proxmox VE API`,
}

func init() {
	rootCmd.AddCommand(PsCmd)
	rootCmd.AddCommand(StartCmd)
	rootCmd.AddCommand(StopCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
