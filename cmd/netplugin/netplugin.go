package main

import (
	"flag"
	"log"

	"dockercan/internal/driver"

	"github.com/docker/go-plugins-helpers/network"
)

func main() {

	ip := "127.0.0.1:4343"

	flag.StringVar(&ip, "addr", ip, "IPv4 address with port")

	flag.Parse()
	d := driver.Driver{}
	h := network.NewHandler(&d)

	log.Printf("Starting docker CAN driver at %s\n", ip)
	err := h.ServeTCP("dockercan", ip, "", nil)

	if err != nil {
		log.Panicln(err)
	}
}
