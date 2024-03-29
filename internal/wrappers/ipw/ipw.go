package ipw

import (
	"os/exec"
)

const ip = "ip"

type IfType int64

const (
	Vcan IfType = iota
)

type IfPType int64

const (
	Vxcan IfPType = iota
	Veth
)

type IfState int64

const (
	DOWN IfState = iota
	UP
)

func (t IfType) String() string {
	return []string{"vcan"}[t]
}

func (t IfPType) String() string {
	return []string{"vxcan", "veth"}[t]
}

func (t IfState) String() string {
	return []string{"down", "up"}[t]
}

func CreateNetworkNamespace(name string) *exec.Cmd {
	return exec.Command(ip, "netns", "add", name)
}

func DeleteNetworkNamespace(name string) *exec.Cmd {
	return exec.Command(ip, "netns", "del", name)
}

func CreateInterface(ifName string, t IfType) *exec.Cmd {
	return exec.Command(ip, "link", "add", ifName, "type", t.String())
}

func CreateInterfacePair(ifName, peerName string, t IfPType) *exec.Cmd {
	return exec.Command(ip, "link", "add", ifName, "type", t.String(), "peer", "name", peerName)
}

func SetInterfaceState(ifName string, t IfState) *exec.Cmd {
	return exec.Command(ip, "link", "set", "dev", ifName, t.String())
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
