package utils

const (
	VmTypeQemu = "qemu"
	VmTypeLxc  = "lxc"
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
