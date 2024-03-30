# CAN network plugin for Docker and Docker Compose
## About
Docker network plugin for connecting containers over vcan and vxcan interfaces. 

Rewritten [can4docker](https://gitlab.com/chgans/can4docker/-/tree/master/can4docker) using [offical Go helper](https://github.com/docker/go-plugins-helpers/tree/master/network). Added different ways of connecting containers, CAN FD support. Hides created interfaces from the user in a separate network namespace.

## Requirements

Requires standard Go toolchain for build (version 1.22).

Plugin relies on these kernel modules for networking:
- can_gw
- vxcan

Wraps around the `ip` command for configuring interfaces and namespaces.

## Installation and usage
### Run plugin without installation

```
make run
```
### Install plugin

#### Systemd service installation
Run commands:
```bash
git clone https://github.com/lgrguricmileusnic/dockercan.git
cd dockercan
make install
```
And follow setup prompts.

##### Uninstall

``` 
make uninstall
```

### Creating networks

#### Docker CLI

```
docker network create -o centralised={true|false} -o canfd={true|false} --driver dockercan <network_name>
```

#### Docker Compose
Example compose file available in [deployments/example_compose](https://github.com/lgrguricmileusnic/dockercan/blob/master/deployments/example_compose/compose.yml).

### Driver options
**Available driver options:**
- centralised
- canfd

**centralised:**
- `true`  :
  - Creates a single vcan interface which acts as a bus. Each container is then connected to the bus using a vxcan interface pair and cangw rules.
  - requires setting `can_gw` kernel module parameter `max_hops` to `2`
    
    ```
    sudo modprobe can_gw max_hops=2
    ```
- `false` : (default)
  - Interconnects all containers using vxcan and cangw, resulting in a greater number of cangw rules.
  - requires `max_hops=1`

**canfd**
- enables or disables CAN FD support (default: `false`)

