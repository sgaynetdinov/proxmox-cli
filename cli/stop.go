package cli

import (
	"fmt"
	"os"
	"strconv"

	"proxmox-cli/proxmox"

	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop <VM_ID> [VM_ID...]",
	Short: "Stop one or more virtual machines",
	Long:  `Stop one or more virtual machines by their IDs`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Parse all VM IDs first
		var vmIDs []int
		for _, arg := range args {
			vmID, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Invalid VM ID '%s'. VM ID must be a number.\n", arg)
				os.Exit(1)
			}
			vmIDs = append(vmIDs, vmID)
		}

		client := GetClientFromContext(cmd)

		hasErrors := false
		for _, vmID := range vmIDs {
			err := proxmox.StopVM(client, vmID)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error stopping VM %d: %v\n", vmID, err)
				hasErrors = true
			} else {
				fmt.Printf("VM %d stopped successfully\n", vmID)
			}
		}

		if hasErrors {
			os.Exit(1)
		}
	},
}
