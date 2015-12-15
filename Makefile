# Main Makefile for FAFilter
#
# Copyright 2015 © by Ollivier Robert for the EEC
#

GOBIN=   ${GOPATH}/bin

all: main.go types.go farecord.go location.go cli.go
	go build -v ./...
	go test -v ./...

install:
	go install -v

clean:
	go clean -v

push:
	git push --all
	git push --all upstream
	git push --all backup
	git push --tags
	git push --tags upstream
	git push --tags backup
