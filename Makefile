SRC_DIR=./cmd/${BINARY_NAME}/

BUILD_DIR=./build/${BINARY_NAME}/
BINARY_NAME=dockercan

SCRIPTS_DIR=./scripts/

ADDR="127.0.0.1:4343"

.PHONY: build

build:
	go build -o ${BUILD_DIR}${BINARY_NAME} ${SRC_DIR}

run: build
	@sudo ${BUILD_DIR}${BINARY_NAME} -addr ${ADDR}

install: build
	@sudo ./scripts/install/install.sh

uninstall:
	@sudo ./scripts/uninstall/uninstall.sh

clean:
	rm -fdr ${BUILD_DIR}
	go clean