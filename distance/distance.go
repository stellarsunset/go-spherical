/*
 * This Distance class is intended to make working with Distances less error prone because (1) all
 * Distance objects are immutable and (2) the unit is always required and always accounted for.
 *
 * This class is extremely similar in spirit and design to java.time.Duration and java.time.Instant.
 */
package distance

import (
	"math"
	"sort"
)

type Distance struct {
	amount float64
	unit   Unit
}

func Zero() *Distance {
	return &Distance{0., Meters}
}

func Of(amount float64, unit Unit) *Distance {
	return &Distance{amount, unit}
}

func OfNauticalMiles(amount float64) *Distance {
	return &Distance{amount, NauticalMiles}
}

func OfFeet(amount float64) *Distance {
	return &Distance{amount, Feet}
}

func OfMeters(amount float64) *Distance {
	return &Distance{amount, Meters}
}

func OfKilometers(amount float64) *Distance {
	return &Distance{amount, Kilometers}
}

func OfMiles(amount float64) *Distance {
	return &Distance{amount, Miles}
}

// The Unit this distance was originally defined with
func (this *Distance) NativeUnit() Unit {
	return this.unit
}

func (this *Distance) In(desiredUnit Unit) float64 {
	if this.unit == desiredUnit {
		return this.amount
	} else {
		return this.amount * (UnitsPerMeter(desiredUnit) / UnitsPerMeter(this.unit))
	}
}

func (this *Distance) InNauticalMiles() float64 {
	return this.In(NauticalMiles)
}

func (this *Distance) InFeet() float64 {
	return this.In(Feet)
}

func (this *Distance) InMeters() float64 {
	return this.In(Meters)
}

func (this *Distance) InKilometers() float64 {
	return this.In(Kilometers)
}

func (this *Distance) InMiles() float64 {
	return this.In(Miles)
}

func (this *Distance) Negate() *Distance {
	return Of(-this.amount, this.unit)
}

func (this *Distance) Abs() *Distance {
	return Of(math.Abs(this.amount), this.unit)
}

func (this *Distance) IsPositive() bool {
	return this.amount > 0.
}

func (this *Distance) IsNegative() bool {
	return this.amount < 0.
}

func (this *Distance) IsZero() bool {
	return this.amount == 0.
}

func (this *Distance) Times(scalar float64) *Distance {
	return Of(this.amount*scalar, this.unit)
}

func (this *Distance) Plus(that *Distance) *Distance {
	return Of(this.amount+that.In(this.unit), this.unit)
}

func (this *Distance) Minus(that *Distance) *Distance {
	return Of(this.amount-that.In(this.unit), this.unit)
}

func (this *Distance) IsLessThan(that *Distance) bool {
	return this.amount < that.In(this.unit)
}

func (this *Distance) IsLessThanOrEqualTo(that *Distance) bool {
	return this.amount <= that.In(this.unit)
}

func (this *Distance) IsGreaterThan(that *Distance) bool {
	return this.amount > that.In(this.unit)
}

func (this *Distance) IsGreaterThanOrEqualTo(that *Distance) bool {
	return this.amount >= that.In(this.unit)
}

type byAmount []Distance

func (a byAmount) Len() int {
	return len(a)
}

func (a byAmount) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAmount) Less(i, j int) bool {
	return a[i].IsLessThan(&a[j])
}

// Sort the provided slice of distances based on their unit-aligned amounts
func Sort(distances []Distance) {
	sort.Sort(byAmount(distances))
}

func Sum(distances []Distance) *Distance {
	switch l := len(distances); l {
	case 0:
		return Zero()
	case 1:
		return &distances[0]
	default:
		amount, unit := distances[0].amount, distances[0].unit
		for i := 1; i < l; i++ {
			amount += distances[i].In(unit)
		}
		return Of(amount, unit)
	}
}

func Min(one, two *Distance) *Distance {
	if one.IsLessThanOrEqualTo(two) {
		return one
	} else {
		return two
	}
}

func MinOf(distances []Distance) *Distance {
	switch l := len(distances); l {
	case 0:
		return nil
	case 1:
		return &distances[0]
	default:
		min := &distances[0]
		for i := 1; i < l; i++ {
			min = Min(min, &distances[i])
		}
		return min
	}
}

func Max(one, two *Distance) *Distance {
	if one.IsGreaterThanOrEqualTo(two) {
		return one
	} else {
		return two
	}
}

func MaxOf(distances []Distance) *Distance {
	switch l := len(distances); l {
	case 0:
		return nil
	case 1:
		return &distances[0]
	default:
		max := &distances[0]
		for i := 1; i < l; i++ {
			max = Max(max, &distances[i])
		}
		return max
	}
}
