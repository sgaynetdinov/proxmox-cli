package cli

import (
	"fmt"
	"os"
	"strconv"

	"proxmox-cli/proxmox"

	"github.com/spf13/cobra"
)

var ShutdownCmd = &cobra.Command{
	Use:   "shutdown <VM_ID> [VM_ID...]",
	Short: "Gracefully shutdown one or more virtual machines",
	Long:  `Gracefully shutdown one or more virtual machines by their IDs. This sends an ACPI shutdown signal to the guest OS.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
			err := proxmox.ShutdownVM(client, vmID)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error shutting down VM %d: %v\n", vmID, err)
				hasErrors = true
			} else {
				fmt.Printf("VM %d shutdown initiated successfully\n", vmID)
			}
		}

		if hasErrors {
			os.Exit(1)
		}
	},
}
