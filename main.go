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
	readFiles	int

	fhOut	*os.File
)

// Process one file at a time
func processFile(file string, out *os.File) error {
	fh, err := os.Open(file)
	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		// Get current line
		line := scanner.Text()

		// each line is a json record
		// must convert to []byte before handing over to json.Unmarshal
		var record FArecord

		s_line := []byte(line)
		if err := json.Unmarshal(s_line, &record); err != nil {
			return err
		}

		recordStats.TotalRead++
		// handover to our checkRecord
		if good :=  record.checkRecord(); good {
			if _, err = fmt.Fprintf(out, "%s\n", line); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing into %s: %v", fFileOut, err)
			}
		} else {
			recordStats.TotalSkipped++
		}
	}

	return err
}

//Print our stats
func printStats() {
	fmt.Fprintf(os.Stderr, "\n%d files read\n", readFiles)
	fmt.Fprintf(os.Stderr, "Lines read: %d\n", recordStats.TotalRead)
	fmt.Fprintf(os.Stderr, "Lines skipped: %d\n", recordStats.TotalSkipped)

	fmt.Fprintf(os.Stderr, "\nTime-related stats:\n")
	fmt.Fprintf(os.Stderr, "  First seen: %s\n", timeStats.FirstSeen.String())
	fmt.Fprintf(os.Stderr, "  Last seen: %s\n", timeStats.LastSeen.String())
	fmt.Fprintf(os.Stderr, "  First selected: %s\n", timeStats.FirstSelected.String())
	fmt.Fprintf(os.Stderr, "  Last selected: %s\n", timeStats.LastSelected.String())

	fmt.Fprintf(os.Stderr, "  Lowest seen: %s\n", timeStats.Lowest.String())
	fmt.Fprintf(os.Stderr, "  Highest seen: %s\n", timeStats.Highest.String())

	fmt.Fprintf(os.Stderr, "\nRecord-related stats:\n")
	fmt.Fprintf(os.Stderr, "  Skipped AircraftId: %d\n", recordStats.SkippedAircraftId)
	fmt.Fprintf(os.Stderr, "  Skipped Hexid: %d\n", recordStats.SkippedHexid)
	fmt.Fprintf(os.Stderr, "  Skipped UpdateType: %d\n", recordStats.SkippedUpdateType)
	fmt.Fprintf(os.Stderr, "  Skipped Geometric: %d\n", recordStats.SkippedGeometric)
	fmt.Fprintf(os.Stderr, "  Skipped Temporal: %d\n", recordStats.SkippedTemporal)
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
	printStats()
}
