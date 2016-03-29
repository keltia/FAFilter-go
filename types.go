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

type Location struct {
	Latitude  float64
	Longitude float64
}

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

type TimeStats struct {
	FirstSeen		time.Time
	LastSeen		time.Time
	FirstSelected	time.Time
	LastSelected	time.Time
	Lowest	time.Time
	Highest	time.Time
}

type RecordStats struct {
	TotalRead			int
	TotalSkipped		int
	SkippedTemporal		int
	SkippedGeometric	int
	SkippedAircraftID	int
	SkippedHexid		int
	SkippedUpdateType	int
}

type Polygon struct {
	P		[]Location
}
