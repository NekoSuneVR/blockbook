SHELL = /bin/bash
VERSION ?= devel
GITCOMMIT = $(shell cd /src && git config --global --add safe.directory /src && git describe --always --dirty)
BUILDTIME = $(shell date --iso-8601=seconds)
LDFLAGS := -X github.com/decenomy/blockbook/common.version=$(VERSION) -X github.com/decenomy/blockbook/common.gitcommit=$(GITCOMMIT) -X github.com/decenomy/blockbook/common.buildtime=$(BUILDTIME)
BLOCKBOOK_BASE := $(GOPATH)/src/github.com/decenomy
BLOCKBOOK_SRC := $(BLOCKBOOK_BASE)/blockbook
ARGS ?=

all: build tools

build: prepare-sources
	cd $(BLOCKBOOK_SRC) && go build -o $(CURDIR)/blockbook -ldflags="-s -w $(LDFLAGS)" $(ARGS)
	cp $(CURDIR)/blockbook /out/blockbook
	chown $(PACKAGER) /out/blockbook

build-debug: prepare-sources
	cd $(BLOCKBOOK_SRC) && go build -o $(CURDIR)/blockbook -ldflags="$(LDFLAGS)" $(ARGS)
	cp $(CURDIR)/blockbook /out/blockbook
	chown $(PACKAGER) /out/blockbook

tools:
	cp $(CURDIR)/{ldb,sst_dump} /out
	chown $(PACKAGER) /out/{ldb,sst_dump}

test: prepare-sources
	cd $(BLOCKBOOK_SRC) && go test -tags 'unittest' `go list ./... | grep -vP '^github.com/decenomy/blockbook/(contrib|tests)'` $(ARGS)

test-integration: prepare-sources
	cd $(BLOCKBOOK_SRC) && go test -tags 'integration' `go list github.com/decenomy/blockbook/tests/...` $(ARGS)

test-all: prepare-sources
	cd $(BLOCKBOOK_SRC) && go test -tags 'unittest integration' `go list ./... | grep -v '^github.com/decenomy/blockbook/contrib'` $(ARGS)

prepare-sources:
	@ [ -n "`ls /src 2> /dev/null`" ] || (echo "/src doesn't exist or is empty" 1>&2 && exit 1)
	rm -rf $(BLOCKBOOK_SRC)
	mkdir -p $(BLOCKBOOK_BASE)
	cp -r /src $(BLOCKBOOK_SRC)
	cd $(BLOCKBOOK_SRC) && go mod download
	sed -i 's/wsMessageSizeLimit\ =\ 15\ \*\ 1024\ \*\ 1024/wsMessageSizeLimit = 50 * 1024 * 1024/g' $(GOPATH)/pkg/mod/github.com/ethereum/go-ethereum*/rpc/websocket.go
