package utils

const (
	ResourceTypeQemu = "qemu"
	ResourceTypeLxc  = "lxc"
	ResourceTypeNode = "node"
)

const (
	VmStatusStopped   = "stopped"
	VmStatusRunning   = "running"
	VmStatusUnknown   = "unknown"
	VmStatusPrelaunch = "prelaunch"
)

const (
	ClusterStatusUnknown = "unknown"
	ClusterStatusOnline  = "online"
	ClusterStatusOffline = "offline"
)
