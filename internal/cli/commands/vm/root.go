package vm

import "github.com/spf13/cobra"

var VmCmd = &cobra.Command{
	Use:   "vm",
	Short: "VM related commands (QEMU/LXC)",
	Long:  `VM related commands and operations for QEMU virtual machines and LXC containers`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	VmCmd.AddCommand(PsCmd)
	VmCmd.AddCommand(StartCmd)
	VmCmd.AddCommand(ShutdownCmd)
	VmCmd.AddCommand(RebootCmd)
}
