package cli

import (
	"fmt"

	"proxmox-cli/proxmox"

	"github.com/spf13/cobra"
)

var ShutdownCmd = &cobra.Command{
	Use:   "shutdown <VM_ID> [VM_ID...]",
	Short: "Gracefully shutdown one or more virtual machines",
	Long:  `Gracefully shutdown one or more virtual machines by their IDs. This sends an ACPI shutdown signal to the guest OS.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmIDs := parseVMIDs(args)
		client := GetClientFromContext(cmd)

		executeVMOperations(vmIDs,
			func(vmID int) error {
				return proxmox.ShutdownVM(client, vmID)
			},
			func(vmID int) string {
				return fmt.Sprintf("VM %d shutdown initiated successfully", vmID)
			},
		)
	},
}
