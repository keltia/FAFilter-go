// location.go
//
// Location related methods/functions

package main

// Implement XOR for bool
func xor(a, b bool) bool {
	return a != b
}

// Code rewritten in Go from http://alienryderflex.com/polygon/
//
//  The function will return true if the point x,y is inside the polygon, or
//  false if it is not.  If the point is exactly on the edge of the polygon,
//  then the function may return true or false.
//
//  Note that division by zero is avoided because the division is protected
//  by the "if" clause which surrounds it.
func (loc *Location) pointInPolygon(zone []Location) bool {
	if loc == nil {
		return false
	}

	var i int

	polySides := len(zone)
	j := polySides - 1
	oddNodes := false

	for i= 0; i < polySides; i ++ {
		if (zone[i].Latitude < loc.Latitude && zone[j].Latitude >= loc.Latitude ||
			zone[j].Latitude < loc.Latitude && zone[i].Latitude >= loc.Latitude) &&
			(zone[i].Longitude <= loc.Longitude || zone[j].Longitude <= loc.Longitude) {
			oddNodes = xor(oddNodes, (zone[i].Longitude + (loc.Latitude - zone[i].Latitude) / (zone[j].Latitude - zone[i].Latitude) * (zone[j].Longitude - zone[i].Longitude) < loc.Longitude))
		}
	}
	return oddNodes
}
