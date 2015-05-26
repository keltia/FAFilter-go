// main.go
//
// Main driver for FAFilter, the Go version of the Flightaware data filtering
// utility written by J. Van Meenen.
//
// Copyright 2015 © by Ollivier Robert for the EEC
//

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"encoding/json"
)

var (
	timeStats	TimeStats
	recordStats	RecordStats

	fhOut	*os.File
)

// Process one file at a time
func processFile(file string, out *os.File) error {
	fh, err := os.Open(file)
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		// each line is a json record
		line := scanner.Text()

		var record FArecord

		// must convert to []byte before handing over to json.Unmarshal
		s_line := []byte(line)
		err := json.Unmarshal(s_line, &record)
		if err != nil {
			return err
		}
		// handover to our checkRecord
		good, err :=  record.checkRecord()
		if err != nil {
			return err
		}
		if good {
			_, err = fmt.Fprintf(out, "%s\n", line)
		}

		if err != nil {
			return err
		}
	}

	return err
}

// Starts here
func main() {
	flag.Usage = Usage
	flag.Parse()

	if flag.Arg(0) == "" {
		fmt.Fprintln(os.Stderr, "Error: you must specify files")
		Usage()
	}
	// Remaining arguments are in flag.Args()
	if fVerbose {
		fmt.Printf("%v\n", flag.Args())
	}

	var err error

	if fFileOut != "" {
		fhOut, err = os.Create(fFileOut)
	} else {
		fhOut = os.Stdout
	}

	for i := 0; i < len(string(flag.Arg(i))); i++ {
		if fVerbose {
			fmt.Fprintf(os.Stderr, "Reading %v…\n", flag.Arg(i))
		}
		err = processFile(flag.Arg(i), fhOut)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %v", flag.Arg(i))
		}
	}
}
