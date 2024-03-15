package main

import (
	"dockercan/internal/driver"

	"github.com/docker/go-plugins-helpers/network"
)

func main() {
	d := driver.Driver{}
	h := network.NewHandler(&d)
	h.ServeTCP("dockercan", "127.0.0.1", "", nil)
}
