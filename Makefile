# Main Makefile for FAFilter

DEST=   bin

all:    ${DEST}/FAFilter

clean:
	rm -f ${DEST}/FAFilter

${DEST}/FAFilter:    main.go types.go farecord.go location.go cli.go
	go build -v -o $@

