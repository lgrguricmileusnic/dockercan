package netns

import (
	"os/exec"
)

const command = "ip"

func CreateNetworkNamespace(name string) error {
	cmd := exec.Command(command, "netns", "add", name)
	return cmd.Run()
}
