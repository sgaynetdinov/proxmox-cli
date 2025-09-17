package proxmox

import (
	"context"
	"fmt"

	"proxmox-cli/proxmox/utils"

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

func VMFromMap(vmInfo map[string]interface{}) VM {
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

	return VM{
		ID:         id,
		Name:       name,
		Status:     status,
		IsTemplate: isTemplate,
		TypeVM:     typeVM,
		Node:       node,
		Uptime:     uptime,
	}
}

func getVmInfo(ctx context.Context, client *ProxmoxClient, vmID int) (VM, *pveSDK.VmRef, error) {
	vmr := pveSDK.NewVmRef(pveSDK.GuestID(vmID))

	vmInfo, err := client.GetVmInfo(ctx, vmr)
	if err != nil {
		return VM{}, vmr, err
	}

	vm := VMFromMap(vmInfo)
	return vm, vmr, nil
}

func VMList(ctx context.Context, client *ProxmoxClient) ([]VM, error) {
	vmList, err := client.GetVmList(ctx)
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

func StartVM(ctx context.Context, client *ProxmoxClient, vmID int) error {
	vm, vmr, err := getVmInfo(ctx, client, vmID)
	if err != nil {
		return err
	}

	if vm.Status == utils.VmStatusRunning {
		return fmt.Errorf("VM %d is already running", vmID)
	}

	_, err = client.StartVm(ctx, vmr)
	return err
}

func StopVM(ctx context.Context, client *ProxmoxClient, vmID int) error {
	vm, vmr, err := getVmInfo(ctx, client, vmID)
	if err != nil {
		return err
	}

	if vm.Status == utils.VmStatusStopped {
		return fmt.Errorf("VM %d is already stopped", vmID)
	}

	err = vmr.Stop(ctx, client)
	return err
}

func ShutdownVM(ctx context.Context, client *ProxmoxClient, vmID int) error {
	vm, vmr, err := getVmInfo(ctx, client, vmID)
	if err != nil {
		return err
	}

	if vm.Status == utils.VmStatusStopped {
		return fmt.Errorf("VM %d is already stopped", vmID)
	}

	_, err = client.ShutdownVm(ctx, vmr)
	return err
}

func ResetVM(ctx context.Context, client *ProxmoxClient, vmID int) error {
	vm, vmr, err := getVmInfo(ctx, client, vmID)
	if err != nil {
		return err
	}

	if vm.TypeVM == utils.ResourceTypeLxc {
		return fmt.Errorf("VM %d reset operation is not supported for LXC containers", vmID)
	}

	if vm.Status != utils.VmStatusRunning {
		return fmt.Errorf("VM %d is not running", vmID)
	}

	_, err = client.ResetVm(ctx, vmr)
	return err
}

func RebootVM(ctx context.Context, client *ProxmoxClient, vmID int) error {
	vm, vmr, err := getVmInfo(ctx, client, vmID)
	if err != nil {
		return err
	}

	if vm.Status != utils.VmStatusRunning {
		return fmt.Errorf("VM %d is not running", vmID)
	}

	_, err = client.RebootVm(ctx, vmr)
	return err
}
