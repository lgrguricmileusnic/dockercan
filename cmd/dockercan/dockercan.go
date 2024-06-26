package main

import (
	"log"

	"dockercan/internal/driver"

	"github.com/docker/go-plugins-helpers/network"
)

func main() {
	dPtr, err := driver.NewDriver()

	if err != nil {
		log.Fatalln(err)
	}

	h := network.NewHandler(dPtr)
	err = h.ServeUnix("/run/docker/plugins/dockercan.sock", 0)

	if err != nil {
		log.Fatal(err)
	}
}
