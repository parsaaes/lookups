package lookups

import "github.com/golang/geo/s2"

type GeoIndex interface {
	Find(Coordinate) string
	Cover(*s2.Polygon) []string
}

// S2Index is a S2-based polygon indexer.
type S2Index struct {
	level int
}

// NewS2Index returns a new S2 index instance.
func NewS2Index(level int) S2Index {
	return S2Index{level: level}
}

// Find return an S2 hash that contains the given point.
func (s S2Index) Find(coordinate Coordinate) string {
	ll := s2.LatLngFromDegrees(coordinate.Latitude, coordinate.Longitude)

	return s2.CellIDFromLatLng(ll).Parent(s.level).ToToken()
}

// Covers returns all S2 cells that cover the given polygon.
func (s S2Index) Cover(polygon *s2.Polygon) []string {
	if polygon.NumEdges() < 1 {
		return []string{}
	}

	ids := s2.SimpleRegionCovering(polygon, polygon.Edge(0).V0, s.level)

	out := make([]string, 0, len(ids))

	for _, id := range ids {
		out = append(out, id.ToToken())
	}

	return out
}
