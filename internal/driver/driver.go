package driver

import (
	"dockercan/internal/wrappers"
	"fmt"
	"log"

	"github.com/docker/go-plugins-helpers/network"
)

type Endpoint struct {
	vxcanHidden    string
	vxcanContainer string
}

type Network struct {
	ns       string
	enpoints map[string]Endpoint
}

type Driver struct {
	networks map[string]Network
}

func NewDriver() (*Driver, error) {
	d := Driver{make(map[string]Network)}
	return &d, nil
}

func (d *Driver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	log.Println("GetCapabilities: received")
	rs := &network.CapabilitiesResponse{Scope: network.LocalScope}

	return rs, nil
}

func (d *Driver) CreateNetwork(rq *network.CreateNetworkRequest) error {
	// create namespace for hiding vxcan interfaces on the host
	log.Println("CreateNetwork: CreateNetwork received")
	nsName := fmt.Sprintf("canns_%s", rq.NetworkID[:4])

	log.Printf("CreateNetwork: Creating network namespace %s", nsName)

	err := wrappers.CreateNetworkNamespace(nsName)
	if err != nil {
		return fmt.Errorf("CreateNetwork: error creating network namespace: %s", err.Error())
	}

	d.networks[rq.NetworkID] = Network{nsName, map[string]Endpoint{}}
	return nil
}

func (d *Driver) DeleteNetwork(rq *network.DeleteNetworkRequest) error {
	nsName := d.networks[rq.NetworkID].ns

	log.Println("DeleteNetwork: DeleteNetwork received")
	log.Printf("DeleteNetwork: Deleting network namespace %s", nsName)

	err := wrappers.DeleteNetworkNamespace(nsName)
	if err != nil {
		return fmt.Errorf("DeleteNetwork: error deleteing network namespace: %s", err.Error())
	}

	delete(d.networks, nsName)
	return nil
}

func (d *Driver) CreateEndpoint(rq *network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	eid := rq.EndpointID
	nid := rq.NetworkID
	log.Println("CreateEndpoint: CreateEndpoint received")

	net, ok := d.networks[nid]
	if !ok {
		return nil, fmt.Errorf("CreateEndpoint: network with id %s does not exist", nid)
	}
	ep := Endpoint{vxcanHidden: fmt.Sprintf("hcan%s", eid[:6]), vxcanContainer: fmt.Sprintf("ccan%s", eid[:6])}

	net.enpoints[eid] = ep

	err := wrappers.CreateInterfacePair(net.enpoints[eid].vxcanHidden, net.enpoints[eid].vxcanContainer, wrappers.Vxcan)
	if err != nil {
		return nil, fmt.Errorf("CreateEndpoint: error creating interface pair : %s", err.Error())
	}

	err = wrappers.MoveInterfaceToNamespace(ep.vxcanHidden, net.ns)
	if err != nil {
		return nil, fmt.Errorf("CreateEndpoint: error moving interface %s to namespace %s: %s", ep.vxcanHidden, net.ns, err.Error())
	}

	return &network.CreateEndpointResponse{}, nil
}

func (d *Driver) DeleteEndpoint(rq *network.DeleteEndpointRequest) error {
	eid := rq.EndpointID
	nid := rq.NetworkID
	log.Println("DeleteEndpoint: DeleteEndpoint received")

	net, ok := d.networks[nid]

	if !ok {
		return fmt.Errorf("DeleteEndpoint: network with id %s does not exist", nid)
	}

	_, ok = net.enpoints[eid]
	if !ok {
		return fmt.Errorf("DeleteEndpoint: endpoint with id %s does not exist", eid)
	}

	delete(d.networks, eid)

	return nil
}

func (d *Driver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	log.Println("EndpointInfo: EndpointInfo received")
	return nil, nil
}

func (d *Driver) Join(*network.JoinRequest) (*network.JoinResponse, error) {
	log.Println("Join: Join received")
	return nil, nil
}

func (d *Driver) Leave(*network.LeaveRequest) error {
	log.Println("Leave: Leave received")
	return nil
}

// unimplemented stubs
func (d *Driver) AllocateNetwork(rq *network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	log.Println("AllocateNetwork: AllocateNetwork received")
	rs := &network.AllocateNetworkResponse{Options: rq.Options}
	return rs, nil
}

func (d *Driver) FreeNetwork(*network.FreeNetworkRequest) error {
	log.Println("FreeNetwork: FreeNetwork received")
	return nil
}
func (d *Driver) DiscoverNew(*network.DiscoveryNotification) error {
	log.Println("DiscoverNew: DiscoverNew received")
	return nil
}

func (d *Driver) DiscoverDelete(*network.DiscoveryNotification) error {
	log.Println("DiscoverDelete: DiscoverDelete received")
	return nil
}

func (d *Driver) ProgramExternalConnectivity(*network.ProgramExternalConnectivityRequest) error {
	log.Println("ProgramExternalConnectivity: ProgramExternalConnectivity received")

	return nil
}

func (d *Driver) RevokeExternalConnectivity(*network.RevokeExternalConnectivityRequest) error {
	log.Println("RevokeExternalConnectivity: RevokeExternalConnectivity received")
	return nil
}
