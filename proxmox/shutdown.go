package proxmox

import (
	"context"
	"fmt"

	"proxmox-cli/proxmox/utils"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

func ShutdownVM(client *pveSDK.Client, vmID int) error {
	vmr := pveSDK.NewVmRef(pveSDK.GuestID(vmID))

	vmInfo, err := client.GetVmInfo(context.Background(), vmr)
	if err != nil {
		return err
	}

	vm := VMFromMap(vmInfo)

	if vm.Status == utils.VmStatusStopped {
		return fmt.Errorf("VM %d is already stopped", vmID)
	}

	_, err = client.ShutdownVm(context.Background(), vmr)
	return err
}
