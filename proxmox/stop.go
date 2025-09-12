package proxmox

import (
	"context"
	"fmt"

	"proxmox-cli/proxmox/utils"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

func StopVM(client *pveSDK.Client, vmID int) error {
	vm, vmr, err := getVmInfo(client, vmID)
	if err != nil {
		return err
	}

	if vm.Status == utils.VmStatusStopped {
		return fmt.Errorf("VM %d is already stopped", vmID)
	}

	err = vmr.Stop(context.Background(), client)
	return err
}
