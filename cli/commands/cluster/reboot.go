package cluster

import (
	"fmt"

	"proxmox-cli/cli/utils"
	"proxmox-cli/proxmox"

	"github.com/spf13/cobra"
)

var RebootCmd = &cobra.Command{
	Use:   "reboot <NODE_NAME>",
	Short: "Reboot cluster node",
	Long:  `Reboot cluster node by their names.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nodeName := args[0]
		client := utils.GetClientFromContext(cmd)

		err := proxmox.ClusterRebootNode(client, nodeName)
		if err != nil {
			fmt.Printf("Error rebooting node %s: %v\n", nodeName, err)
		} else {
			fmt.Printf("Node %s reboot initiated successfully\n", nodeName)
		}
	},
}

func init() {
	ClusterCmd.AddCommand(RebootCmd)
}
