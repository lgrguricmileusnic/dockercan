package cangww

import (
	"os/exec"
)

const cangw = "cangw"

func AddRuleCmd(srcIf string, dstIf string, echo bool, canfd bool, routeIncoming bool) *exec.Cmd {
	args := []string{"-A", "-s", srcIf, "-d", dstIf}
	if echo {
		args = append(args, "-e")
	}
	if routeIncoming {
		args = append(args, "-i")
	}
	return exec.Command(cangw, args...)
}
