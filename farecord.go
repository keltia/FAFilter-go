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
func (line *FArecord) checkRecord() bool {
	var myTimestamp time.Time

	if line.Type == "position" {
		// Record timestamps
		if line.Clock != "" {
			var value int64

			if _, err := fmt.Sscanf(line.Clock, "%d", &value); err != nil {
				fmt.Fprintf(os.Stderr, "Invalid value %d\n", value)
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
		if fAircraftId != "" {
			cont, _ := regexp.MatchString(fAircraftId, line.Ident)
			if cont == false {
				recordStats.SkippedAircraftId++
				return cont
			}
		}

		// Check for -x
		if fHexid != "" {
			cont, _ := regexp.MatchString(fHexid, line.Hexid)
			if cont == false {
				recordStats.SkippedHexid++
				return cont
			}
		}

		// Check for -t
		if fUpdateType != "" {
			cont, _ := regexp.MatchString(fUpdateType, line.UpdateType)
			if cont == false {
				recordStats.SkippedUpdateType++
				return cont
			}
		}

		// Check for -g
		var found	bool = false

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
						found = true
						break
					}
				}
			}
			if !found {
				recordStats.SkippedGeometric++
				return false
			}
		}

		// Everything has been checked
		// fill in stats again
		if (timeStats.FirstSelected == time.Time{}) {
			timeStats.FirstSelected = myTimestamp
		}
		timeStats.LastSelected = myTimestamp

		return true
	} else {
		return false
	}// position

	// fallthrough
	return true
}

