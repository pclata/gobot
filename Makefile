GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test

install:
	$(GOINSTALL)

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/gobot
	rm -f gobot

test:
	$(GOTEST) -v -cover