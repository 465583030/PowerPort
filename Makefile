.PHONY: default
export GOPATH:=$(shell pwd)
export GOBIN:=$(shell pwd)/bin
default: PowerPortd

all: PowerPortd PowerPorts PowerPortc
	
clean:
	go clean -i -r ./

PowerPortd:
	go install PowerPortd

PowerPorts:
	go install PowerPorts

PowerPortc:
	go install PowerPortc

install:

uninstall:
	