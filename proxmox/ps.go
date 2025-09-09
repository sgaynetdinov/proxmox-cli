package proxmox

import (
	"context"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
)

type VM struct {
	ID         int
	Name       string
	Status     string
	IsTemplate bool
	TypeVM     string
	Node       string
	Uptime     int64
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

		var isTemplate bool
		if t, ok := vmInfo["template"].(float64); ok {
			isTemplate = t == 1
		}

		var typeVM string
		if t, ok := vmInfo["type"].(string); ok {
			typeVM = t
		}

		var node string
		if n, ok := vmInfo["node"].(string); ok {
			node = n
		}

		var uptime int64
		if u, ok := vmInfo["uptime"].(float64); ok {
			uptime = int64(u)
		}

		vms = append(vms, VM{
			ID:         id,
			Name:       name,
			Status:     status,
			IsTemplate: isTemplate,
			TypeVM:     typeVM,
			Node:       node,
			Uptime:     uptime,
		})
	}
	return vms, nil
}
