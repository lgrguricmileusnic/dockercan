#!/bin/bash

id=$(docker create dockercan_rootfs:latest true)
sudo mkdir -p plugin/rootfs
sudo docker export "$id" | sudo tar -x -C plugin/rootfs
docker rm -vf "$id"
docker rmi dockercan_rootfs:latest

cp scripts/build_plugin/config.json plugin/config.json