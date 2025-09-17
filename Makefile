BINARY_VERSION=1.0.4
BINARY_NAME=snmp-proxy

build:
	go build -o $(BINARY_NAME) -ldflags "-X main.version=${BINARY_VERSION}" main.go

run: build
	./${BINARY_NAME}

test:
	go test -v ./...

clean:
	go clean
	rm -f ${BINARY_NAME}
