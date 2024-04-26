# CAN network plugin for Docker and Docker Compose
## About
Docker network plugin for connecting containers over vcan and vxcan interfaces. 

Rewritten [can4docker](https://gitlab.com/chgans/can4docker/-/tree/master/can4docker) using [offical Go helper](https://github.com/docker/go-plugins-helpers/tree/master/network). Added another way of connecting containers and CAN FD support, see driver options below. Plugin hides created interfaces from the user in a separate network namespace.

## Requirements

Requires standard Go toolchain for build (version 1.22).

Plugin relies on these kernel modules for networking:
- can_gw
- vxcan

Wraps around the `ip` command for configuring interfaces and namespaces.

## Installation and usage
### Run plugin without installation

```
make run # Starts remote driver at localhost:4343
```

**OR**

```
make run ADDR=<address>:<port>
```

### Install plugin

#### Install plugin from dockerhub

```
docker plugin install lovrogm/dockercan:latest
docker plugin enable  lovrogm/dockercan:latest
```

#### Build and install as a Systemd service
Run commands:
```bash
git clone https://github.com/lgrguricmileusnic/dockercan.git
cd dockercan
make install
```
And follow setup prompts.

This installs and starts the plugin at `localhost:4343`. Edit dockercan.service file to modify address and port.


##### Uninstall

``` 
make uninstall
```

### Creating networks

#### Docker CLI

##### Dockerhub installation
```
docker network create -o centralised={true|false} -o canfd={true|false} -o host_if=<if_name> --driver lovrogm/dockerhub:latest <network_name>
```

##### Systemd installation or manual run
```
docker network create -o centralised={true|false} -o canfd={true|false} -o host_if=dcan0 --driver dockercan_remote <network_name>
```

##### 

#### Docker Compose
Example compose files are available in [deployments/example_compose](https://github.com/lgrguricmileusnic/dockercan/blob/master/deployments/example_compose/compose.yml).

### Driver options
**Available driver options:**
- centralised
- canfd
- hostif

**centralised:**
- `true`  :
  - Creates a single vcan interface which acts as a bus. Each container is then connected to the bus using a vxcan interface pair and cangw rules.
  - requires setting `can_gw` kernel module parameter `max_hops` to `2`
    
    ```
    sudo modprobe can_gw max_hops=2
    ```
- `false` : (default)
  - Interconnects all containers using vxcan and cangw, resulting in a greater number of cangw rules.
  - requires setting `can_gw` kernel module parameter `max_hops=1` (default)

**canfd**
- enables or disables CAN FD support (default: `false`)

**hostif**
- specify name for interface on host
