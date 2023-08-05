package spherical

import (
	"math"
	"testing"
)

func almostEquals(expected, actual, maxError float64) bool {
	return math.Abs(expected-actual) <= maxError
}

func withinError(t *testing.T, expected, actual, maxError float64, s string) {
	if math.Abs(expected-actual) > maxError {
		t.Errorf("%s: want = %f, got = %f, tol = %f", s, expected, actual, maxError)
	}
}

func TestDistanceInNm(t *testing.T) {

	// http://www.movable-type.co.uk/scripts/latlong.html
	expectedDistanceKm := 1569.
	kmPerNm := 1.852
	withinError(t, expectedDistanceKm/kmPerNm, DistanceInNm(0., 0., 10., 10.), 1., "DistanceInNm()")
}

func TestCourseInDegreesDueNorth(t *testing.T) {
	withinError(t, 360., CourseInDegrees(0., 0., 10., 0.), .01, "CourseInDegrees(North)")
}

func TestCourseInDegreesDueEast(t *testing.T) {
	withinError(t, 90., CourseInDegrees(0., 0., 0., 10.), .01, "CourseInDegrees(East)")
}

func TestCourseInDegreesDueWest(t *testing.T) {
	withinError(t, 270., CourseInDegrees(0., 10., 0., 0.), .01, "CourseInDegrees(West)")
}

func TestCourseInDegreesDueSouth(t *testing.T) {
	withinError(t, 180., CourseInDegrees(10., 0., 0., 0.), .01, "CourseInDegrees(South)")
}

func TestCourseInDegreesNorthEast(t *testing.T) {
	withinError(t, 45., CourseInDegrees(0., 0., 1., 1.), .01, "CourseInDegrees(NorthEast)")
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
