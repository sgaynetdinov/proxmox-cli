package commands

import (
	"fmt"

	"proxmox-cli/cli/utils"
	"proxmox-cli/proxmox"

	"github.com/spf13/cobra"
)

var RebootCmd = &cobra.Command{
	Use:   "reboot <VM_ID> [VM_ID...]",
	Short: "Reboot one or more virtual machines",
	Long:  `Reboot one or more virtual machines by their IDs. Performs a graceful reboot via the guest OS (ACPI).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmIDs := utils.ParseVMIDs(args)
		client := utils.GetClientFromContext(cmd)

		utils.ExecuteVMOperations(vmIDs,
			func(vmID int) error {
				return proxmox.RebootVM(client, vmID)
			},
			func(vmID int) string {
				return fmt.Sprintf("VM %d reboot initiated successfully", vmID)
			},
		)
	},
}
