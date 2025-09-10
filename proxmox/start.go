package proxmox

import (
	"context"
	"fmt"

	"proxmox-cli/proxmox/utils"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

func StartVM(client *pveSDK.Client, vmID int) error {
	vmr := pveSDK.NewVmRef(pveSDK.GuestID(vmID))

	vmInfo, err := client.GetVmInfo(context.Background(), vmr)
	if err != nil {
		return err
	}

	vm := VMFromMap(vmInfo)

	if vm.Status == utils.VmStatusRunning {
		return fmt.Errorf("VM %d is already running", vmID)
	}

	_, err = client.StartVm(context.Background(), vmr)
	return err
}
