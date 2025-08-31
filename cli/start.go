package cli

import (
	"fmt"
	"os"
	"strconv"

	"proxmox-cli/proxmox"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start <VM_ID> [VM_ID...]",
	Short: "Start one or more virtual machines",
	Long:  `Start one or more virtual machines by their IDs`,
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

		client := proxmox.Login(cmd.Context())

		hasErrors := false
		for _, vmID := range vmIDs {
			err := proxmox.StartVM(client, vmID)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error starting VM %d: %v\n", vmID, err)
				hasErrors = true
			} else {
				fmt.Printf("VM %d started successfully\n", vmID)
			}
		}

		if hasErrors {
			os.Exit(1)
		}
	},
}
