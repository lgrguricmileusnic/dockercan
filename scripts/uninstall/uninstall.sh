#!/bin/bash
# This script should be started using Makefile from repo root directory

uninstall () {
    systemctl disable dockercan.service
    systemctl stop dockercan.service

    rm /etc/systemd/system/dockercan.service
    rm /usr/lib/docker/dockercan_tcp

    systemctl daemon-reload
}

while true; do
read -p "Uninstall dockercan systemd service? (y/n) " yn

case $yn in 
	[yY] ) uninstall;
		break;;
	*  ) 
		exit;;
esac

done
exit