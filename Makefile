export GO111MODULE 	= on

SHELL 				=  /bin/bash -o pipefail
PWD 				=  $(shell pwd)
GO 					?= go
DC 					?= docker-compose
DU 					?= du

all: ctrlt cn

ctrlt: get
	$(GO) build ./cmd/ctrlt
	@$(DU) -k $@

cn: get
	$(GO) build ./cmd/cn
	@$(DU) -k $@

ctrlt-static: get
	$(GO) build -a -tags netgo -ldflags '-s -w -extldflags "-static"' ./cmd/ctrlt
	@$(DU) -k ctrlt

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

run:
	$(DC) pull immudb
	$(DC) up --build --remove-orphans --abort-on-container-exit --exit-code-from ctrlt
	$(DC) down

clean:
	$(RM) ctrlt cn
	$(GO) clean ./...
	$(GO) clean -testcache

.PHONY: all ctrlt cnt ctrlt-static cn-static test integration-test get vendor tidy generate run clean
