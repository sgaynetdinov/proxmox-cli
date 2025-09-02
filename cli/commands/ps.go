package commands

import (
	"fmt"
	"os"

	"proxmox-cli/cli/utils"
	"proxmox-cli/proxmox"

	"github.com/spf13/cobra"
)

var PsCmd = &cobra.Command{
	Use:   "ps",
	Short: "List virtual machines",
	Long:  `List all virtual machines from Proxmox VE. By default shows only running VMs.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := utils.GetClientFromContext(cmd)

		vms, err := proxmox.VMList(client)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting VM list: %v\n", err)
			os.Exit(1)
		}

		showAll, _ := cmd.Flags().GetBool("all")

		var filteredVMs []proxmox.VM
		for _, vm := range vms {
			if !showAll && vm.Status != "running" {
				continue
			}

			if vm.IsTemplate {
				continue
			}

			filteredVMs = append(filteredVMs, vm)
		}

		fmt.Printf("%-8s %-30s %-10s\n", "VM ID", "NAME", "STATUS")

		for _, vm := range filteredVMs {
			name := vm.Name
			if name == "" {
				name = "<no name>"
			}
			fmt.Printf("%-8d %-30s %-10s\n", vm.ID, name, vm.Status)
		}
	},
}

func init() {
	PsCmd.Flags().BoolP("all", "a", false, "Show all VMs (including stopped)")
}
