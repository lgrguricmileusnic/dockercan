package main

import (
	"dockercan/internal/driver"
	"os/user"
	"strconv"

	"github.com/docker/go-plugins-helpers/network"
)

func main() {
	d := driver.Driver{}
	h := network.NewHandler(&d)
	u, err := user.Current()

	if err != nil {
		panic(err)
	}

	gid, err := strconv.Atoi(u.Gid)

	if err != nil {
		panic(err)
	}

	h.ServeUnix("/run/docker/plugins/dockercan.sock", gid)
}
