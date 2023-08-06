package course_test

import (
	"math"
	crs "stellarsunset/spherical/course"
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

func isEqual(t *testing.T, expected, actual crs.Course) {
	if expected != actual {
		t.Errorf("want = %+v, got = %+v", expected, actual)
	}
}

func withinError(t *testing.T, expected, actual, maxError float64, s string) {
	if math.Abs(expected-actual) > maxError {
		t.Errorf("%s: want = %f, got = %f, tol = %f", s, expected, actual, maxError)
	}
}

func withinFractionOfExpected(t *testing.T, expected, actual, percentError float64, s string) {
	if math.Abs(expected-actual) > (percentError * expected) {
		t.Errorf("%s: want = %f, got = %f, tol = %f", s, expected, actual, percentError)
	}
}
