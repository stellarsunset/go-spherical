/*
 * This Distance class is intended to make working with Distances less error prone because (1) all
 * Distance objects are immutable and (2) the unit is always required and always accounted for.
 *
 * This class is extremely similar in spirit and design to java.time.Duration and java.time.Instant.
 */
package distance

import "math"

type distance struct {
	amount float64
	unit   Unit
}

func Zero() *distance {
	return &distance{0., Meters}
}

func Of(amount float64, unit Unit) *distance {
	return &distance{amount, unit}
}

func OfNauticalMiles(amount float64) *distance {
	return &distance{amount, NauticalMiles}
}

func OfFeet(amount float64) *distance {
	return &distance{amount, Feet}
}

func OfMeters(amount float64) *distance {
	return &distance{amount, Meters}
}

func OfKilometers(amount float64) *distance {
	return &distance{amount, Kilometers}
}

func OfMiles(amount float64) *distance {
	return &distance{amount, Miles}
}

// The Unit this distance was originally defined with
func (this *distance) NativeUnit() Unit {
	return this.unit
}

func (this *distance) In(desiredUnit Unit) float64 {
	if this.unit == desiredUnit {
		return this.amount
	} else {
		return this.amount * (UnitsPerMeter(desiredUnit) / UnitsPerMeter(this.unit))
	}
}

func (this *distance) InNauticalMiles() float64 {
	return this.In(NauticalMiles)
}

func (this *distance) InFeet() float64 {
	return this.In(Feet)
}

func (this *distance) InMeters() float64 {
	return this.In(Meters)
}

func (this *distance) InKilometers() float64 {
	return this.In(Kilometers)
}

func (this *distance) InMiles() float64 {
	return this.In(Miles)
}

func (this *distance) Negate() *distance {
	return Of(-this.amount, this.unit)
}

func (this *distance) Abs() *distance {
	return Of(math.Abs(this.amount), this.unit)
}

func (this *distance) IsPositive() bool {
	return this.amount > 0.
}

func (this *distance) IsNegative() bool {
	return this.amount < 0.
}

func (this *distance) IsZero() bool {
	return this.amount == 0.
}

func (this *distance) Times(scalar float64) *distance {
	return Of(this.amount*scalar, this.unit)
}

func (this *distance) Plus(that *distance) *distance {
	return Of(this.amount+that.In(this.unit), this.unit)
}

func (this *distance) Minus(that *distance) *distance {
	return Of(this.amount-that.In(this.unit), this.unit)
}

func (this *distance) IsLessThan(that *distance) bool {
	return this.amount < that.In(this.unit)
}

func (this *distance) IsLessThanOrEqualTo(that *distance) bool {
	return this.amount <= that.In(this.unit)
}

func (this *distance) IsGreaterThan(that *distance) bool {
	return this.amount > that.In(this.unit)
}

func (this *distance) IsGreaterThanOrEqualTo(that *distance) bool {
	return this.amount >= that.In(this.unit)
}
