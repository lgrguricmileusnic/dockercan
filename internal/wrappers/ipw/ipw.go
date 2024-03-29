package ipw

import (
	"os/exec"
)

const ip = "ip"

type IfType int64

const (
	Vxcan IfType = iota
	Vcan
	Veth
)

func (t IfType) String() string {
	return []string{"vxcan", "vcan", "veth"}[t]
}

func CreateNetworkNamespace(name string) *exec.Cmd {
	return exec.Command(ip, "netns", "add", name)
}

func DeleteNetworkNamespace(name string) *exec.Cmd {
	return exec.Command(ip, "netns", "del", name)
}

func CreateInterfacePair(ifName, peerName string, t IfType) *exec.Cmd {
	return exec.Command(ip, "link", "add", ifName, "type", t.String(), "peer", "name", peerName)
}

func DeleteInterface(ifName string) *exec.Cmd {
	return exec.Command(ip, "link", "del", ifName)
}

func MoveInterfaceToNamespace(ifName, nsName string) *exec.Cmd {
	return exec.Command(ip, "link", "set", ifName, "netns", nsName)
}

func ExecCommandInNamespace(nsName string, c exec.Cmd) *exec.Cmd {
	cs := make([]string, len(c.Args)+3)
	cs[0] = "netns"
	cs[1] = "exec"
	cs[2] = nsName
	copy(cs[3:], c.Args)

	cmd := exec.Command(ip, cs...)
	return cmd
}
