package main

import (
	"time"
	"fmt"
	"regexp"
)

// Generic regex-based check
func checkRegex(value string, regex regexp.Regexp) bool {
	return regex.Match([]byte(value))
}

// Implements main checks
//
// Having several parameters specified on the CLI means AND, not OR because
// we only break on false matches. As long as we match, we keep on.
func (line *FArecord) checkRecord() bool {
	var cont		bool
	var myTimestamp time.Time

	// Check for -a
	if fAircraftId != "" {
		cont = checkRegex(line.Ident, rAircraftId)
		if cont == false {
			recordStats.SkippedAircraftId++
			return cont
		}
	}

	// Check for -x
	if fHexid != "" {
		cont = checkRegex(line.Hexid, rHexid)
		if cont == false {
			recordStats.SkippedHexid++
			return cont
		}
	}

	// Check for -t
	if fUpdateType != "" {
		if line.UpdateType != fUpdateType {
			recordStats.SkippedUpdateType++
			return false
		}
	}
	
	// fallthrough
	return true

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
	return true
}

