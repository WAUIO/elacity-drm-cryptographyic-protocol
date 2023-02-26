.PHONY:

NPROC=$(shell grep -c 'processor' /proc/cpuinfo)
MAKEFLAGS+=-j$(NPROCS)

TAG=$(shell git describe --tags)
PROTOVERSION=1.0
TARGET=output/crypto.protocol.wasm

all: wasm

fmt:
	@gofmt -l -w ./..

test:
	go test -race -v $(shell go list ./... | grep -v /cmd | grep -v /runtime)

wasm: fmt
	GO111MODULE=on CGO_ENABLED=0 GOOS=js GOARCH=wasm go build \
	 -ldflags "-X github.com/elacity/crypto-protocol/protocol.Version=$(PROTOVERSION) -X github.com/elacity/crypto-protocol/protocol.Tag=$(TAG)" \
	 -o $(TARGET) ./cmd/wasm/main.go

tiny: fmt
	GO111MODULE=on CGO_ENABLED=0 GOOS=js GOARCH=wasm tinygo build \
	 -target=wasm \
	 -ldflags "-X github.com/elacity/crypto-protocol/protocol.Version=$(PROTOVERSION) -X github.com/elacity/crypto-protocol/protocol.Tag=$(TAG)" \
	 -o $(TARGET) ./cmd/wasm/main.go

http: fmt
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	 -gcflags="all=-N -l" \
	 -ldflags "-X github.com/elacity/crypto-protocol/protocol.Version=$(PROTOVERSION) -X github.com/elacity/crypto-protocol/protocol.Tag=$(TAG)" \
	 -o $(TARGET) ./cmd/http	 