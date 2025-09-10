package proxmox

import (
	"context"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

func VMList(client *pveSDK.Client) ([]VM, error) {
	vmList, err := client.GetVmList(context.Background())
	if err != nil {
		return nil, err
	}

	var vms []VM
	for _, vm := range vmList["data"].([]interface{}) {
		vmInfo := vm.(map[string]interface{})
		vm := VMFromMap(vmInfo)
		vms = append(vms, vm)
	}
	return vms, nil
}
