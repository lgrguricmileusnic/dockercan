package driver

import (
	"log"
)

type NetworkOptions struct {
	centralised bool
}

func ExtractNetworkOptions(options map[string]interface{}) (opts NetworkOptions) {
	opts = NetworkOptions{centralised: false}
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

	if cs == "true" {
		opts.centralised = true
	}
	return
}
