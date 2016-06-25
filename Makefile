PKG ?= github.com/tasdomas/pix

GOBIN = $(abspath ./bin)

TESTS := $(shell go list github.com/tasdomas/pix/... | grep -v /vendor/)

all: deps

.PHONY: deps
deps: $(GOBIN)/rice

$(GOBIN)/rice:
	go install github.com/GeertJohan/go.rice/rice

.PHONY: check
check:
	go test $(TESTS)

.PHONY: run
run:
	go run main.go

print-%:
	@echo '$*=$($*)'
