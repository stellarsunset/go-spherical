package spherical

import (
	"math"
	"testing"
)

func almostEquals(expected, actual, maxError float64) bool {
	return math.Abs(expected-actual) <= maxError
}

func TestDistanceInNm(t *testing.T) {

	// http://www.movable-type.co.uk/scripts/latlong.html
	expectedDistanceKm := 1569.
	kmPerNm := 1.852

	want := expectedDistanceKm / kmPerNm
	if got := DistanceInNm(0., 0., 10., 10.); !almostEquals(got, want, 1.) {
		t.Errorf("DistanceInNm(...) = %f, want %f", got, want)
	}
}

func TestCourseInDegreesDueNorth(t *testing.T) {

	want := 360.
	if got := CourseInDegrees(0., 0., 10., 0.); !almostEquals(got, want, .01) {
		t.Errorf("CourseInDegrees(...) = %f, want %f", got, want)
	}
}

func TestCourseInDegreesDueEast(t *testing.T) {

	want := 90.
	if got := CourseInDegrees(0., 0., 0., 10.); !almostEquals(got, want, .01) {
		t.Errorf("CourseInDegrees(...) = %f, want %f", got, want)
	}
}

func TestCourseInDegreesDueWest(t *testing.T) {

	want := 270.
	if got := CourseInDegrees(0., 10., 0., 0.); !almostEquals(got, want, .01) {
		t.Errorf("CourseInDegrees(...) = %f, want %f", got, want)
	}
}

func TestCourseInDegreesDueSouth(t *testing.T) {

	want := 180.
	if got := CourseInDegrees(10., 0., 0., 0.); !almostEquals(got, want, .01) {
		t.Errorf("CourseInDegrees(...) = %f, want %f", got, want)
	}
}

func TestCourseInDegreesNorthEast(t *testing.T) {

	want := 45.
	if got := CourseInDegrees(0., 0., 1., 1.); !almostEquals(got, want, .01) {
		t.Errorf("CourseInDegrees(...) = %f, want %f", got, want)
	}
}

func TestProjectOut(t *testing.T) {

	startLat, startLon := 0., 0.
	expectedLat, expectedLon := 10., 10.

	course := CourseInDegrees(startLat, startLon, expectedLat, expectedLon)
	distance := DistanceInNm(startLat, startLon, expectedLat, expectedLon)

	if actualLat, actualLon := ProjectOut(startLat, startLon, course, distance); !almostEquals(expectedLat, actualLat, .01) || !almostEquals(expectedLon, actualLon, .01) {
		t.Errorf("ProjectOut(...) = (%f, %f), wanted (%f, %f)", actualLat, actualLon, expectedLat, expectedLon)
	}
}

func TestCrossTrackDistanceLeft(t *testing.T) {

	expected := -60.00686673640662
	actual := CrossTrackDistanceNm(0., 0., 0., 10., 1., .5)

	if !almostEquals(expected, actual, .0001) {
		t.Errorf("CrossTrackDistance(...) = %f, wanted %f", actual, expected)
	}
}

func TestCrossTrackDistanceRight(t *testing.T) {

	expected := 60.00686673640662
	actual := CrossTrackDistanceNm(0., 0., 0., 10., -1., .5)

	if !almostEquals(expected, actual, .0001) {
		t.Errorf("CrossTrackDistance(...) = %f, wanted %f", actual, expected)
	}
}

func TestAngleDifferencePositiveInputs(t *testing.T) {
	tolerance := .0001

	expected, actual := 5., angleDifference(5.)
	if !almostEquals(actual, expected, tolerance) {
		t.Errorf("AngleDifference(...) = %f, wanted %f", actual, expected)
	}

	expected, actual = 175., angleDifference(175.)
	if !almostEquals(actual, expected, tolerance) {
		t.Errorf("AngleDifference(...) = %f, wanted %f", actual, expected)
	}

	expected, actual = -175., angleDifference(185.)
	if !almostEquals(actual, expected, tolerance) {
		t.Errorf("AngleDifference(...) = %f, wanted %f", actual, expected)
	}

	expected, actual = -5., angleDifference(355.)
	if !almostEquals(actual, expected, tolerance) {
		t.Errorf("AngleDifference(...) = %f, wanted %f", actual, expected)
	}
}

func TestAngleDifferenceNegativeInputs(t *testing.T) {
	tolerance := .0001

	expected, actual := -5., angleDifference(-5.)
	if !almostEquals(actual, expected, tolerance) {
		t.Errorf("AngleDifference(...) = %f, wanted %f", actual, expected)
	}

	expected, actual = -175., angleDifference(-175.)
	if !almostEquals(actual, expected, tolerance) {
		t.Errorf("AngleDifference(...) = %f, wanted %f", actual, expected)
	}

	expected, actual = 175., angleDifference(-185.)
	if !almostEquals(actual, expected, tolerance) {
		t.Errorf("AngleDifference(...) = %f, wanted %f", actual, expected)
	}
}

func TestAngleDifferencePublic(t *testing.T) {
	tolerance := .0001

	expected, actual := 10., AngleDifference(5., 355.)
	if !almostEquals(expected, actual, tolerance) {
		t.Errorf("AngDifference(...) = %f, wanted %f", actual, expected)
	}

	expected, actual = -10., AngleDifference(355., 5.)
	if !almostEquals(expected, actual, tolerance) {
		t.Errorf("AngDifference(...) = %f, wanted %f", actual, expected)
	}
}

func TestAlongTrackDistancePositive(t *testing.T) {
	tolerance := .0001

	expected := 30.00343415285915
	actual := alongTrackDistanceNm(0., 0., 0., 10., 1., .5)

	if !almostEquals(expected, actual, tolerance) {
		t.Errorf("AlongTrackDistanceNm(...) = %f, wanted %f", actual, expected)
	}
}

func TestAlongTrackDistanceNegative(t *testing.T) {
	tolerance := .0001

	expected := -30.00343415285915
	actual := alongTrackDistanceNm(0., 0., 0., 10., 1., -.5)

	if !almostEquals(expected, actual, tolerance) {
		t.Errorf("AlongTrackDistanceNm(...) = %f, wanted %f", actual, expected)
	}
}
