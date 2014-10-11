.PHONY: .FORCE
GO=go
GOPATH := $(shell pwd)

PROGS = main

all: $(PROGS)

$(PROGS):
	GOPATH=$(GOPATH) $(GO) install main

clean:
	rm -rf bin pkg
