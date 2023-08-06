/*
This class is intended to make working with Courses less error prone because (1) all Course
objects are immutable and (2) the unit is always required and always accounted for.

Course is similar in spirit to the LatLong, Distance, and Speed classes. These classes (along
with java.time.Instant and java.time.Duration) are particularly powerful when used to clarify
method signatures. For example "doSomthing(LatLong, Speed, Distance, Course)" is easier to
understand than "doSomthing(Double, Double, Double, Double, Double)"

There is a difference between a "Course" and a "Heading". The Course of an aircraft is the
direction this aircraft is MOVING. The Heading of an aircraft is the direction the aircraft is
POINTING. This distinction is important when (a) an aircraft is impacted by the wind and (b) when
a helicopter flies in a direction is it's pointed.
*/
package course

import (
	"math"
	"stellarsunset/spherical/spherical"
)

type Course struct {
	angle float64
	unit  Unit
}

func North() *Course {
	return OfDegrees(0)
}

func South() *Course {
	return OfDegrees(180)
}

func East() *Course {
	return OfDegrees(90)
}

func West() *Course {
	return OfDegrees(270)
}

func Of(angle float64, unit Unit) *Course {
	return &Course{angle, unit}
}

func OfDegrees(angle float64) *Course {
	return &Course{angle, Degrees}
}

func OfRadians(angle float64) *Course {
	return &Course{angle, Radians}
}

func AngleBetween(one, two *Course) *Course {
	return OfDegrees(spherical.AngleDifference(one.InDegrees(), two.InDegrees()))
}

func (this *Course) NativeUnit() Unit {
	return this.unit
}

func (this *Course) In(desiredUnit Unit) float64 {
	if this.unit == desiredUnit {
		return this.angle
	} else {
		return this.angle * (UnitsPerDegree(desiredUnit) / UnitsPerDegree(this.unit))
	}
}

func (this *Course) InDegrees() float64 {
	return this.In(Degrees)
}

func (this *Course) InRadians() float64 {
	return this.In(Radians)
}

func (this *Course) Negate() *Course {
	return Of(-1*this.angle, this.unit)
}

func (this *Course) Abs() *Course {
	return Of(math.Abs(this.angle), this.unit)
}

func (this *Course) IsPositive() bool {
	return this.angle > 0
}

func (this *Course) IsNegative() bool {
	return this.angle < 0
}

func (this *Course) IsZero() bool {
	return this.angle == 0
}

func (this *Course) Times(scalar float64) *Course {
	return Of(this.angle*scalar, this.unit)
}

func (this *Course) Plus(that *Course) *Course {
	return Of(this.angle+that.In(this.unit), this.unit)
}

func (this *Course) Minus(that *Course) *Course {
	return Of(this.angle+that.In(this.unit), this.unit)
}

func (this *Course) IsLessThan(that *Course) bool {
	return this.angle < that.In(this.unit)
}

func (this *Course) IsLessThanOrEqualTo(that *Course) bool {
	return this.angle <= that.In(this.unit)
}

func (this *Course) IsGreaterThan(that *Course) bool {
	return this.angle > that.In(this.unit)
}

func (this *Course) IsGreaterThanOrEqualTo(that *Course) bool {
	return this.angle >= that.In(this.unit)
}

func (this *Course) Sin() float64 {
	return math.Sin(this.InRadians())
}

func (this *Course) Cos() float64 {
	return math.Cos(this.InRadians())
}

func (this *Course) Tan() float64 {
	return math.Tan(this.InRadians())
}
