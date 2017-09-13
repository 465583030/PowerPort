.PHONY: default
export GOPATH:=$(shell pwd)
export GOBIN:=$(shell pwd)/bin
default: PowerPortd

all: PowerPortd PowerPorts PowerPortc
	
clean:
	go clean -i -r ./

PowerPortd: simplejson
	go install main/PowerPortd

PowerPorts:
	go install main/PowerPorts

PowerPortc:
	go install main/PowerPortc

simplejson:
	go get github.com/bitly/go-simplejson

install:

uninstall:
	