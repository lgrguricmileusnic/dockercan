#!/bin/bash
# This script should be started using Makefile from repo root directory

install () {
    cp ./bin/dockercan/dockercan           /usr/lib/docker/dockercan
    cp ./scripts/install/dockercan.service /etc/systemd/system/dockercan.service

    systemctl daemon-reload
    systemctl enable dockercan.service
    systemctl start dockercan.service
}

while true; do
read -p "Install dockercan as a systemd service? (y/n) " yn

case $yn in 
	[yY] ) install;
		break;;
	*  ) 
		exit;;
esac

done
exit