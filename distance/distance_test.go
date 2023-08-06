package distance_test

import (
	"math"
	"stellarsunset/spherical/distance"
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

func isEqual(t *testing.T, expected, actual distance.Distance) {
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

func TestOf(t *testing.T) {

	tol, d := .00001, distance.Of(1., distance.NauticalMiles)

	withinError(t, 1., d.InNauticalMiles(), tol, "InNauticalMiles()")
	withinFractionOfExpected(t, 1852., d.InMeters(), tol, "InMeters()")
	withinFractionOfExpected(t, 6076.12, d.InFeet(), tol, "InFeet()")
	withinFractionOfExpected(t, 1.15078, d.InMiles(), tol, "InMiles()")
}

func TestIn(t *testing.T) {

	tol := .00001

	oneNm := distance.OfNauticalMiles(1.)
	withinError(t, 1., oneNm.In(distance.NauticalMiles), tol, "In(NauticalMiles)")
	withinFractionOfExpected(t, 1852., oneNm.In(distance.Meters), tol, "In(Meters)")
	withinFractionOfExpected(t, 6076.12, oneNm.In(distance.Feet), tol, "In(Feet)")
	withinFractionOfExpected(t, 1.15078, oneNm.In(distance.Miles), tol, "In(Miles)")

	oneMeter := distance.OfMeters(1.)
	withinFractionOfExpected(t, 0.000539956803456, oneMeter.In(distance.NauticalMiles), tol, "In(NauticalMiles)")
	withinError(t, 1., oneMeter.In(distance.Meters), tol, "In(Meters)")
	withinError(t, .001, oneMeter.In(distance.Kilometers), tol, "In(Kilometers)")
	withinFractionOfExpected(t, 3.28084, oneMeter.In(distance.Feet), tol, "In(Feet)")
	withinFractionOfExpected(t, .000621371, oneMeter.In(distance.Miles), tol, "In(Miles)")
}

func TestComparisonMethods(t *testing.T) {

	halfMeter, oneMeter := distance.OfMeters(.5), distance.OfMeters(1.)
	oneThousandMeters, oneKilometer := distance.OfMeters(1000.), distance.OfKilometers(1.)

	isTrue(t, halfMeter.IsLessThan(oneMeter), ".5M < 1M")
	isTrue(t, halfMeter.IsLessThanOrEqualTo(oneMeter), ".5M <= 1M")
	isTrue(t, oneMeter.IsGreaterThan(halfMeter), "1M > .5M")
	isTrue(t, oneMeter.IsGreaterThanOrEqualTo(halfMeter), "1M >= .5M")

	isTrue(t, oneKilometer.IsGreaterThanOrEqualTo(oneThousandMeters), "1KM >= 1000M")
	isTrue(t, oneKilometer.IsLessThanOrEqualTo(oneThousandMeters), "1KM <= 1000M")
}

func TestMiles(t *testing.T) {

	tol, oneMile := .00001, distance.OfMiles(1.)
	withinError(t, 1., oneMile.InMiles(), tol, "OneMile.inMiles()")
	withinFractionOfExpected(t, 5280., oneMile.InFeet(), tol, "OneMile.inFeet()")
}

func TestNegate(t *testing.T) {

	oneMeter, negativeMeter := distance.OfMeters(1.), distance.OfMeters(-1.)
	isTrue(t, *negativeMeter == *oneMeter.Negate(), "Abs(-1M) == 1M")
}

func TestIsPositive(t *testing.T) {

	negativeOne, zero, one := distance.OfFeet(-1), distance.OfFeet(0.), distance.OfFeet(1)

	isFalse(t, negativeOne.IsPositive(), "IsPositive(-1)")
	isTrue(t, negativeOne.Negate().IsPositive(), "IsPositive(Negate(-1))")

	isFalse(t, zero.IsPositive(), "IsPositive(0)")
	isFalse(t, zero.Negate().IsPositive(), "IsPositive(Negate(0))")

	isTrue(t, one.IsPositive(), "IsPositive(1)")
	isFalse(t, one.Negate().IsPositive(), "IsPositive(Negate(1))")
}

func TestIsNegative(t *testing.T) {

	negativeOne, zero, one := distance.OfFeet(-1), distance.OfFeet(0.), distance.OfFeet(1)

	isTrue(t, negativeOne.IsNegative(), "IsNegative(-1)")
	isFalse(t, negativeOne.Negate().IsNegative(), "IsNegative(Negate(-1))")

	isFalse(t, zero.IsNegative(), "IsNegative(0)")
	isFalse(t, zero.Negate().IsNegative(), "IsNegative(Negate(0))")

	isFalse(t, one.IsNegative(), "IsNegative(1)")
	isTrue(t, one.Negate().IsNegative(), "IsNegative(Negate(1))")
}

func TestIsZero(t *testing.T) {

	negativeOne, zero, one := distance.OfFeet(-1), distance.OfFeet(0.), distance.OfFeet(1)

	isFalse(t, negativeOne.IsZero(), "IsZero(-1)")
	isFalse(t, negativeOne.Negate().IsZero(), "IsZero(Negate(-1))")

	isTrue(t, zero.IsZero(), "IsZero(0)")
	isTrue(t, zero.Negate().IsZero(), "IsZero(Negate(0))")

	isFalse(t, one.IsZero(), "IsZero(1)")
	isFalse(t, one.Negate().IsZero(), "IsZero(Negate(1))")
}

func TestAbs(t *testing.T) {

	oneMeter, negativeMeter := distance.OfMeters(1.), distance.OfMeters(-1.)

	isTrue(t, *negativeMeter.Abs() == *oneMeter, "Abs(-1M) == 1M")
	isTrue(t, *oneMeter.Abs() == *distance.OfMeters(1.), "Abs(1M) == 1M")
}

func TestTimes(t *testing.T) {

	tol, oneMeter, halfMeter := .00001, distance.OfMeters(1.), distance.OfMeters(.5)
	withinError(t, halfMeter.InMeters(), oneMeter.Times(.5).InMeters(), tol, "1M * .5")
}

func TestPlus(t *testing.T) {

	tol, oneFoot, fiveHalvesFeet := .00001, distance.OfFeet(1.), distance.OfFeet(2.5)

	sum := oneFoot.Plus(fiveHalvesFeet)
	withinError(t, 3.5, sum.InFeet(), tol, "1Ft + 2.5Ft")
}

func TestMinus(t *testing.T) {

	tol, oneFoot, fiveHalvesFeet := .00001, distance.OfFeet(1.), distance.OfFeet(2.5)

	sum := oneFoot.Minus(fiveHalvesFeet)

	withinError(t, -1.5, sum.InFeet(), tol, "1Ft - 2.5Ft")
	isTrue(t, sum.IsNegative(), "sum.IsNegative()")
}

func TestSort(t *testing.T) {

	oneMeter, zero, negativeOneFeet, oneFoot := distance.OfMeters(1), distance.Zero(), distance.OfFeet(-1), distance.OfFeet(1)
	oneNm, fourFeet, oneKm, fiveFeet := distance.OfNauticalMiles(1), distance.OfFeet(4), distance.OfKilometers(1), distance.OfFeet(5)

	ds := []distance.Distance{
		*oneMeter,
		*zero,
		*negativeOneFeet,
		*oneFoot,
		*oneNm,
		*fourFeet,
		*oneKm,
		*fiveFeet,
	}

	distance.Sort(ds)

	isEqual(t, *negativeOneFeet, ds[0])
	isEqual(t, *zero, ds[1])
	isEqual(t, *oneFoot, ds[2])
	isEqual(t, *oneMeter, ds[3])
	isEqual(t, *fourFeet, ds[4])
	isEqual(t, *fiveFeet, ds[5])
	isEqual(t, *oneKm, ds[6])
	isEqual(t, *oneNm, ds[7])
}

func TestSum(t *testing.T) {

	isEqual(t, *distance.Zero(), *distance.Sum(nil))

	empty := []distance.Distance{}
	isEqual(t, *distance.Zero(), *distance.Sum(empty))

	one := []distance.Distance{*distance.OfFeet(1)}
	isEqual(t, *distance.OfFeet(1), *distance.Sum(one))

	sameUnits := []distance.Distance{
		*distance.OfFeet(12),
		*distance.OfFeet(22),
	}
	isEqual(t, *distance.OfFeet(34), *distance.Sum(sameUnits))

	differentUnits := []distance.Distance{
		*distance.OfFeet(12),
		*distance.OfFeet(22),
		*distance.OfMeters(1),
	}
	withinError(t, 37.28084, distance.Sum(differentUnits).InFeet(), .00001, "Different Units Sum")
}

func TestMinOf(t *testing.T) {

	one := []distance.Distance{*distance.OfFeet(1)}
	isEqual(t, *distance.OfFeet(1), *distance.MinOf(one))

	sameUnits := []distance.Distance{
		*distance.OfFeet(12),
		*distance.OfFeet(22),
	}
	isEqual(t, *distance.OfFeet(12), *distance.MinOf(sameUnits))

	differentUnits := []distance.Distance{
		*distance.OfFeet(12),
		*distance.OfFeet(22),
		*distance.OfMeters(1),
	}
	isEqual(t, *distance.OfMeters(1), *distance.MinOf(differentUnits))
}

func TestMaxOf(t *testing.T) {

	one := []distance.Distance{*distance.OfFeet(1)}
	isEqual(t, *distance.OfFeet(1), *distance.MaxOf(one))

	sameUnits := []distance.Distance{
		*distance.OfFeet(12),
		*distance.OfFeet(22),
	}
	isEqual(t, *distance.OfFeet(22), *distance.MaxOf(sameUnits))

	differentUnits := []distance.Distance{
		*distance.OfFeet(12),
		*distance.OfFeet(22),
		*distance.OfMeters(1),
	}
	isEqual(t, *distance.OfFeet(22), *distance.MaxOf(differentUnits))
}
