// cli.go
//
// Everything related to command-line flag handling

package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"time"
)

var (
	// cli
	fStartTime	string
	fEndTime	string
	fFileOut	string
 	fAircraftId	string
	fHexid 		string
	fUpdateType	string
	fVerbose	bool

	tsStart		time.Time
	tsEnd		time.Time
)

// my usage string
const (
	cliUsage	= `
Usage: %s [-o file] [-b time -e time] [-a regex|-x regex|-t regex] [-v] files...
`
	TIMEFMT = "2006-01-02 15:04:05"
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
	flag.BoolVar(&fVerbose, "v", false, "Be verbose")

	//Treat these differently
	flag.StringVar(&fAircraftId, "a", "", "AircraftId regexp")
	flag.StringVar(&fHexid, "x", "", "Hexid regexp")
	flag.StringVar(&fUpdateType, "t", "", "Update type regex")

	// Compile and check, if a given regex is invalid, panic()
	if (fAircraftId != "") {
		_ = *regexp.MustCompile(fAircraftId)
	}
	if (fHexid != "") {
		_ = *regexp.MustCompile(fHexid)
	}
	if fUpdateType != "" {
		_ = *regexp.MustCompile(fUpdateType)
	}

	var err  error

	// Check dates
	tsStart, err = time.Parse(TIMEFMT, fStartTime)
	if err != nil {
		fmt.Println(err)
		tsStart = time.Now()
	}
	tsEnd, err = time.Parse(TIMEFMT, fEndTime)
	if err != nil {
		fmt.Println(err)
	}

}
