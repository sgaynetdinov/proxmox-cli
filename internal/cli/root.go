package cli

import (
	"context"
	"fmt"
	"os"

	commands_cluster "proxmox-cli/internal/cli/commands/cluster"
	commands_vm "proxmox-cli/internal/cli/commands/vm"
	clicontext "proxmox-cli/internal/cli/context"
	"proxmox-cli/internal/cli/utils"
	"proxmox-cli/internal/proxmox"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "proxmox-cli",
	Short: "A CLI tool for managing Proxmox VE",
	Long:  `A command line interface for interacting with Proxmox VE API`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		apiURL, username, password, err := utils.GetCredentialsFromEnv()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting credentials: %v\n", err)
			os.Exit(1)
		}

		client, err := proxmox.NewClient(cmd.Context(), apiURL, username, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating Proxmox client: %v\n", err)
			os.Exit(1)
		}
		ctx := context.WithValue(cmd.Context(), clicontext.ClientKey, client)
		cmd.SetContext(ctx)
	},
}

func init() {
	// Cluster commands
	rootCmd.AddCommand(commands_cluster.ClusterCmd)

	// VM commands
	rootCmd.AddCommand(commands_vm.PsCmd)
	rootCmd.AddCommand(commands_vm.StartCmd)
	rootCmd.AddCommand(commands_vm.StopCmd)
	rootCmd.AddCommand(commands_vm.ShutdownCmd)
	rootCmd.AddCommand(commands_vm.ResetCmd)
	rootCmd.AddCommand(commands_vm.RebootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
