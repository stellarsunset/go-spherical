package course

import "math"

type Unit int

const (
	Degrees Unit = iota
	Radians
)

type info struct {
	perDegree float64
	abbr      string
}

var units = [...]info{
	Degrees: {perDegree: 1., abbr: "deg"},
	Radians: {perDegree: math.Pi / 180., abbr: "rad"},
}

func UnitsPerDegree(unit Unit) float64 {
	return units[unit].perDegree
}

func Abbr(unit Unit) string {
	return units[unit].abbr
}
