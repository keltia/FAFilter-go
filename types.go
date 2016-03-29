// types.go
//
// All new types needed for FAFilter
//
// Copyright 2015 © by Ollivier Robert for the EEC
//

package main

import (
	"time"
)

// Location is the generic position
type Location struct {
	Latitude  float64
	Longitude float64
}

// FArecord is the main type we are playing with
type FArecord struct {
	Type		string
	AirGround	string
	AltChange	string
	Clock		string
	Gs			string
	Heading		string
	Hexid		string
	ID			string
	Ident		string
	Lat			string
	Lon			string
	Reg			string
	Squawk		string
	UpdateType	string
}

// TimeStats is for recording time statistics for records
type TimeStats struct {
	FirstSeen		time.Time
	LastSeen		time.Time
	FirstSelected	time.Time
	LastSelected	time.Time
	Lowest	time.Time
	Highest	time.Time
}

// RecordStats is general statistics for our filtering
type RecordStats struct {
	TotalRead			int
	TotalSkipped		int
	SkippedTemporal		int
	SkippedGeometric	int
	SkippedAircraftID	int
	SkippedHexid		int
	SkippedUpdateType	int
}

// Polygon is a series of — hopefully — closed points
type Polygon struct {
	P		[]Location
}
