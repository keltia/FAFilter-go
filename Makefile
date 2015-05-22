# Main Makefile for FAFilter

DEST=   bin

all:    ${DEST}/FAFilter

clean:
	rm -f ${DEST}/FAFilter

${DEST}/FAFilter:    main.go types.go
	go build -v -o $@

