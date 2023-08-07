package spherical_test

import (
	"math"
	sph "stellarsunset/spherical"
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
	withinError(t, expectedDistanceKm/kmPerNm, sph.DistanceInNm(0., 0., 10., 10.), 1., "DistanceInNm()")
}

func TestCourseInDegreesDueNorth(t *testing.T) {
	withinError(t, 360., sph.CourseInDegrees(0., 0., 10., 0.), .01, "CourseInDegrees(North)")
}

func TestCourseInDegreesDueEast(t *testing.T) {
	withinError(t, 90., sph.CourseInDegrees(0., 0., 0., 10.), .01, "CourseInDegrees(East)")
}

func TestCourseInDegreesDueWest(t *testing.T) {
	withinError(t, 270., sph.CourseInDegrees(0., 10., 0., 0.), .01, "CourseInDegrees(West)")
}

func TestCourseInDegreesDueSouth(t *testing.T) {
	withinError(t, 180., sph.CourseInDegrees(10., 0., 0., 0.), .01, "CourseInDegrees(South)")
}

func TestCourseInDegreesNorthEast(t *testing.T) {
	withinError(t, 45., sph.CourseInDegrees(0., 0., 1., 1.), .01, "CourseInDegrees(NorthEast)")
}

func TestProjectOut(t *testing.T) {

	startLat, startLon := 0., 0.
	expectedLat, expectedLon := 10., 10.

	course := sph.CourseInDegrees(startLat, startLon, expectedLat, expectedLon)
	distance := sph.DistanceInNm(startLat, startLon, expectedLat, expectedLon)

	actualLat, actualLon := sph.ProjectOut(startLat, startLon, course, distance)

	withinError(t, expectedLat, actualLat, .01, "Latitude")
	withinError(t, expectedLon, actualLon, .01, "Longitude")
}

func TestProjectOutPositiveDistance(t *testing.T) {

	startLat, startLon := 0., 0.
	expectedLat, expectedLon := 1., 1.

	// approx course/distance from (0,0)->(1,1)
	crs, dist := 45., 84.860371
	actualLat, actualLon := sph.ProjectOut(startLat, startLon, crs, dist)

	withinError(t, expectedLat, actualLat, .01, "Latitude")
	withinError(t, expectedLon, actualLon, .01, "Longitude")
}

func TestProjectOutNegativeDistance(t *testing.T) {

	startLat, startLon := 0., 0.
	expectedLat, expectedLon := -1., -1.

	// approx course/distance from (0,0)->(1,1)
	crs, dist := 45., 84.860371
	actualLat, actualLon := sph.ProjectOut(startLat, startLon, crs, -dist)

	withinError(t, expectedLat, actualLat, .01, "Latitude")
	withinError(t, expectedLon, actualLon, .01, "Longitude")
}

func TestCrossTrackDistanceLeft(t *testing.T) {

	expected := -60.00686673640662
	actual := sph.CrossTrackDistanceNm(0., 0., 0., 10., 1., .5)

	if !almostEquals(expected, actual, .0001) {
		t.Errorf("CrossTrackDistance(...) = %f, wanted %f", actual, expected)
	}
}

func TestCrossTrackDistanceRight(t *testing.T) {

	expected := 60.00686673640662
	actual := sph.CrossTrackDistanceNm(0., 0., 0., 10., -1., .5)

	if !almostEquals(expected, actual, .0001) {
		t.Errorf("CrossTrackDistance(...) = %f, wanted %f", actual, expected)
	}
}

func TestAlongTrackDistancePositive(t *testing.T) {
	tolerance, cross := .0001, sph.CrossTrackDistanceNm(0., 0., 0., 10., 1., .5)

	expected := 30.00343415285915
	actual := sph.AlongTrackDistanceNm(0., 0., 0., 10., 1., .5, cross)

	if !almostEquals(expected, actual, tolerance) {
		t.Errorf("AlongTrackDistanceNm(...) = %f, wanted %f", actual, expected)
	}
}

func TestAlongTrackDistanceNegative(t *testing.T) {
	tolerance := .0001

	cross := sph.CrossTrackDistanceNm(0., 0., 0., 10., 1., -.5)

	expected := -30.00343415285915
	actual := sph.AlongTrackDistanceNm(0., 0., 0., 10., 1., -.5, cross)

	if !almostEquals(expected, actual, tolerance) {
		t.Errorf("AlongTrackDistanceNm(...) = %f, wanted %f", actual, expected)
	}
}
