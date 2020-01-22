export GO111MODULE 	= on

SHELL 				=  /bin/bash -o pipefail
PWD 				=  $(shell pwd)
GO 					?= go
DU 					?= du

all: cn

cn: get
	$(GO) build ./cmd/cn
	@$(DU) -k $@

cn-static: get
	$(GO) build -a -tags netgo -ldflags '-s -w -extldflags "-static"' ./cmd/cn
	@$(DU) -k cn

test: get
	$(GO) vet ./...
	$(GO) test --race ./...

integration-test: get
	$(GO) test -tags integration -run Integration ./...

get:
	$(GO) get ./...

vendor:
	$(GO) mod vendor

tidy:
	$(GO) fix ./...
	$(GO) fmt ./...
	$(GO) mod tidy

generate:
	$(GO) generate ./...

clean:
	$(RM) cn
	$(GO) clean ./...
	$(GO) clean -testcache

.PHONY: all cn cn-static test integration-test get vendor tidy generate run clean
