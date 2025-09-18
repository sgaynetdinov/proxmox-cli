package cluster

import (
	"fmt"
	"os"

	"proxmox-cli/internal/cli/utils"
	"proxmox-cli/internal/proxmox"

	"github.com/spf13/cobra"
)

var rebootCmd = &cobra.Command{
	Use:     "reboot <NODE_NAME>",
	Short:   "Reboot cluster node",
	Long:    `Reboot cluster node by their names.`,
	Aliases: []string{"restart"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nodeName := args[0]
		client := utils.GetClientFromContext(cmd)

		err := proxmox.ClusterRebootNode(cmd.Context(), client, nodeName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error rebooting node %s: %v\n", nodeName, err)
			return
		}
		fmt.Printf("Node %s reboot initiated successfully\n", nodeName)
	},
}

func init() {
	ClusterCmd.AddCommand(rebootCmd)
}
