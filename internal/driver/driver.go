package driver

import (
	"fmt"

	"github.com/docker/go-plugins-helpers/network"
)

type Driver struct {
	Name string
}

func (d *Driver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	return nil, nil
}
func (d *Driver) CreateNetwork(*network.CreateNetworkRequest) error {
	fmt.Println("Hello")
	return nil
}
func (d *Driver) AllocateNetwork(*network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	fmt.Println("Hello")
	return nil, nil
}
func (d *Driver) DeleteNetwork(*network.DeleteNetworkRequest) error {
	fmt.Println("Hello")
	return nil
}
func (d *Driver) FreeNetwork(*network.FreeNetworkRequest) error {
	fmt.Println("Hello")
	return nil
}
func (d *Driver) CreateEndpoint(*network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	fmt.Println("Hello")
	return nil, nil
}
func (d *Driver) DeleteEndpoint(*network.DeleteEndpointRequest) error {
	fmt.Println("Hello")
	return nil
}
func (d *Driver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	fmt.Println("Hello")
	return nil, nil
}
func (d *Driver) Join(*network.JoinRequest) (*network.JoinResponse, error) {
	fmt.Println("Hello")
	return nil, nil
}
func (d *Driver) Leave(*network.LeaveRequest) error {
	fmt.Println("Hello")
	return nil
}
func (d *Driver) DiscoverNew(*network.DiscoveryNotification) error {
	fmt.Println("Hello")
	return nil
}
func (d *Driver) DiscoverDelete(*network.DiscoveryNotification) error {
	fmt.Println("Hello")
	return nil
}
func (d *Driver) ProgramExternalConnectivity(*network.ProgramExternalConnectivityRequest) error {
	fmt.Println("Hello")
	return nil
}
func (d *Driver) RevokeExternalConnectivity(*network.RevokeExternalConnectivityRequest) error {
	fmt.Println("Hello")
	return nil
}
