package distance_test

import (
	"stellarsunset/spherical/distance"
	"testing"
)

func TestMetersPerUnit(t *testing.T) {

	want := 1.
	if got := distance.UnitsPerMeter(distance.Meters); got != want {
		t.Errorf("UnitsPerMeter(Meter) = %f, want %f", got, want)
	}
}

func TestAbbreviation(t *testing.T) {

	want := distance.NauticalMiles
	if got := distance.OfNauticalMiles(1.).NativeUnit(); got != want {
		t.Errorf("OfNauticalMiles(1.) = %q, want %q", distance.Abbr(got), distance.Abbr(want))
	}
}
