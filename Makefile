PKG ?= github.com/tasdomas/pix

GOBIN = $(abspath ./bin)

PKGS := $(shell go list github.com/tasdomas/pix/... | grep -v /vendor/)

all: deps

.PHONY: deps
deps: $(GOBIN)/rice

$(GOBIN)/rice:
	go install github.com/GeertJohan/go.rice/rice

.PHONY: check
check:
	go test $(PKGS)

.PHONY: run
run:
	go run main.go

print-%:
	@echo '$*=$($*)'

build: embed
	go build github.com/tasdomas/pix

.PHONY: embed
embed: $(GOBIN)/rice $(PKGS)

github.com/tasdomas/%:
	rice embed-go -i $@
