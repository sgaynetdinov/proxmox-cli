package cluster

import (
	"fmt"
	"os"

	"proxmox-cli/cli/utils"
	"proxmox-cli/proxmox"

	"github.com/spf13/cobra"
)

var shutdownCmd = &cobra.Command{
	Use:     "shutdown <NODE_NAME>",
	Short:   "Shutdown cluster node",
	Long:    `Shutdown cluster node by their names.`,
	Aliases: []string{"halt", "poweroff"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nodeName := args[0]
		client := utils.GetClientFromContext(cmd)

		err := proxmox.ClusterShutdownNode(cmd.Context(), client, nodeName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error shutdown node %s: %v\n", nodeName, err)
			return
		}
		fmt.Printf("Node %s shutdown initiated successfully\n", nodeName)
	},
}

func init() {
	ClusterCmd.AddCommand(shutdownCmd)
}
