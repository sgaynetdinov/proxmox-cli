package cli

import (
	"context"
	"fmt"
	"os"

	"proxmox-cli/cli/commands"
	"proxmox-cli/cli/utils"
	"proxmox-cli/proxmox"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "proxmox-cli",
	Short: "A CLI tool for managing Proxmox VE",
	Long:  `A command line interface for interacting with Proxmox VE API`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client := proxmox.Login(cmd.Context())
		ctx := context.WithValue(cmd.Context(), utils.ClientKey, client)
		cmd.SetContext(ctx)
	},
}

func init() {
	rootCmd.AddCommand(commands.PsCmd)
	rootCmd.AddCommand(commands.StartCmd)
	rootCmd.AddCommand(commands.StopCmd)
	rootCmd.AddCommand(commands.ShutdownCmd)
	rootCmd.AddCommand(commands.ResetCmd)
	rootCmd.AddCommand(commands.RebootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
