// cli.go
//
// Everything related to command-line flag handling
//
// Copyright 2015 © by Ollivier Robert for the EEC
//

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
	fGeoFile	string
	gFileList	[]string
	fVerbose	bool

	tsStart		time.Time
	tsEnd		time.Time
)

// my usage string
const (
	cliUsage	= `
Usage: %s [-o file] [-b time -e time] [-a regex|-x regex|-t regex] [-v] [-g f1,f2,...] files...
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
	flag.StringVar(&fStartTime, "b", "2001-01-01 00:00:00", "Start time")
	flag.StringVar(&fEndTime, "e", "2038-01-01 00:00:00", "End time")
	flag.StringVar(&fFileOut, "o", "", "Output into file")
	flag.BoolVar(&fVerbose, "v", false, "Be verbose")
	flag.StringVar(&fGeoFile, "g", "", "Geofile for specific area")

	// Treat these differently
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
