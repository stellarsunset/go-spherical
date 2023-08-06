package distance_test

import (
	"math"
	d "stellarsunset/spherical/distance"
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

func isEqual(t *testing.T, expected, actual d.Distance) {
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

	tol, d := .00001, d.Of(1., d.NauticalMiles)

	withinError(t, 1., d.InNauticalMiles(), tol, "InNauticalMiles()")
	withinFractionOfExpected(t, 1852., d.InMeters(), tol, "InMeters()")
	withinFractionOfExpected(t, 6076.12, d.InFeet(), tol, "InFeet()")
	withinFractionOfExpected(t, 1.15078, d.InMiles(), tol, "InMiles()")
}

func TestIn(t *testing.T) {

	tol := .00001

	oneNm := d.OfNauticalMiles(1.)
	withinError(t, 1., oneNm.In(d.NauticalMiles), tol, "In(NauticalMiles)")
	withinFractionOfExpected(t, 1852., oneNm.In(d.Meters), tol, "In(Meters)")
	withinFractionOfExpected(t, 6076.12, oneNm.In(d.Feet), tol, "In(Feet)")
	withinFractionOfExpected(t, 1.15078, oneNm.In(d.Miles), tol, "In(Miles)")

	oneMeter := d.OfMeters(1.)
	withinFractionOfExpected(t, 0.000539956803456, oneMeter.In(d.NauticalMiles), tol, "In(NauticalMiles)")
	withinError(t, 1., oneMeter.In(d.Meters), tol, "In(Meters)")
	withinError(t, .001, oneMeter.In(d.Kilometers), tol, "In(Kilometers)")
	withinFractionOfExpected(t, 3.28084, oneMeter.In(d.Feet), tol, "In(Feet)")
	withinFractionOfExpected(t, .000621371, oneMeter.In(d.Miles), tol, "In(Miles)")
}

func TestComparisonMethods(t *testing.T) {

	halfMeter, oneMeter := d.OfMeters(.5), d.OfMeters(1.)
	oneThousandMeters, oneKilometer := d.OfMeters(1000.), d.OfKilometers(1.)

	isTrue(t, halfMeter.IsLessThan(oneMeter), ".5M < 1M")
	isTrue(t, halfMeter.IsLessThanOrEqualTo(oneMeter), ".5M <= 1M")
	isTrue(t, oneMeter.IsGreaterThan(halfMeter), "1M > .5M")
	isTrue(t, oneMeter.IsGreaterThanOrEqualTo(halfMeter), "1M >= .5M")

	isTrue(t, oneKilometer.IsGreaterThanOrEqualTo(oneThousandMeters), "1KM >= 1000M")
	isTrue(t, oneKilometer.IsLessThanOrEqualTo(oneThousandMeters), "1KM <= 1000M")
}

func TestMiles(t *testing.T) {

	tol, oneMile := .00001, d.OfMiles(1.)
	withinError(t, 1., oneMile.InMiles(), tol, "OneMile.inMiles()")
	withinFractionOfExpected(t, 5280., oneMile.InFeet(), tol, "OneMile.inFeet()")
}

func TestNegate(t *testing.T) {

	oneMeter, negativeMeter := d.OfMeters(1.), d.OfMeters(-1.)
	isTrue(t, *negativeMeter == *oneMeter.Negate(), "Abs(-1M) == 1M")
}

func TestIsPositive(t *testing.T) {

	negativeOne, zero, one := d.OfFeet(-1), d.OfFeet(0.), d.OfFeet(1)

	isFalse(t, negativeOne.IsPositive(), "IsPositive(-1)")
	isTrue(t, negativeOne.Negate().IsPositive(), "IsPositive(Negate(-1))")

	isFalse(t, zero.IsPositive(), "IsPositive(0)")
	isFalse(t, zero.Negate().IsPositive(), "IsPositive(Negate(0))")

	isTrue(t, one.IsPositive(), "IsPositive(1)")
	isFalse(t, one.Negate().IsPositive(), "IsPositive(Negate(1))")
}

func TestIsNegative(t *testing.T) {

	negativeOne, zero, one := d.OfFeet(-1), d.OfFeet(0.), d.OfFeet(1)

	isTrue(t, negativeOne.IsNegative(), "IsNegative(-1)")
	isFalse(t, negativeOne.Negate().IsNegative(), "IsNegative(Negate(-1))")

	isFalse(t, zero.IsNegative(), "IsNegative(0)")
	isFalse(t, zero.Negate().IsNegative(), "IsNegative(Negate(0))")

	isFalse(t, one.IsNegative(), "IsNegative(1)")
	isTrue(t, one.Negate().IsNegative(), "IsNegative(Negate(1))")
}

func TestIsZero(t *testing.T) {

	negativeOne, zero, one := d.OfFeet(-1), d.OfFeet(0.), d.OfFeet(1)

	isFalse(t, negativeOne.IsZero(), "IsZero(-1)")
	isFalse(t, negativeOne.Negate().IsZero(), "IsZero(Negate(-1))")

	isTrue(t, zero.IsZero(), "IsZero(0)")
	isTrue(t, zero.Negate().IsZero(), "IsZero(Negate(0))")

	isFalse(t, one.IsZero(), "IsZero(1)")
	isFalse(t, one.Negate().IsZero(), "IsZero(Negate(1))")
}

func TestAbs(t *testing.T) {

	oneMeter, negativeMeter := d.OfMeters(1.), d.OfMeters(-1.)

	isTrue(t, *negativeMeter.Abs() == *oneMeter, "Abs(-1M) == 1M")
	isTrue(t, *oneMeter.Abs() == *d.OfMeters(1.), "Abs(1M) == 1M")
}

func TestTimes(t *testing.T) {

	tol, oneMeter, halfMeter := .00001, d.OfMeters(1.), d.OfMeters(.5)
	withinError(t, halfMeter.InMeters(), oneMeter.Times(.5).InMeters(), tol, "1M * .5")
}

func TestPlus(t *testing.T) {

	tol, oneFoot, fiveHalvesFeet := .00001, d.OfFeet(1.), d.OfFeet(2.5)

	sum := oneFoot.Plus(fiveHalvesFeet)
	withinError(t, 3.5, sum.InFeet(), tol, "1Ft + 2.5Ft")
}

func TestMinus(t *testing.T) {

	tol, oneFoot, fiveHalvesFeet := .00001, d.OfFeet(1.), d.OfFeet(2.5)

	sum := oneFoot.Minus(fiveHalvesFeet)

	withinError(t, -1.5, sum.InFeet(), tol, "1Ft - 2.5Ft")
	isTrue(t, sum.IsNegative(), "sum.IsNegative()")
}

func TestSort(t *testing.T) {

	oneMeter, zero, negativeOneFeet, oneFoot := d.OfMeters(1), d.Zero(), d.OfFeet(-1), d.OfFeet(1)
	oneNm, fourFeet, oneKm, fiveFeet := d.OfNauticalMiles(1), d.OfFeet(4), d.OfKilometers(1), d.OfFeet(5)

	ds := []d.Distance{
		*oneMeter,
		*zero,
		*negativeOneFeet,
		*oneFoot,
		*oneNm,
		*fourFeet,
		*oneKm,
		*fiveFeet,
	}

	d.Sort(ds)

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

	isEqual(t, *d.Zero(), *d.Sum(nil))

	empty := []d.Distance{}
	isEqual(t, *d.Zero(), *d.Sum(empty))

	one := []d.Distance{*d.OfFeet(1)}
	isEqual(t, *d.OfFeet(1), *d.Sum(one))

	sameUnits := []d.Distance{
		*d.OfFeet(12),
		*d.OfFeet(22),
	}
	isEqual(t, *d.OfFeet(34), *d.Sum(sameUnits))

	differentUnits := []d.Distance{
		*d.OfFeet(12),
		*d.OfFeet(22),
		*d.OfMeters(1),
	}
	withinError(t, 37.28084, d.Sum(differentUnits).InFeet(), .00001, "Different Units Sum")
}

func TestMinOf(t *testing.T) {

	one := []d.Distance{*d.OfFeet(1)}
	isEqual(t, *d.OfFeet(1), *d.MinOf(one))

	sameUnits := []d.Distance{
		*d.OfFeet(12),
		*d.OfFeet(22),
	}
	isEqual(t, *d.OfFeet(12), *d.MinOf(sameUnits))

	differentUnits := []d.Distance{
		*d.OfFeet(12),
		*d.OfFeet(22),
		*d.OfMeters(1),
	}
	isEqual(t, *d.OfMeters(1), *d.MinOf(differentUnits))
}

func TestMaxOf(t *testing.T) {

	one := []d.Distance{*d.OfFeet(1)}
	isEqual(t, *d.OfFeet(1), *d.MaxOf(one))

	sameUnits := []d.Distance{
		*d.OfFeet(12),
		*d.OfFeet(22),
	}
	isEqual(t, *d.OfFeet(22), *d.MaxOf(sameUnits))

	differentUnits := []d.Distance{
		*d.OfFeet(12),
		*d.OfFeet(22),
		*d.OfMeters(1),
	}
	isEqual(t, *d.OfFeet(22), *d.MaxOf(differentUnits))
}
