GOBIN=$(abspath ./bin)
GOPATH=$(abspath ./gopath)

all: deps

.PHONY: deps
deps: $(GOBIN)/rice

$(GOBIN)/rice:
	go get github.com/GeertJohan/go.rice/rice
