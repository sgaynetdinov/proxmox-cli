package commands

import (
	"fmt"

	"proxmox-cli/cli/utils"
	"proxmox-cli/proxmox"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
)

func shutdownMany(client *pveSDK.Client, vmIDs []int) {
	utils.ExecuteVMOperations(vmIDs,
		func(vmID int) error { return proxmox.ShutdownVM(client, vmID) },
		func(vmID int) string { return fmt.Sprintf("VM %d shutdown initiated successfully", vmID) },
	)
}

var ShutdownCmd = &cobra.Command{
	Use:   "shutdown <VM_ID> [VM_ID...]",
	Short: "Gracefully shutdown one or more virtual machines",
	Long:  `Gracefully shutdown one or more virtual machines by their IDs. This sends an ACPI shutdown signal to the guest OS.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmIDs := utils.ParseVMIDs(args)
		client := utils.GetClientFromContext(cmd)
		force, _ := cmd.Flags().GetBool("force")

		if force {
			stopMany(client, vmIDs)
		} else {
			shutdownMany(client, vmIDs)
		}
	},
}

func init() {
	ShutdownCmd.Flags().BoolP("force", "f", false, "Force immediate stop instead of graceful shutdown")
}
