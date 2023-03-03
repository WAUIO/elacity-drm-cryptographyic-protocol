.PHONY:

NPROC=$(shell grep -c 'processor' /proc/cpuinfo)
MAKEFLAGS+=-j$(NPROCS)

TAG=$(shell git describe --tags)
PROTOVERSION=1.0
TARGET=output/crypto.protocol.wasm
SRC=/src/build/contracts
output=$(shell mktemp -u --tmpdir=$(PWD)/output)
root=$(shell go env GORROT)

all: wasm

fmt:
	@gofmt -l -w ./..

test:
	go test -race -v $(shell go list ./... | grep -v /cmd | grep -v /runtime)

wasm: fmt
	GO111MODULE=on CGO_ENABLED=0 GOOS=js GOARCH=wasm go build \
	 -ldflags "-s -w -X github.com/elacity/crypto-protocol/core.Version=$(PROTOVERSION) -X github.com/elacity/crypto-protocol/core.Tag=$(TAG)" \
	 -o $(TARGET) ./cmd/wasm/main.go

wasm-gz: wasm
	gzip -9 -c $(TARGET) > $(TARGET).gz

tiny: fmt
	GO111MODULE=on CGO_ENABLED=0 GOOS=js GOARCH=wasm tinygo build \
	 -target=wasm \
	 -ldflags "-X github.com/elacity/crypto-protocol/core.Version=$(PROTOVERSION) -X github.com/elacity/crypto-protocol/core.Tag=$(TAG)" \
	 -o $(TARGET) ./cmd/wasm/main.go

gen-contracts: fmt
	cat $(SRC)/AuthorityGateway.json | jq .abi > $(output)/AuthorityGateway.abi
	abigen --abi $(output)/AuthorityGateway.abi --pkg=contracts --type AuthorityGateway --out=./contracts/authority.go
	cat $(SRC)/IOperative.json | jq .abi > $(output)/IOperative.abi
	abigen --abi $(output)/IOperative.abi --pkg=contracts --type Operative --out=./contracts/operative.go
	rm -rf $(output)

contracts: gen-contracts fmt	

http: fmt
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	 -gcflags="all=-N -l" \
	 -ldflags "-X github.com/elacity/crypto-protocol/core.Version=$(PROTOVERSION) -X github.com/elacity/crypto-protocol/core.Tag=$(TAG)" \
	 -o $(TARGET) ./cmd/http