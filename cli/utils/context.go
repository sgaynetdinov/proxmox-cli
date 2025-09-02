package utils

import (
	"fmt"
	"os"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
)

type contextKey string

const ClientKey contextKey = "proxmox-client"

func GetClientFromContext(cmd *cobra.Command) *pveSDK.Client {
	client, ok := cmd.Context().Value(ClientKey).(*pveSDK.Client)
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Proxmox client not found in context\n")
		os.Exit(1)
	}
	return client
}
