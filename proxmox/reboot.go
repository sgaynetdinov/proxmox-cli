package proxmox

import (
	"context"
	"fmt"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"

	"proxmox-cli/proxmox/utils"
)

func RebootVM(client *pveSDK.Client, vmID int) error {
	vm, vmr, err := getVmInfo(client, vmID)
	if err != nil {
		return err
	}

	if vm.Status != utils.VmStatusRunning {
		return fmt.Errorf("VM %d is not running", vmID)
	}

	_, err = client.RebootVm(context.Background(), vmr)
	return err
}
