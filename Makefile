
# The binary to build (just the basename).
IMAGE := miner
#ALL_ARCH := amd64 arm arm64 ppc64le
ALL_ARCH ?= amd64

# This repo's root import path (under GOPATH).
# PKG ?= $(shell realpath --relative-to=${GOPATH}/src `pwd`)
PKG ?= github.com/innoq-blockchain-go-1

VERSION ?= "latest"

# go source files without tests, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -name '*_test.go' -not -path './vendor/*')

.PHONY: all-image
all-image: $(addprefix image-, $(ALL_ARCH))

.PHONY: image-%
image-%: ARCH = $*
image-%: Dockerfile-%
	docker build -t $(IMAGE):$(VERSION) -f Dockerfile-$(ARCH) .

Dockerfile-%: ARCH = $*
Dockerfile-%: Dockerfile.in
	@sed \
		-e 's|ARG_PKG|$(PKG)|g' \
		-e 's|ARG_ARCH|$(ARCH)|g' \
		Dockerfile.in > $@

## dev stuff
Gopkg.toml:
	dep init

.PHONY: vendor
vendor: Gopkg.toml
	dep ensure

.PHONY: install
install: vendor
	go install ./...

.PHONY: run
run: vendor
	go run $(SRC)
