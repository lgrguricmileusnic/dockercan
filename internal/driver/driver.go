package driver

import (
	"github.com/docker/go-plugins-helpers/network"
)

type Driver struct {
	Name string
}

func (d *Driver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	res := &network.CapabilitiesResponse{Scope: network.LocalScope}
	return res, nil
}
func (d *Driver) CreateNetwork(*network.CreateNetworkRequest) error {
	// TODO create namespace
	return nil
}
func (d *Driver) AllocateNetwork(*network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	res := &network.AllocateNetworkResponse{}
	return res, nil
}
func (d *Driver) DeleteNetwork(*network.DeleteNetworkRequest) error {
	return nil
}
func (d *Driver) FreeNetwork(*network.FreeNetworkRequest) error {
	return nil
}
func (d *Driver) CreateEndpoint(*network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	return nil, nil
}
func (d *Driver) DeleteEndpoint(*network.DeleteEndpointRequest) error {
	return nil
}
func (d *Driver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	return nil, nil
}
func (d *Driver) Join(*network.JoinRequest) (*network.JoinResponse, error) {
	return nil, nil
}
func (d *Driver) Leave(*network.LeaveRequest) error {
	return nil
}
func (d *Driver) DiscoverNew(*network.DiscoveryNotification) error {
	return nil
}
func (d *Driver) DiscoverDelete(*network.DiscoveryNotification) error {
	return nil
}
func (d *Driver) ProgramExternalConnectivity(*network.ProgramExternalConnectivityRequest) error {
	return nil
}
func (d *Driver) RevokeExternalConnectivity(*network.RevokeExternalConnectivityRequest) error {
	return nil
}
