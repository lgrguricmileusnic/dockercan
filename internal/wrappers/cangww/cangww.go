package cangww

import (
	"os/exec"
)

const cangw = "cangw"

func AddRule(srcIf string, dstIf string, echo bool, canfd bool, routeIncoming bool) *exec.Cmd {
	args := []string{"-A", "-s", srcIf, "-d", dstIf}
	appendFlags(args, echo, canfd, routeIncoming)

	return exec.Command(cangw, args...)
}

func RemoveRule(srcIf string, dstIf string, echo bool, canfd bool, routeIncoming bool) *exec.Cmd {
	args := []string{"-D", "-s", srcIf, "-d", dstIf}
	appendFlags(args, echo, canfd, routeIncoming)

	return exec.Command(cangw, args...)
}

func appendFlags(args []string, echo bool, canfd bool, routeIncoming bool) []string {
	if echo {
		args = append(args, "-e")
	}
	if routeIncoming {
		args = append(args, "-i")
	}
	if canfd {
		args = append(args, "-X")
	}
	return args
}
