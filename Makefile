PKG ?= github.com/tasdomas/pix

GOBIN = $(abspath ./bin)
PKGS := $(shell go list github.com/tasdomas/pix/... | grep -v /vendor/)
VERSION := 0.0.1
APP := pix
IMAGE = $(APP):$(VERSION)

all: deps

.PHONY: deps
deps: $(GOBIN)/rice

$(GOBIN)/vendetta:
	go get github.com/dpw/vendetta

$(GOBIN)/rice:
	go install github.com/GeertJohan/go.rice/rice

update-deps: $(GOBIN)/vendetta
	$(GOBIN)/vendetta -n $(PKG) -u

.PHONY: check
check:
	go test $(PKGS)

.PHONY: run
run:
	go run main.go -cfg cfg.yaml

print-%:
	@echo '$*=$($*)'

pix: embed
	go build ./cmd/pix

.PHONY: embed
embed: $(GOBIN)/rice $(PKGS)

github.com/tasdomas/%:
	$(GOBIN)/rice embed-go -i $@

docker: pix
	docker build ./ -t ${IMAGE}
