.DEFAULT_GOAL = protoc

build:
	@go build .

protoc: build
	@PATH=${CURDIR}:${PATH} protoc -I${CURDIR}/testdata --go_out=$(shell go env GOPATH)/src --go-jsonshim_out=$(shell go env GOPATH)/src $(shell find ${CURDIR}/testdata -name '*.proto')
	@cat testdata/generatedpb/types_jsonshim.pb.go testdata/generatedpb/external_jsonshim.pb.go
