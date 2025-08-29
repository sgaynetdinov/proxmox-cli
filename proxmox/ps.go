package proxmox

import (
	"context"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

type VM struct {
	ID     int
	Name   string
	Status string
}

func VMList(client *pveSDK.Client) ([]VM, error) {
	vmList, err := client.GetVmList(context.Background())
	if err != nil {
		return nil, err
	}

	var vms []VM
	for _, vm := range vmList["data"].([]interface{}) {
		vmInfo := vm.(map[string]interface{})

		var id int
		if vmid, ok := vmInfo["vmid"].(float64); ok {
			id = int(vmid)
		}

		var name string
		if n, ok := vmInfo["name"].(string); ok {
			name = n
		}

		var status string
		if s, ok := vmInfo["status"].(string); ok {
			status = s
		}

		vms = append(vms, VM{
			ID:     id,
			Name:   name,
			Status: status,
		})
	}
	return vms, nil
}
