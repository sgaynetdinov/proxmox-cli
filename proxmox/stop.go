package proxmox

import (
	"context"
	"fmt"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

func StopVM(client *pveSDK.Client, vmID int) error {
	vmr := pveSDK.NewVmRef(pveSDK.GuestID(vmID))

	vmInfo, err := client.GetVmInfo(context.Background(), vmr)
	if err != nil {
		return err
	}

	status := vmInfo["status"].(string)
	if status == "stopped" {
		return fmt.Errorf("VM %d is already stopped", vmID)
	}

	err = vmr.Stop(context.Background(), client)
	return err
}
