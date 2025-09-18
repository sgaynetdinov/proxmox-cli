package proxmox

import (
	"context"
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

func ClusterNodeList(ctx context.Context, client *ProxmoxClient) ([]ClusterNode, error) {
	nodeList, err := client.GetNodeList(ctx)
	if err != nil {
		return nil, err
	}

	var nodes []ClusterNode
	for _, item := range nodeList["data"].([]interface{}) {
		info := item.(map[string]interface{})
		clusterNode := ClusterFromMap(info)
		nodes = append(nodes, clusterNode)
	}

	return nodes, nil
}

func ClusterRebootNode(ctx context.Context, client *ProxmoxClient, nodeName string) error {
	_, err := client.RebootNode(ctx, nodeName)
	return err
}

func ClusterShutdownNode(ctx context.Context, client *ProxmoxClient, nodeName string) error {
	_, err := client.ShutdownNode(ctx, nodeName)
	return err
}
