package proxmox

import (
	"context"
	"fmt"

	"proxmox-cli/proxmox/utils"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

func StartVM(client *pveSDK.Client, vmID int) error {
	vm, vmr, err := getVmInfo(client, vmID)
	if err != nil {
		return err
	}

	if vm.Status == utils.VmStatusRunning {
		return fmt.Errorf("VM %d is already running", vmID)
	}

	_, err = client.StartVm(context.Background(), vmr)
	return err
}
