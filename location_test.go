package main

import (
	"testing"
)

var (
	// Complete
	data1 = Polygon{P: []Location{
		{52.3890110622346, 3.2958984375},
		{44.2924010852901, -6.943359375},
		{41.6236553906864, 13.798828125},
		{49.6818468994013, 9.0087890625},
		{52.3890110622346, 3.2958984375},
	}}

	// Incomplete
	data2 = Polygon{P: []Location{
		{51.7406361640977, 0.0439453125},
		{48.2978124924372, -5.625},
		{42.7954006530372, 1.669921875},
		{42.8954006530372, 3.669921875},
		{47.8574028946582, 8.701171875},
	}}
)

func TestCheckComplete(t *testing.T) {
	if !data1.checkComplete() {
		t.Errorf("Error loading polygon %v", data1)
	}

	if data2.checkComplete() {
		t.Errorf("Error loading polygon %v", data2)
	}
}

func TestPolygonLen(t *testing.T) {
	mylen := data1.len()
	if mylen != len(data1.P) {
		t.Errorf("Error loading polygon %v, invalid length %d", data1, mylen)
	}

	mylen = data2.len()
	if mylen != len(data2.P) {
		t.Errorf("Error loading polygon %v, invalid length %d", data2, mylen)
	}
}

func TestPointInPolygon(t *testing.T) {
	good := Location{45,5}
	if !good.pointInPolygon(data1.P) {
		t.Errorf("Point %v should be in polygon %v", good, data1.P)
	}

	bad := Location{50,-5}
	if bad.pointInPolygon(data1.P) {
		t.Errorf("Point %v should not be in polygon %v", bad, data1.P)
	}
}