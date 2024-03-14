package main

import (
	"dockercan/internal/driver"
	"fmt"

	"github.com/docker/go-plugins-helpers/network"
)

func main() {
	d := driver.Driver{Name: "name"}
	network.NewHandler(&d)

	fmt.Println(d.Name)

}
