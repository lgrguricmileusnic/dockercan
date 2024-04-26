SRC_DIR=./cmd/

BUILD_DIR=./build/

DRIVER=dockercan
DRIVER_REMOTE=dockercan_remote

SCRIPTS_DIR=./scripts/

ADDR="127.0.0.1:5555"



build: dockercan dockercan_remote

dockercan_remote:
	go build -o ${BUILD_DIR}${DRIVER_REMOTE} ${SRC_DIR}${DRIVER_REMOTE}

dockercan:
	go build -o ${BUILD_DIR}${DRIVER} ${SRC_DIR}${DRIVER}

run: dockercan_remote
	@sudo ${BUILD_DIR}${DRIVER_REMOTE} -addr ${ADDR}

install: dockercan_remote
	@sudo ./scripts/install/install.sh

uninstall:
	@sudo ./scripts/uninstall/uninstall.sh

build_rootfsimage:
	docker build . -t dockercan_rootfs

build_plugin: build_rootfsimage
	@sudo ./scripts/build_plugin/build_plugin.sh



.PHONY: build run install uninstall build_rootfsimage

clean:
	sudo rm -fdr ./build ./plugin
	go clean