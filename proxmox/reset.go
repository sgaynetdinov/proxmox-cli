package proxmox

import (
	"context"
	"fmt"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"

	"proxmox-cli/proxmox/utils"
)

func ResetVM(client *pveSDK.Client, vmID int) error {
	vmr := pveSDK.NewVmRef(pveSDK.GuestID(vmID))

	vmInfo, err := client.GetVmInfo(context.Background(), vmr)
	if err != nil {
		return err
	}

	if vmr.GetVmType() == utils.VmTypeLxc {
		return fmt.Errorf("VM %d reset operation is not supported for LXC containers", vmID)
	}

	status := vmInfo["status"].(string)
	if status != utils.VmStatusRunning {
		return fmt.Errorf("VM %d is not running", vmID)
	}

	_, err = client.ResetVm(context.Background(), vmr)
	return err
}
