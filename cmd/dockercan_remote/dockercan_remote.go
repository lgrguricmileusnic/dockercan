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
	dPtr, err := driver.NewDriver()

	if err != nil {
		log.Fatalln(err)
	}

	h := network.NewHandler(dPtr)

	log.Printf("Starting docker CAN driver at %s\n", ip)

	err = h.ServeTCP("dockercan_remote", ip, "", nil)

	if err != nil {
		log.Fatal(err)
	}
}
