BINARY_DIR=./bin/
BINARY_NAME=netplugin

build:
	go build -o ${BINARY_DIR}${BINARY_NAME} ./cmd/netplugin

run: build
	sudo ${BINARY_DIR}${BINARY_NAME}

clean:
	rm -fdr ./bin
	go clean