package driver

import (
	"dockercan/internal/netns"
	"fmt"

	"github.com/docker/go-plugins-helpers/network"
)

type Endpoint struct {
	vxcanHost string
	vxcanCont string
}

type Network struct {
	ns       string
	enpoints map[string]Endpoint
}

type Driver struct {
	networks map[string]Network
}

func (d *Driver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	rs := &network.CapabilitiesResponse{Scope: network.LocalScope}
	return rs, nil
}

func (d *Driver) CreateNetwork(rq *network.CreateNetworkRequest) error {
	// create namespace for hiding vxcan interfaces on the host
	nsName := fmt.Sprintf("canns_%s", rq.NetworkID[:4])

	err := netns.CreateNetworkNamespace(nsName)
	if err != nil {
		return fmt.Errorf("error creating network namespace: %s", err.Error())
	}

	d.networks[rq.NetworkID] = Network{nsName, map[string]Endpoint{}}
	return nil
}

func (d *Driver) DeleteNetwork(rq *network.DeleteNetworkRequest) error {
	nsName := d.networks[rq.NetworkID].ns
	err := netns.DeleteNetworkNamespace(nsName)
	if err != nil {
		return fmt.Errorf("error deleteing network namespace: %s", err.Error())
	}

	delete(d.networks, nsName)
	return nil
}

func (d *Driver) CreateEndpoint(rq *network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	eid := rq.EndpointID
	nid := rq.NetworkID

	net, ok := d.networks[nid]

	if !ok {
		return nil, fmt.Errorf("network with id %s does not exist", nid)
	}

	net.enpoints[eid] = Endpoint{}
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

// unimplemented stubs
func (d *Driver) FreeNetwork(*network.FreeNetworkRequest) error {
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

func (d *Driver) AllocateNetwork(rq *network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	rs := &network.AllocateNetworkResponse{Options: rq.Options}
	return rs, nil
}
