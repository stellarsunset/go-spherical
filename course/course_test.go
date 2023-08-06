package course_test

import (
	"math"
	crs "stellarsunset/spherical/course"
	"testing"
)

const (
	tolerance = 0.0000001
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

func withinError(t *testing.T, expected, actual float64, s string) {
	if math.Abs(expected-actual) > tolerance {
		t.Errorf("%s: want = %f, got = %f, tol = %f", s, expected, actual, tolerance)
	}
}

func withinFractionOfExpected(t *testing.T, expected, actual float64, s string) {
	if math.Abs(expected-actual) > (tolerance * expected) {
		t.Errorf("%s: want = %f, got = %f, tol = %f", s, expected, actual, tolerance)
	}
}

func TestOf(t *testing.T) {

	oneDegree, twoRadians := crs.Of(1, crs.Degrees), crs.Of(2, crs.Radians)

	withinError(t, 1., oneDegree.InDegrees(), "InDegrees(OneDegree)")
	isTrue(t, oneDegree.NativeUnit() == crs.Degrees, "NativeUnit(OneDegree)")

	withinError(t, 2., twoRadians.InRadians(), "InRadians(TwoRadians)")
	isTrue(t, twoRadians.NativeUnit() == crs.Radians, "NativeUnit(TwoRadians)")
}

func TestOfDegrees(t *testing.T) {

	oneDegree := crs.OfDegrees(1)

	withinError(t, 1., oneDegree.InDegrees(), "InDegrees(OneDegree)")
	isTrue(t, oneDegree.NativeUnit() == crs.Degrees, "NativeUnit(OneDegree)")
}

func TestOfRadians(t *testing.T) {

	twoRadians := crs.OfRadians(2)

	withinError(t, 2., twoRadians.InRadians(), "InRadians(TwoRadians)")
	isTrue(t, twoRadians.NativeUnit() == crs.Radians, "NativeUnit(TwoRadians)")
}

func TestUnitConversions(t *testing.T) {

	ninetyDegrees := crs.OfDegrees(90.)
	withinError(t, math.Pi/2., ninetyDegrees.InRadians(), "InRadians(90)")

	threeSixtyDegrees := crs.OfDegrees(360)
	withinError(t, math.Pi*2, threeSixtyDegrees.InRadians(), "InRadians(360)")

	zeroDegrees := crs.OfDegrees(0)
	withinError(t, 0., zeroDegrees.InRadians(), "InRadians(0)")

	twoRadians := crs.OfRadians(2)
	withinError(t, 2*(180/math.Pi), twoRadians.InDegrees(), "InDegrees(2rad)")

	piRadians := crs.OfRadians(math.Pi)
	withinError(t, 180., piRadians.InDegrees(), "InDegrees(Pirad)")

	zeroRadians := crs.OfRadians(0)
	withinError(t, 0, zeroRadians.InDegrees(), "InDegrees(0rad)")
}

func TestUnitConversionsNegatives(t *testing.T) {

	ninetyDegrees := crs.OfDegrees(-90.)
	withinError(t, -math.Pi/2., ninetyDegrees.InRadians(), "InRadians(90)")

	threeSixtyDegrees := crs.OfDegrees(-360)
	withinError(t, -math.Pi*2, threeSixtyDegrees.InRadians(), "InRadians(360)")

	zeroDegrees := crs.OfDegrees(0)
	withinError(t, 0., zeroDegrees.InRadians(), "InRadians(0)")

	twoRadians := crs.OfRadians(-2)
	withinError(t, -2*(180/math.Pi), twoRadians.InDegrees(), "InDegrees(2rad)")

	piRadians := crs.OfRadians(-math.Pi)
	withinError(t, -180., piRadians.InDegrees(), "InDegrees(Pirad)")

	zeroRadians := crs.OfRadians(0)
	withinError(t, 0, zeroRadians.InDegrees(), "InDegrees(0rad)")
}

func TestNegate(t *testing.T) {
	ninetyDegrees, negativeNinetyDegrees := crs.OfDegrees(90), crs.OfDegrees(-90)
	isEqual(t, *ninetyDegrees, *negativeNinetyDegrees.Negate())
	isEqual(t, *negativeNinetyDegrees, *ninetyDegrees.Negate())
}

func TestAbs(t *testing.T) {
	ninetyDegrees, negativeNinetyDegrees := crs.OfDegrees(90), crs.OfDegrees(-90)
	isEqual(t, *ninetyDegrees, *ninetyDegrees.Abs())
	isEqual(t, *ninetyDegrees, *negativeNinetyDegrees.Abs())
}

func TestIsPositive(t *testing.T) {
	isTrue(t, crs.OfDegrees(20).IsPositive(), "IsPositive(20)")
	isFalse(t, crs.OfDegrees(0).IsPositive(), "IsPositive(0)")
	isFalse(t, crs.OfDegrees(-20).IsPositive(), "IsPositive(-20)")
}

func TestIsNegative(t *testing.T) {
	isFalse(t, crs.OfDegrees(20).IsNegative(), "IsNegative(20)")
	isFalse(t, crs.OfDegrees(0).IsNegative(), "IsNegative(0)")
	isTrue(t, crs.OfDegrees(-20).IsNegative(), "IsNegative(-20)")
}

func TestIsZero(t *testing.T) {
	isFalse(t, crs.OfDegrees(20).IsZero(), "IsZero(20)")
	isTrue(t, crs.OfDegrees(0).IsZero(), "IsZero(0)")
	isFalse(t, crs.OfDegrees(-20).IsZero(), "IsZero(-20)")
}

func TestTimes(t *testing.T) {

}

func TestPlus(t *testing.T) {

}

func TestMinus(t *testing.T) {

}

func TestComparisonMethods(t *testing.T) {

}
