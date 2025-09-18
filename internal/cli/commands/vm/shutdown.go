package vm

import (
	"context"
	"fmt"

	"proxmox-cli/internal/cli/utils"
	"proxmox-cli/internal/proxmox"

	"github.com/spf13/cobra"
)

func shutdownMany(ctx context.Context, client *proxmox.ProxmoxClient, vmIDs []int) {
	utils.ExecuteVMOperations(vmIDs,
		func(vmID int) error { return proxmox.ShutdownVM(ctx, client, vmID) },
		func(vmID int) string { return fmt.Sprintf("VM %d shutdown initiated successfully", vmID) },
	)
}

var ShutdownCmd = &cobra.Command{
	Use:     "shutdown <VM_ID> [VM_ID...]",
	Short:   "Gracefully shutdown one or more virtual machines",
	Long:    `Gracefully shutdown one or more virtual machines by their IDs. This sends an ACPI shutdown signal to the guest OS.`,
	Aliases: []string{"halt", "poweroff"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmIDs := utils.ParseVMIDs(args)
		client := utils.GetClientFromContext(cmd)
		force, _ := cmd.Flags().GetBool("force")

		if force {
			stopMany(cmd.Context(), client, vmIDs)
		} else {
			shutdownMany(cmd.Context(), client, vmIDs)
		}
	},
}

func init() {
	ShutdownCmd.Flags().BoolP("force", "f", false, "Force immediate stop instead of graceful shutdown")
}
