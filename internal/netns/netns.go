package netns

import (
	"os/exec"
)

const command = "ip"

type IfType int64

const (
	Vxcan IfType = iota
	Veth
)

func (t IfType) String() string {
	return []string{"vxcan", "veth"}[t]
}

func CreateNetworkNamespace(name string) error {
	cmd := exec.Command(command, "netns", "add", name)
	return cmd.Run()
}

func DeleteNetworkNamespace(name string) error {
	cmd := exec.Command(command, "netns", "del", name)
	return cmd.Run()
}

func CreateInterfacePair(peer1 string, peer2 string, ifType IfType) error {
	// ip link add vxcan1 type vxcan peer name vxcan2
	cmd := exec.Command(command, "link", "add", peer1, "type", ifType.String(), "peer", "name", peer2)
	return cmd.Run()
}
