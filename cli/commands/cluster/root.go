package cluster

import (
	"github.com/spf13/cobra"
)

var ClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Cluster related commands",
	Long:  `Manage Proxmox cluster resources and operations`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
