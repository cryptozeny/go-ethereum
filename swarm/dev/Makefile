.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/ethereumprogpow/ethereumprogpow/cmd/swarm
	go build -o bin/geth     github.com/ethereumprogpow/ethereumprogpow/cmd/geth
	go build -o bin/bootnode github.com/ethereumprogpow/ethereumprogpow/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/ethereumprogpow/ethereumprogpow/swarm/...
