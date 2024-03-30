SRC_DIR=./cmd/dockercan

BINARY_DIR=./bin/dockercan/
BINARY_NAME=dockercan

SCRIPTS_DIR=./scripts/

build:
	go build -o ${BINARY_DIR}${BINARY_NAME} ${SRC_DIR}

run: build
	sudo ${BINARY_DIR}${BINARY_NAME}

install: build
	sudo ./scripts/install/install.sh

uninstall:
	sudo ./scripts/uninstall/uninstall.sh

clean:
	rm -fdr ${BINARY_DIR}
	go clean