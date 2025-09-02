package proxmox

import (
	"context"
	"fmt"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

func ShutdownVM(client *pveSDK.Client, vmID int) error {
	vmr := pveSDK.NewVmRef(pveSDK.GuestID(vmID))

	vmInfo, err := client.GetVmInfo(context.Background(), vmr)
	if err != nil {
		return err
	}

	status := vmInfo["status"].(string)
	if status == "stopped" {
		return fmt.Errorf("VM %d is already stopped", vmID)
	}

	_, err = client.ShutdownVm(context.Background(), vmr)
	return err
}
