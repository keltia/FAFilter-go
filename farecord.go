// farecord.go
//
// Main processing on files
//
// Copyright 2015 © by Ollivier Robert for the EEC
//

package main

import (
	"time"
	"fmt"
	"regexp"
	"os"
	"strconv"
)

// Implements main checks
//
// Having several parameters specified on the CLI means AND, not OR because
// we only break on false matches. As long as we match, we keep on.
func (line *FArecord) checkRecord() (valid bool) {
	var myTimestamp time.Time

	valid = false
	if line.Type == "position" {
		// Record timestamps
		if line.Clock != "" {
			var value int64

			if _, err := fmt.Sscanf(line.Clock, "%d", &value); err != nil {
				fmt.Fprintf(os.Stderr, "Invalid value %d\n", value)
				return
			}
			myTimestamp = time.Unix(value, 0)

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

		// Check for -a
		if fAircraftID != "" {
			valid, _ = regexp.MatchString(fAircraftID, line.Ident)
			if !valid {
				recordStats.SkippedAircraftID++
				return
			}
		}

		// Check for -x
		if fHexid != "" {
			valid, _ = regexp.MatchString(fHexid, line.Hexid)
			if !valid {
				recordStats.SkippedHexid++
				return
			}
		}

		// Check for -t
		if fUpdateType != "" {
			valid, _ = regexp.MatchString(fUpdateType, line.UpdateType)
			if !valid {
				recordStats.SkippedUpdateType++
				return
			}
		}

		// Check for -g
		if line.Lat != "" && line.Lon != "" && fGeoFile != "" {
			myLat, _ := strconv.ParseFloat(line.Lat, 64)
			myLon, _ := strconv.ParseFloat(line.Lon, 64)
			myLocation := Location{myLat, myLon}
			//
			// Look at all polygons stored in polygonList
			//
			for _, polygon := range polygonList {
				if len(polygon.P) > 0 {
					if myLocation.pointInPolygon(polygon.P) {
						valid = true
						break
					}
				}
			}
			if !valid {
				recordStats.SkippedGeometric++
				return
			}
		}

		// Everything has been checked
		// fill in stats again
		if (timeStats.FirstSelected == time.Time{}) {
			timeStats.FirstSelected = myTimestamp
		}
		timeStats.LastSelected = myTimestamp
		valid = true
	}
	// !position
	return
}

