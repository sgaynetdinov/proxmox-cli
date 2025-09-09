package vm

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"proxmox-cli/cli/utils"
	"proxmox-cli/proxmox"
	proxmox_utils "proxmox-cli/proxmox/utils"

	"github.com/spf13/cobra"
)

func sortVmByStatus(filteredVMs []proxmox.VM) func(i, j int) bool {
	return func(i, j int) bool {
		if filteredVMs[i].Status == proxmox_utils.VmStatusRunning && filteredVMs[j].Status != proxmox_utils.VmStatusRunning {
			return true
		}
		if filteredVMs[i].Status != proxmox_utils.VmStatusRunning && filteredVMs[j].Status == proxmox_utils.VmStatusRunning {
			return false
		}
		return filteredVMs[i].ID < filteredVMs[j].ID
	}
}

var rowFormat = "%-8s %-30s %-10s %-5s %-10s\n"

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
			if !showAll && vm.Status != proxmox_utils.VmStatusRunning {
				continue
			}

			if vm.IsTemplate {
				continue
			}

			filteredVMs = append(filteredVMs, vm)
		}

		if showAll {
			sort.SliceStable(filteredVMs, sortVmByStatus(filteredVMs))
		}

		fmt.Printf(rowFormat, "VM ID", "NAME", "STATUS", "TYPE", "NODE")

		for _, vm := range filteredVMs {
			name := vm.Name
			if name == "" {
				name = "<no name>"
			}
			fmt.Printf(rowFormat, strconv.Itoa(vm.ID), name, vm.Status, vm.TypeVM, vm.Node)
		}
	},
}

func init() {
	PsCmd.Flags().BoolP("all", "a", false, "Show all VMs (including stopped)")
}
