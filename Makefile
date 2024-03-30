SRC_DIR=./cmd/dockercan

BUILD_DIR=./build/dockercan/
BINARY_NAME=dockercan

SCRIPTS_DIR=./scripts/

build:
	go build -o ${BUILD_DIR}${BINARY_NAME} ${SRC_DIR}

run: build
	sudo ${BUILD_DIR}${BINARY_NAME}

install: build
	sudo ./scripts/install/install.sh

uninstall:
	sudo ./scripts/uninstall/uninstall.sh

clean:
	rm -fdr ${BUILD_DIR}
	go clean