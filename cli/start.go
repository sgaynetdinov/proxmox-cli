package cli

import (
	"fmt"
	"os"
	"strconv"

	"proxmox-cli/proxmox"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start <VM_ID>",
	Short: "Start a virtual machine",
	Long:  `Start a virtual machine by its ID`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vmID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid VM ID '%s'. VM ID must be a number.\n", args[0])
			os.Exit(1)
		}

		client := proxmox.Login(cmd.Context())

		err = proxmox.StartVM(client, vmID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error starting VM %d: %v\n", vmID, err)
			os.Exit(1)
		}

		fmt.Printf("VM %d started successfully\n", vmID)
	},
}
