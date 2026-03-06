package vm

import (
	"context"
	"fmt"

	clicontext "proxmox-cli/internal/cli/context"
	"proxmox-cli/internal/cli/utils"
	"proxmox-cli/internal/proxmox"

	"github.com/spf13/cobra"
)

func rebootMany(ctx context.Context, client *proxmox.ProxmoxClient, vmIDs []int) {
	utils.ExecuteVMOperations(vmIDs,
		func(vmID int) error { return proxmox.RebootVM(ctx, client, vmID) },
		func(vmID int) string { return fmt.Sprintf("VM %d reboot initiated successfully", vmID) },
	)
}

func resetMany(ctx context.Context, client *proxmox.ProxmoxClient, vmIDs []int) {
	utils.ExecuteVMOperations(vmIDs,
		func(vmID int) error { return proxmox.ResetVM(ctx, client, vmID) },
		func(vmID int) string { return fmt.Sprintf("VM %d reset initiated successfully", vmID) },
	)
}

var RebootCmd = &cobra.Command{
	Use:     "reboot <VM_ID> [VM_ID...]",
	Short:   "Gracefully reboot one or more virtual machines (use -f for hard reset)",
	Long:    `Reboot one or more guests by their IDs. Performs a graceful reboot via the guest OS (ACPI). Use -f for a hard reset where supported (QEMU only; LXC may not support hard reset).`,
	Example: "  proxmox-cli vm reboot 101 102\n  proxmox-cli vm reboot -f 101",
	Aliases: []string{"restart"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmIDs := utils.ParseVMIDs(args)
		client := clicontext.GetClientFromContext(cmd)
		force, _ := cmd.Flags().GetBool("force")

		if force {
			resetMany(cmd.Context(), client, vmIDs)
		} else {
			rebootMany(cmd.Context(), client, vmIDs)
		}
	},
}

func init() {
	RebootCmd.Flags().BoolP("force", "f", false, "Force reboot by issuing a hard reset")
}
