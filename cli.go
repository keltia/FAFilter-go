// cli.go
//
// Everything related to command-line flag handling

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// cli
	fStartTime	string
	fEndTime	string
	fFileOut	string
	fAircraftId	string
	fHexid		string
	fUpdateType	string

	fVerbose	bool
)

// my usage string
const (
	cliUsage	= `
Usage: %s [-o file] [-b time -e time] [-a reg|-x reg|-t TYPE] [-v] files...
`
)

// Redefine Usage
var Usage = func() {
        fmt.Fprintf(os.Stderr, cliUsage, os.Args[0])
        flag.PrintDefaults()
}

// called by flag.Parse()
func init() {
	// cli
	flag.StringVar(&fStartTime, "b", "", "Start time")
	flag.StringVar(&fEndTime, "e", "", "End time")
	flag.StringVar(&fFileOut, "o", "", "Output into file")
	flag.StringVar(&fAircraftId, "a", "", "AircraftId regexp")
	flag.StringVar(&fHexid, "x", "", "Hexid regexp")
	flag.StringVar(&fUpdateType, "t", "", "Update type filter")

	flag.BoolVar(&fVerbose, "v", false, "Be verbose")

}
