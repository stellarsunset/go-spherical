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

func isEqual(t *testing.T, expected, actual any) {
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
	isEqual(t, crs.Degrees, oneDegree.NativeUnit())

	withinError(t, 2., twoRadians.InRadians(), "InRadians(TwoRadians)")
	isEqual(t, crs.Radians, twoRadians.NativeUnit())
}

func TestOfDegrees(t *testing.T) {

	oneDegree := crs.OfDegrees(1)

	withinError(t, 1., oneDegree.InDegrees(), "InDegrees(OneDegree)")
	isEqual(t, crs.Degrees, oneDegree.NativeUnit())
}

func TestOfRadians(t *testing.T) {

	twoRadians := crs.OfRadians(2)

	withinError(t, 2., twoRadians.InRadians(), "InRadians(TwoRadians)")
	isEqual(t, crs.Radians, twoRadians.NativeUnit())
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
	oneDegree, fourDegrees := crs.OfDegrees(1), crs.OfDegrees(4)
	isEqual(t, *fourDegrees, *oneDegree.Times(4))
}

func TestPlus(t *testing.T) {
	oneRadian, ninetyDegrees := crs.OfRadians(1), crs.OfDegrees(90)

	isEqual(t, crs.Radians, oneRadian.Plus(ninetyDegrees).NativeUnit())
	isEqual(t, crs.Degrees, ninetyDegrees.Plus(oneRadian).NativeUnit())

	isEqual(t, 1.0+math.Pi/2., oneRadian.Plus(ninetyDegrees).InRadians())
	isEqual(t, 1.0+math.Pi/2., ninetyDegrees.Plus(oneRadian).InRadians())
}

func TestMinus(t *testing.T) {
	oneRadian, ninetyDegrees := crs.OfRadians(1), crs.OfDegrees(90)

	isEqual(t, crs.Radians, oneRadian.Minus(ninetyDegrees).NativeUnit())
	isEqual(t, crs.Degrees, ninetyDegrees.Minus(oneRadian).NativeUnit())

	withinError(t, 1.0-(math.Pi/2.), oneRadian.Minus(ninetyDegrees).InRadians(), "Rad(1) - Deg(90)")
	withinError(t, (math.Pi/2.)-1., ninetyDegrees.Minus(oneRadian).InRadians(), "Deg(90) - Rad(1)")
}

func TestComparisonMethods(t *testing.T) {

	zeroRadians, zeroDegrees, oneDegree := crs.OfRadians(0), crs.OfDegrees(0), crs.OfDegrees(1)

	isTrue(t, zeroRadians.IsGreaterThanOrEqualTo(zeroDegrees), "Rad(0) >= Deg(0)")
	isTrue(t, zeroDegrees.IsGreaterThanOrEqualTo(zeroRadians), "Deg(0) >= Rad(0)")
	isFalse(t, zeroRadians.IsGreaterThan(zeroDegrees), "Rad(0) > Deg(0)")
	isFalse(t, zeroDegrees.IsGreaterThan(zeroRadians), "Deg(0) > Rad(0)")

	isTrue(t, zeroRadians.IsLessThanOrEqualTo(zeroDegrees), "Rad(0) <= Deg(0)")
	isTrue(t, zeroDegrees.IsLessThanOrEqualTo(zeroRadians), "Deg(0) <= Rad(0)")
	isFalse(t, zeroRadians.IsLessThan(zeroDegrees), "Rad(0) < Deg(0)")
	isFalse(t, zeroDegrees.IsLessThan(zeroRadians), "Deg(0) < Rad(0)")

	isTrue(t, oneDegree.IsGreaterThanOrEqualTo(zeroDegrees), "Deg(1) >= Deg(0)")
	isTrue(t, oneDegree.IsGreaterThan(zeroDegrees), "Deg(1) > Deg(0)")
	isFalse(t, oneDegree.IsLessThanOrEqualTo(zeroDegrees), "Deg(1) <= Deg(0)")
	isFalse(t, oneDegree.IsLessThan(zeroDegrees), "Deg(1) < Deg(0)")
}

func TestAngleBetween(t *testing.T) {

	diff := crs.AngleBetween(crs.OfDegrees(5), crs.OfDegrees(355))
	isEqual(t, *crs.OfDegrees(10), *diff)

	diff = crs.AngleBetween(crs.OfDegrees(355), crs.OfDegrees(5))
	isEqual(t, *crs.OfDegrees(-10), *diff)
}

func TestAngleDifference(t *testing.T) {

	// positive
	isEqual(t, 5., crs.AngleDifference(5., 0.))
	isEqual(t, 175., crs.AngleDifference(175., 0))
	isEqual(t, -175., crs.AngleDifference(185., 0))
	isEqual(t, -5., crs.AngleDifference(355., 0.))

	// negative
	isEqual(t, -5., crs.AngleDifference(-5., 0))
	isEqual(t, -175., crs.AngleDifference(-175., 0))
	isEqual(t, 175., crs.AngleDifference(-185., 0))
}

func TestSin(t *testing.T) {
	withinError(t, -1., crs.OfDegrees(-90).Sin(), "Sin(-90)")
	withinError(t, -math.Sqrt(2.)/2., crs.OfDegrees(-45).Sin(), "Sin(-45)")
	withinError(t, 0., crs.OfDegrees(0).Sin(), "Sin(0)")
	withinError(t, math.Sqrt(2.)/2., crs.OfDegrees(45).Sin(), "Sin(45)")
	withinError(t, 1., crs.OfDegrees(90).Sin(), "Sin(90)")
}

func TestCos(t *testing.T) {
	withinError(t, 0., crs.OfDegrees(-90).Cos(), "Cos(-90)")
	withinError(t, math.Sqrt(2.)/2, crs.OfDegrees(-45).Cos(), "Cos(-45)")
	withinError(t, 1., crs.OfDegrees(0).Cos(), "Cos(0)")

	withinError(t, 0., crs.OfDegrees(90).Cos(), "Cos(90)")
	withinError(t, -math.Sqrt(2.)/2, crs.OfDegrees(135).Cos(), "Cos(135)")
	withinError(t, -1., crs.OfDegrees(180).Cos(), "Cos(180)")
}

func TestTan(t *testing.T) {
	withinError(t, -1., crs.OfDegrees(-45).Tan(), "Tan(-45)")
	withinError(t, 0., crs.OfDegrees(0).Tan(), "Tan(0)")
	withinError(t, 1., crs.OfDegrees(45).Tan(), "Tan(45)")
	isTrue(t, 2e14 < crs.OfDegrees(90.).Tan(), "Tan(90)")
	withinError(t, -1., crs.OfDegrees(135).Tan(), "Tan(135)")
}
