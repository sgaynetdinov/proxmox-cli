package utils

import (
	"fmt"
	"os"
	"strconv"
)

func ParseVMIDs(args []string) []int {
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

func ExecuteVMOperations(vmIDs []int, operation func(int) error, successMessage func(int) string) {
	for _, vmID := range vmIDs {
		err := operation(vmID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error with VM %d: %v\n", vmID, err)
		} else {
			fmt.Printf("%s\n", successMessage(vmID))
		}
	}
}

func GetCredentialsFromEnv() (string, string, string, error) {
	apiURL := os.Getenv("PM_API_URL")
	username := os.Getenv("PM_USER")
	password := os.Getenv("PM_PASS")

	if apiURL == "" {
		return "", "", "", fmt.Errorf("PM_API_URL environment variable must be set")
	}

	if username == "" || password == "" {
		return "", "", "", fmt.Errorf("PM_USER and PM_PASS environment variables must be set")
	}
	return apiURL, username, password, nil
}
