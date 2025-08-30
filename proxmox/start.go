package proxmox

import (
	"context"
	"fmt"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

func StartVM(client *pveSDK.Client, vmID int) error {
	vmr := pveSDK.NewVmRef(pveSDK.GuestID(vmID))

	vmInfo, err := client.GetVmInfo(context.Background(), vmr)
	if err != nil {
		return err
	}

	status := vmInfo["status"].(string)
	if status == "running" {
		return fmt.Errorf("VM %d is already running", vmID)
	}

	_, err = client.StartVm(context.Background(), vmr)
	return err
}
