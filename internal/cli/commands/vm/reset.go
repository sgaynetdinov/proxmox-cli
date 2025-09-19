package vm

import (
	"context"
	"fmt"

	clicontext "proxmox-cli/internal/cli/context"
	"proxmox-cli/internal/cli/utils"
	"proxmox-cli/internal/proxmox"

	"github.com/spf13/cobra"
)

func resetMany(ctx context.Context, client *proxmox.ProxmoxClient, vmIDs []int) {
	utils.ExecuteVMOperations(vmIDs,
		func(vmID int) error { return proxmox.ResetVM(ctx, client, vmID) },
		func(vmID int) string { return fmt.Sprintf("VM %d reset initiated successfully", vmID) },
	)
}

var ResetCmd = &cobra.Command{
	Use:   "reset <VM_ID> [VM_ID...]",
	Short: "Reset one or more virtual machines",
	Long:  `Reset one or more virtual machines by their IDs. This performs a hard reset, equivalent to pressing the reset button.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmIDs := utils.ParseVMIDs(args)
		client := clicontext.GetClientFromContext(cmd)

		resetMany(cmd.Context(), client, vmIDs)
	},
}
