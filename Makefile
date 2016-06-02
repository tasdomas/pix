PKG=github.com/tasdomas/pixserver

GOBIN=$(abspath ./bin)

all: deps

.PHONY: deps
deps: $(GOBIN)/rice

$(GOBIN)/rice:
	go get github.com/GeertJohan/go.rice/rice

.PHONY: check
check:
	@echo $(go list PKG/... | grep -v /vendor/)

.PHONY: run
run:
	go run main.go
