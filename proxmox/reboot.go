package proxmox

import (
	"context"
	"fmt"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"

	"proxmox-cli/proxmox/utils"
)

func RebootVM(client *pveSDK.Client, vmID int) error {
	vmr := pveSDK.NewVmRef(pveSDK.GuestID(vmID))

	vmInfo, err := client.GetVmInfo(context.Background(), vmr)
	if err != nil {
		return err
	}

	vm := VMFromMap(vmInfo)

	if vm.Status != utils.VmStatusRunning {
		return fmt.Errorf("VM %d is not running", vmID)
	}

	_, err = client.RebootVm(context.Background(), vmr)
	return err
}
