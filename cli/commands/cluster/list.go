package cluster

import (
	"fmt"
	"os"
	"sort"

	"proxmox-cli/cli/utils"
	"proxmox-cli/proxmox"
	proxmox_utils "proxmox-cli/proxmox/utils"

	"github.com/spf13/cobra"
)

var rowFormat = "%-20s %-10s %-10s\n"

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List cluster nodes",
	Long:  `List all nodes in the Proxmox cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		client := utils.GetClientFromContext(cmd)

		nodes, err := proxmox.ClusterNodeList(client)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting cluster nodes: %v\n", err)
			os.Exit(1)
		}

		sort.Slice(nodes, func(i, j int) bool {
			// Put "online" status at the top
			if nodes[i].Status == proxmox_utils.ClusterStatusOnline && nodes[j].Status != proxmox_utils.ClusterStatusOnline {
				return true
			}
			if nodes[i].Status != proxmox_utils.ClusterStatusOnline && nodes[j].Status == proxmox_utils.ClusterStatusOnline {
				return false
			}
			// If status is the same, sort by name
			return nodes[i].Name < nodes[j].Name
		})

		fmt.Printf(rowFormat, "NAME", "STATUS", "UPTIME")
		for _, n := range nodes {
			name := n.Name
			if name == "" {
				name = "<unknown>"
			}

			uptime := proxmox_utils.FormatSecondsHMS(n.Uptime)
			if n.Status != proxmox_utils.ClusterStatusOnline {
				uptime = "-"
			}

			fmt.Printf(rowFormat, name, n.Status, uptime)
		}
	},
}

func init() {
	ClusterCmd.AddCommand(ListCmd)
}
