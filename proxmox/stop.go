package proxmox

import (
	"context"
	"fmt"

	"proxmox-cli/proxmox/utils"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

func StopVM(client *pveSDK.Client, vmID int) error {
	vmr := pveSDK.NewVmRef(pveSDK.GuestID(vmID))

	vmInfo, err := client.GetVmInfo(context.Background(), vmr)
	if err != nil {
		return err
	}

	vm := VMFromMap(vmInfo)

	if vm.Status == utils.VmStatusStopped {
		return fmt.Errorf("VM %d is already stopped", vmID)
	}

	err = vmr.Stop(context.Background(), client)
	return err
}
