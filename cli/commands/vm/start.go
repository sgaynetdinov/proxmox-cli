package vm

import (
	"fmt"

	"proxmox-cli/cli/utils"
	"proxmox-cli/proxmox"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start <VM_ID> [VM_ID...]",
	Short: "Start one or more virtual machines",
	Long:  `Start one or more virtual machines by their IDs`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmIDs := utils.ParseVMIDs(args)
		client := utils.GetClientFromContext(cmd)

		utils.ExecuteVMOperations(vmIDs,
			func(vmID int) error {
				return proxmox.StartVM(client, vmID)
			},
			func(vmID int) string {
				return fmt.Sprintf("VM %d started successfully", vmID)
			},
		)
	},
}
