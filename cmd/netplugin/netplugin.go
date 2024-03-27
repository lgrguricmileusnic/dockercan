package main

import (
	"dockercan/internal/driver"
	"log"

	"github.com/docker/go-plugins-helpers/network"
)

func main() {
	d := driver.Driver{}
	h := network.NewHandler(&d)

	log.Println("Starting CAN docker network driver at 127.0.0.1:1337")
	err := h.ServeTCP("lgm_dockercan", "127.0.0.1:1338", "", nil)

	if err != nil {
		log.Panicln(err)
	}
}
