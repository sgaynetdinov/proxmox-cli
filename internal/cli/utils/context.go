package utils

import (
	"fmt"
	"os"

	"proxmox-cli/internal/proxmox"

	"github.com/spf13/cobra"
)

type contextKey string

const ClientKey contextKey = "proxmox-client"

func GetClientFromContext(cmd *cobra.Command) *proxmox.ProxmoxClient {
	client, ok := cmd.Context().Value(ClientKey).(*proxmox.ProxmoxClient)
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Proxmox client not found in context\n")
		os.Exit(1)
	}
	return client
}
