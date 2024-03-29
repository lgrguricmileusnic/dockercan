package driver

import (
	"fmt"
	"log"
)

type NetworkOptions struct {
	centralised bool
	canfd       bool
}

func ExtractNetworkOptions(options map[string]interface{}) (opts NetworkOptions) {
	opts = NetworkOptions{centralised: false, canfd: false}
	rqOpts, ok := options["com.docker.network.generic"].(map[string]interface{})

	// if request contains no options, assume default can_gw kernel module configuration
	if !ok {
		log.Printf("Error extracting options")
		return
	}

	cs, ok := rqOpts["centralised"].(string)
	// if request contains no 'centralised' option, assume default can_gw kernel module configuration
	if !ok {
		log.Printf("No centralised option, defaulting to p2p")
		return
	}

	opts.centralised = strToBool(cs)

	fds, ok := rqOpts["canfd"].(string)
	// if request contains no 'canfd' option, assume default can_gw
	if !ok {
		log.Printf("No canfd option, defaulting to regular CAN")
		return
	}

	opts.canfd = strToBool(fds)
	return
}

func strToBool(s string) bool {
	return s == "true"
}

func NetworkAndEndpointById(nid, eid string, networks map[string]Network) (*Network, *Endpoint, error) {

	net, ok := networks[nid]
	if !ok {
		return nil, nil, fmt.Errorf("network with id %s does not exist", nid)
	}

	ep, ok := net.endpoints[eid]
	if !ok {
		return nil, nil, fmt.Errorf("endpoint with id %s does not exist", eid)
	}

	return &net, &ep, nil

}
