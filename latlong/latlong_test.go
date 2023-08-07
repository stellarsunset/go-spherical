package latlong_test

import (
	"math"
	dist "stellarsunset/spherical/distance"
	ll "stellarsunset/spherical/latlong"
	"testing"
)

func isTrue(t *testing.T, condition bool, s string) {
	if !condition {
		t.Error(s)
	}
}

func isFalse(t *testing.T, condition bool, s string) {
	if condition {
		t.Error(s)
	}
}

func isEqual(t *testing.T, expected, actual any) {
	if expected != actual {
		t.Errorf("want = %+v, got = %+v", expected, actual)
	}
}

func withinError(t *testing.T, expected, actual, tolerance float64) {
	if math.Abs(expected-actual) > tolerance {
		t.Errorf("want = %f, got = %f, tol = %f", expected, actual, tolerance)
	}
}

func TestNewLatLong(t *testing.T) {
	ll := ll.NewLatLong(1., -1.)
	isEqual(t, 1., ll.Latitude())
	isEqual(t, -1., ll.Longitude())
}

func TestDistanceTo(t *testing.T) {

	one, two := ll.NewLatLong(0., 0.), ll.NewLatLong(1., 1.)

	actual := one.DistanceTo(two)
	expected := dist.OfKilometers(157.2)

	withinError(t, expected.InNauticalMiles(), actual.InNauticalMiles(), .1)
}

func TestDistanceToTriangle(t *testing.T) {

	a, b, c := ll.NewLatLong(0., 0.), ll.NewLatLong(0., 1.), ll.NewLatLong(1., 0.)

	l1, l2, hypotenuse := a.DistanceTo(b).InMiles(), a.DistanceTo(c).InMiles(), b.DistanceTo(c).InMiles()
	isTrue(t, math.Pow(hypotenuse, 2) <= math.Pow(l1, 2)+math.Pow(l2, 2), "c^2 <= a^2 + b^2 on a sphere")
}

func TestCourseInDegrees(t *testing.T) {

	one, two := ll.NewLatLong(0., 0.), ll.NewLatLong(1., 1.)

	withinError(t, 45., one.CourseInDegrees(two), .1)
}

func TestIsWithinSamePoint(t *testing.T) {

	one, two := ll.NewLatLong(0., 0.), ll.NewLatLong(0., 0.)

	isTrue(t, one.IsWithin(dist.OfFeet(0), two), "Oft")
	isTrue(t, one.IsWithin(dist.OfFeet(.1), two), ".1ft")
	isFalse(t, one.IsWithin(dist.OfFeet(-1.), two), "-1ft")
}

func TestIsWithinDifferentPoints(t *testing.T) {

	one, two := ll.NewLatLong(0., 0.), ll.NewLatLong(1., 1.)

	isFalse(t, one.IsWithin(dist.OfKilometers(0), two), "Okm")
	isFalse(t, one.IsWithin(dist.OfKilometers(50), two), "50km")
	isTrue(t, one.IsWithin(dist.OfKilometers(160), two), "160km")
}

func TestProjectOut(t *testing.T) {

	source := ll.NewLatLong(0., 0.)

	deg, nm := 45., dist.OfKilometers(157.2).InNauticalMiles()

	actual := source.ProjectOut(deg, nm)
	expected := ll.NewLatLong(1., 1.)

	withinError(t, expected.Latitude(), actual.Latitude(), .01)
	withinError(t, expected.Longitude(), actual.Longitude(), .01)
}

func TestProjectOutNegativeDistance(t *testing.T) {

	source := ll.NewLatLong(0., 0.)

	deg, nm := 45., dist.OfKilometers(157.2).Negate().InNauticalMiles()

	actual := source.ProjectOut(deg, nm)
	expected := ll.NewLatLong(-1., -1.)

	withinError(t, expected.Latitude(), actual.Latitude(), .01)
	withinError(t, expected.Longitude(), actual.Longitude(), .01)
}

func TestCrossTrackDistanceTo(t *testing.T) {

	start, end, point := ll.NewLatLong(0., 0.), ll.NewLatLong(0., 10.), ll.NewLatLong(1., 5.)

	expected := dist.OfNauticalMiles(-60.00686673640662)
	actual := point.CrossTrackDistanceTo(start, end)

	withinError(t, expected.InNauticalMiles(), actual.InNauticalMiles(), .01)
}

func TestAlongTrackDistanceTo(t *testing.T) {

	start, end, point := ll.NewLatLong(0., 0.), ll.NewLatLong(0., 10.), ll.NewLatLong(1., -.5)

	cross := point.CrossTrackDistanceTo(start, end)

	expected := dist.OfNauticalMiles(-30.00343415285915)
	actual := point.AlongTrackDistanceTo(start, end, cross)

	withinError(t, expected.InNauticalMiles(), actual.InNauticalMiles(), .01)
}
