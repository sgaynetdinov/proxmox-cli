package cli

import (
	"fmt"
	"os"
	"strconv"
)

func parseVMIDs(args []string) []int {
	var vmIDs []int
	for _, arg := range args {
		vmID, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid VM ID '%s'. VM ID must be a number.\n", arg)
			os.Exit(1)
		}
		vmIDs = append(vmIDs, vmID)
	}
	return vmIDs
}

func executeVMOperations(vmIDs []int, operation func(int) error, successMessage func(int) string) {
	hasErrors := false
	for _, vmID := range vmIDs {
		err := operation(vmID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error with VM %d: %v\n", vmID, err)
			hasErrors = true
		} else {
			fmt.Printf("%s\n", successMessage(vmID))
		}
	}

	if hasErrors {
		os.Exit(1)
	}
}
