package course_test

import (
	crs "stellarsunset/spherical/course"
	"testing"
)

func TestDegreesPerUnit(t *testing.T) {

	want := 1.
	if got := crs.UnitsPerDegree(crs.Degrees); got != want {
		t.Errorf("UnitsPerDegree(Meter) = %f, want %f", got, want)
	}
}

func TestAbbreviation(t *testing.T) {

	want := crs.Degrees
	if got := crs.OfDegrees(1.).NativeUnit(); got != want {
		t.Errorf("OfDegrees(1.) = %q, want %q", crs.Abbr(got), crs.Abbr(want))
	}
}
