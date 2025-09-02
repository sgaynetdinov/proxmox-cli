package cli

import (
	"context"
	"fmt"
	"os"

	"proxmox-cli/proxmox"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
)

type contextKey string

const ClientKey contextKey = "proxmox-client"

var rootCmd = &cobra.Command{
	Use:   "proxmox-cli",
	Short: "A CLI tool for managing Proxmox VE",
	Long:  `A command line interface for interacting with Proxmox VE API`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client := proxmox.Login(cmd.Context())
		ctx := context.WithValue(cmd.Context(), ClientKey, client)
		cmd.SetContext(ctx)
	},
}

func init() {
	rootCmd.AddCommand(PsCmd)
	rootCmd.AddCommand(StartCmd)
	rootCmd.AddCommand(StopCmd)
	rootCmd.AddCommand(ShutdownCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func GetClientFromContext(cmd *cobra.Command) *pveSDK.Client {
	client, ok := cmd.Context().Value(ClientKey).(*pveSDK.Client)
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Proxmox client not found in context\n")
		os.Exit(1)
	}
	return client
}
