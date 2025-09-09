package vm

import (
	"fmt"

	"proxmox-cli/cli/utils"
	"proxmox-cli/proxmox"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
)

func stopMany(client *pveSDK.Client, vmIDs []int) {
	utils.ExecuteVMOperations(vmIDs,
		func(vmID int) error { return proxmox.StopVM(client, vmID) },
		func(vmID int) string { return fmt.Sprintf("VM %d stopped successfully", vmID) },
	)
}

var StopCmd = &cobra.Command{
	Use:   "stop <VM_ID> [VM_ID...]",
	Short: "Stop one or more virtual machines",
	Long:  `Stop one or more virtual machines by their IDs`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmIDs := utils.ParseVMIDs(args)
		client := utils.GetClientFromContext(cmd)

		stopMany(client, vmIDs)
	},
}
