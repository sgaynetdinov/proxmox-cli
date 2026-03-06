package vm

import (
	"context"
	"fmt"

	clicontext "proxmox-cli/internal/cli/context"
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

func stopMany(ctx context.Context, client *proxmox.ProxmoxClient, vmIDs []int) {
	utils.ExecuteVMOperations(vmIDs,
		func(vmID int) error { return proxmox.StopVM(ctx, client, vmID) },
		func(vmID int) string { return fmt.Sprintf("VM %d stopped successfully", vmID) },
	)
}

var ShutdownCmd = &cobra.Command{
	Use:     "shutdown <VM_ID> [VM_ID...]",
	Short:   "Gracefully shutdown one or more virtual machines (use -f for force stop)",
	Long:    `Gracefully shutdown one or more guests by their IDs. This sends an ACPI shutdown signal to the guest OS. Use -f to force an immediate stop instead.`,
	Example: "  proxmox-cli vm shutdown 101 102\n  proxmox-cli vm shutdown -f 101 102",
	Aliases: []string{"halt", "poweroff"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmIDs := utils.ParseVMIDs(args)
		client := clicontext.GetClientFromContext(cmd)
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
