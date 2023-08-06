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

type Course struct {
	angle float64
	unit  Unit
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

func (this *Course) NativeUnit() Unit {
	return this.unit
}
