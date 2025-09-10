package proxmox

import (
	"context"

	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"proxmox-cli/proxmox/utils"
)

type ClusterNode struct {
	Name   string
	Status string
	Uptime int64
}

func ClusterFromMap(info map[string]interface{}) ClusterNode {
	var nodeName string
	if v, ok := info["node"].(string); ok {
		nodeName = v
	}

	var status string
	if v, ok := info["status"].(string); ok {
		status = v
	}

	var uptime int64
	if v, ok := info["uptime"].(float64); ok {
		uptime = int64(v)
	}

	return ClusterNode{
		Name:   nodeName,
		Status: status,
		Uptime: uptime,
	}
}


func ClusterNodeList(client *pveSDK.Client) ([]ClusterNode, error) {
	list, err := client.GetResourceList(context.Background(), utils.ResourceTypeNode)
	if err != nil {
		return nil, err
	}

	var nodes []ClusterNode
	for _, item := range list {
		info := item.(map[string]interface{})
		clusterNode := ClusterFromMap(info)
		nodes = append(nodes, clusterNode)
	}

	return nodes, nil
}
