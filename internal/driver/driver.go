package driver

import (
	"dockercan/internal/wrappers/cangww"
	"dockercan/internal/wrappers/ipw"
	"fmt"
	"log"

	"github.com/docker/go-plugins-helpers/network"
)

const (
	hIfPrefix string = "hcan_"
	cIfPrefix string = "ccan_"
)

type Endpoint struct {
	vxcanHidden    string
	vxcanContainer string
}

type Network struct {
	ns        string
	vcan      string
	endpoints map[string]Endpoint
	opts      NetworkOptions
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
	// create namespace for hiding vxcan/vcan interfaces on the host
	log.Println("CreateNetwork: CreateNetwork received")
	nsName := fmt.Sprintf("canns_%s", rq.NetworkID[:4])
	vcanName := "canbus"
	log.Printf("CreateNetwork: Creating network namespace %s", nsName)

	err := ipw.CreateNetworkNamespace(nsName).Run()
	if err != nil {
		return fmt.Errorf("CreateNetwork: error creating network namespace: %s", err.Error())
	}

	opts := ExtractNetworkOptions(rq.Options)

	d.networks[rq.NetworkID] = Network{nsName, vcanName, map[string]Endpoint{}, opts}

	if !opts.centralised {
		// if not centralised, containers will be conected p2p using vxcan
		return nil
	}

	err = ipw.CreateInterface(vcanName, ipw.Vcan).Run()
	if err != nil {
		return fmt.Errorf("CreateNetwork: error creating virtual bus interface in namespace %s: %s", nsName, err.Error())
	}

	return nil
}

func (d *Driver) DeleteNetwork(rq *network.DeleteNetworkRequest) error {
	nsName := d.networks[rq.NetworkID].ns

	log.Println("DeleteNetwork: DeleteNetwork received")
	log.Printf("DeleteNetwork: Deleting network namespace %s", nsName)

	err := ipw.DeleteNetworkNamespace(nsName).Run()
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

	ep := Endpoint{
		vxcanHidden:    fmt.Sprintf("%s%s", hIfPrefix, eid[:6]),
		vxcanContainer: fmt.Sprintf("%s%s", cIfPrefix, eid[:6]),
	}
	net.endpoints[eid] = ep

	err := ipw.CreateInterfacePair(net.endpoints[eid].vxcanHidden, net.endpoints[eid].vxcanContainer, ipw.Vxcan).Run()
	if err != nil {
		return nil, fmt.Errorf("CreateEndpoint: error creating interface pair : %s", err.Error())
	}

	err = ipw.MoveInterfaceToNamespace(ep.vxcanHidden, net.ns).Run()
	if err != nil {
		return nil, fmt.Errorf("CreateEndpoint: error moving interface %s to namespace %s: %s", ep.vxcanHidden, net.ns, err.Error())
	}

	log.Printf("CreateEndpoint: Created endpoint received")

	return &network.CreateEndpointResponse{}, nil
}

func (d *Driver) DeleteEndpoint(rq *network.DeleteEndpointRequest) error {
	eid := rq.EndpointID
	nid := rq.NetworkID
	log.Println("DeleteEndpoint: DeleteEndpoint received")

	net, ep, err := NetworkAndEndpointById(nid, eid, d.networks)

	if err != nil {
		return fmt.Errorf("DeleteEndpoint: %s", err.Error())
	}

	err = ipw.ExecCommandInNamespace(net.ns, *ipw.DeleteInterface(ep.vxcanHidden)).Run()
	if err != nil {
		return fmt.Errorf("error deleting interface pair %s:%s from hidden network namespace %s: %s", ep.vxcanHidden, ep.vxcanContainer, net.ns, err.Error())
	}
	delete(d.networks, eid)

	return nil
}

func (d *Driver) Join(rq *network.JoinRequest) (*network.JoinResponse, error) {
	log.Println("Join: Join received")
	nid, eid := rq.NetworkID, rq.EndpointID
	net, ep, err := NetworkAndEndpointById(nid, eid, d.networks)

	if err != nil {
		return nil, fmt.Errorf("Join: %s", err.Error())
	}

	if net.opts.centralised {
		// Connect hidden vxcan pair to vcan bus inside network namespace.
		// 2 * N rules total for a network
		cangww.AddRule(ep.vxcanHidden, net.vcan, true, net.opts.canfd, true)
		cangww.AddRule(net.vcan, ep.vxcanHidden, true, net.opts.canfd, true)
	} else {
		// Connect containers to all existing containers using (N-1) * 2 cangw rules.
		// (N choose 2) * 2 rules total for a network
		for i, e := range net.endpoints {
			if i == eid {
				continue
			}

			cmd := *cangww.AddRule(ep.vxcanHidden, e.vxcanHidden, true, net.opts.canfd, true)
			err := ipw.ExecCommandInNamespace(net.ns, cmd).Run()
			if err != nil {
				return nil, fmt.Errorf("Join: error adding cangw rule %s -> %s: %s", ep.vxcanHidden, e.vxcanHidden, err.Error())
			}

			cmd = *cangww.AddRule(e.vxcanHidden, ep.vxcanHidden, true, net.opts.canfd, true)
			err = ipw.ExecCommandInNamespace(net.ns, cmd).Run()
			if err != nil {
				return nil, fmt.Errorf("Join: error adding cangw rule %s -> %s: %s", e.vxcanHidden, ep.vxcanHidden, err.Error())
			}
		}
	}

	ifName := network.InterfaceName{
		SrcName:   ep.vxcanContainer,
		DstPrefix: hIfPrefix,
	}

	return &network.JoinResponse{InterfaceName: ifName}, nil
}

func (d *Driver) Leave(rq *network.LeaveRequest) error {
	log.Println("Leave: Leave received")
	nid, eid := rq.NetworkID, rq.EndpointID

	net, ep, err := NetworkAndEndpointById(nid, eid, d.networks)
	if err != nil {
		return fmt.Errorf("Leave: %s", err.Error())
	}

	if net.opts.centralised {
		cangww.RemoveRule(ep.vxcanHidden, net.vcan, true, net.opts.canfd, true)
		cangww.RemoveRule(net.vcan, ep.vxcanHidden, true, net.opts.canfd, true)
	} else {
		for i, e := range net.endpoints {
			if i == eid {
				continue
			}

			err := cangww.RemoveRule(ep.vxcanHidden, e.vxcanHidden, true, net.opts.canfd, true).Run()
			if err != nil {
				return fmt.Errorf("Join: error adding cangw rule %s -> %s: %s", ep.vxcanHidden, e.vxcanContainer, err.Error())
			}

			err = cangww.RemoveRule(e.vxcanHidden, ep.vxcanHidden, true, net.opts.canfd, true).Run()
			if err != nil {
				return fmt.Errorf("Join: error adding cangw rule %s -> %s: %s", e.vxcanHidden, ep.vxcanContainer, err.Error())
			}
		}
	}
	return nil
}

// unimplemented stubs
func (d *Driver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	log.Println("EndpointInfo: EndpointInfo received")
	return nil, nil
}

func (d *Driver) AllocateNetwork(rq *network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	log.Println("AllocateNetwork: AllocateNetwork received")
	return nil, nil
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
