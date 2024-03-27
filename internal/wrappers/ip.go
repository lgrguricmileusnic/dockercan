package wrappers

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

func CreateInterfacePair(ifName string, peerName string, ifType IfType) error {
	// ip link add vxcan1 type vxcan peer name vxcan2
	cmd := exec.Command(command, "link", "add", ifName, "type", ifType.String(), "peer", "name", peerName)
	return cmd.Run()
}

func MoveInterfaceToNamespace(ifName string, nsName string) error {
	cmd := exec.Command(command, "link", "set", ifName, "netns", nsName)
	return cmd.Run()
}

func ExecCommandInNamespace(nsName string, cmd exec.Cmd) error {
	// TODO:
	// cmd := exec.Command()
	return nil
}
