// main.go
//
// Main driver for FAFilter, the Go version of the Flightaware data filtering
// utility written by J. Van Meenen.
//
// Copyright 2015 © by Ollivier Robert for the EEC
//

package main

import (
	"flag"
	"time"
	"fmt"
	"os"
)

var (
	timeStats	TimeStats
	recordStats	RecordStats
)

// Implements main conversion
func (line *FArecord) checkRecord() (FArecord, error) {
	var myTimestamp time.Time

	if line.Type == "position" {
		if line.Clock != "" {
			var value int64

			if _, err := fmt.Sscanf(line.Clock, "%d", &value); err != nil {
				myTimestamp = time.Unix(value, 0)
			}
			// gather stats
			if (timeStats.FirstSeen == time.Time{}) {
				timeStats.FirstSeen = myTimestamp
			}
			timeStats.LastSeen = myTimestamp

			if (timeStats.Lowest == time.Time{}) {
				timeStats.Lowest = myTimestamp
				timeStats.Highest = myTimestamp
			} else {
				// check lowest/highest
				if myTimestamp.Before(timeStats.Lowest) {
					timeStats.Lowest = myTimestamp
				}
				if myTimestamp.After(timeStats.Highest) {
					timeStats.Highest = myTimestamp
				}
			}


		}
	}
	return *line, nil
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
}
