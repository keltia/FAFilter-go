# Main Makefile for FAFilter
#
# Copyright 2015 © by Ollivier Robert for the EEC
#

DEST=   bin

all:    ${DEST}/FAFilter

install:
	go install -v

clean:
	go clean -v

${DEST}/FAFilter:    main.go types.go farecord.go location.go cli.go
	go build -v -o $@

