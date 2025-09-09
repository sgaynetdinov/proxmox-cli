package vm

import (
	"fmt"

	"proxmox-cli/cli/utils"
	"proxmox-cli/proxmox"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
)

func rebootMany(client *pveSDK.Client, vmIDs []int) {
	utils.ExecuteVMOperations(vmIDs,
		func(vmID int) error { return proxmox.RebootVM(client, vmID) },
		func(vmID int) string { return fmt.Sprintf("VM %d reboot initiated successfully", vmID) },
	)
}

var RebootCmd = &cobra.Command{
	Use:   "reboot <VM_ID> [VM_ID...]",
	Short: "Reboot one or more virtual machines",
	Long:  `Reboot one or more virtual machines by their IDs. Performs a graceful reboot via the guest OS (ACPI).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmIDs := utils.ParseVMIDs(args)
		client := utils.GetClientFromContext(cmd)
		force, _ := cmd.Flags().GetBool("force")

		if force {
			resetMany(client, vmIDs)
		} else {
			rebootMany(client, vmIDs)
		}
	},
}

func init() {
	RebootCmd.Flags().BoolP("force", "f", false, "Force reboot by issuing a hard reset")
}
