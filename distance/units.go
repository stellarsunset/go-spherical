package distance

type Unit int

const (
	NauticalMiles Unit = iota
	Feet
	Meters
	Kilometers
	Miles
)

type info struct {
	perMeter float64
	abbr     string
}

var units = [...]info{
	NauticalMiles: {perMeter: 1. / 1852, abbr: "NM"},
	Feet:          {perMeter: 1. / .3048, abbr: "ft"},
	Meters:        {perMeter: 1., abbr: "m"},
	Kilometers:    {perMeter: .001, abbr: "km"},
	Miles:         {perMeter: 1. / (.3048 * 5280.), abbr: "mi"},
}

func UnitsPerMeter(unit Unit) float64 {
	return units[unit].perMeter
}

func Abbr(unit Unit) string {
	return units[unit].abbr
}
