// Collection of simple geodesic functions for computing courses and distances between points on the surface of the
// Earth assuming it is a Spheroid.
//
// The error introduced by choosing to model as a Spheroid vs a Geoid (e.g. WGS84) is small enough to be sufficient
// for many applications and is far less computationally intensive.
//
// Note: all methods in this class referenceing latitude/longitude are implicitly expecting them in degrees
package spherical

import (
	"fmt"
	"math"
)

const (
	EarthRadiusNm float64 = 3438.14021579022
	MetersPerNm   float64 = 1852.
	MetersPerFoot float64 = .3048
	// Constant by which to multiply an angular value in degrees to obtain an
	// angular value in radians.
	degreesToRadians float64 = .017453292519943295
	// Constant by which to multiply an angular value in radians to obtain an
	// angular value in degrees.
	radiansToDegrees float64 = 57.29577951308232
	twoPi            float64 = 2. * math.Pi
	// Tolerance to be used in AlongTrackDistanceNm to avoid failures due to numeric error
	tolerance float64 = 1e-10
)

// Compute the Great Circle distance between two (latitude, longitude) coordinates in nautical miles
func DistanceInNm(lat1, lon1, lat2, lon2 float64) float64 {

	latRad1, lonRad1 := toRadians(lat1), toRadians(lon1)
	latRad2, lonRad2 := toRadians(lat2), toRadians(lon2)

	latHaver := haversine(latRad2 - latRad1)
	lonHaver := math.Cos(latRad1) * math.Cos(latRad2) * haversine(lonRad2-lonRad1)
	return EarthRadiusNm * ahaversine(latHaver+lonHaver)
}

// Compute the Great Circle course between two (latitude, longitude) coordinates in degrees
func CourseInDegrees(startLat, startLon, endLat, endLon float64) float64 {

	lat1, lon1 := toRadians(startLat), toRadians(startLon)
	lat2, lon2 := toRadians(endLat), toRadians(endLon)

	y := math.Sin(lon1-lon2) * math.Cos(lat2)
	x := (math.Cos(lat1) * math.Sin(lat2)) - (math.Sin(lat1) * math.Cos(lat2) * math.Cos(lon2-lon1))

	courseRad := twoPi - mod(math.Atan2(y, x), twoPi)
	return toDegrees(courseRad)
}

// Compute a new (latitude, longitude) location by projecting along the Great Circle defined by the starting location and direction
// the provided distance in NM
func ProjectOut(lat, lon, headingDegrees, distanceNm float64) (latitude, longitude float64) {

	if math.IsNaN(headingDegrees) {
		panic("Heading cannot be NaN")
	}
	if math.IsNaN(distanceNm) {
		panic("Distance cannot be NaN")
	}

	latRad, lonRad := toRadians(lat), toRadians(lon)

	course, dist := headingDegrees, math.Abs(distanceNm)/EarthRadiusNm
	if headingDegrees < 0. {
		course = mod(headingDegrees+180., 360.)
	}
	course = toRadians(course)

	latProj := asinReal((math.Cos(dist) * math.Sin(latRad)) + (math.Sin(dist) * math.Cos(latRad) * math.Cos(course)))

	lonNum, lonDen := math.Cos(dist)-(math.Sin(latProj)*math.Sin(latRad)), math.Cos(latProj)*math.Cos(latRad)
	dLon := acosReal(lonNum / lonDen)

	var lonProj float64
	if EarthRadiusNm*math.Abs(math.Cos(latProj)) < .01 {
		// North Pole
		lonProj = lonRad
	} else if math.Abs(dLon) < (math.Pi / 4.) {
		// Near Field
		lonProj = lonRad + math.Asin(math.Sin(dist)*math.Sin(course)/math.Cos(latProj))
	} else {
		// Far Field
		var sign float64
		if course < math.Pi {
			sign = 1.
		} else {
			sign = -1.
		}
		lonProj = lonRad + sign*dLon
	}

	return toDegrees(latProj), mod(toDegrees(lonProj)+180., 360.) - 180.
}

// Computes the spherical cross track distance (in nautical miles) between a great circle defined by the provided start and end
// point and the provided position.
//
// This value will be negative if the position is to the left of the GC and positive if it is to the right.
//
// Note: the projected location of the provided position onto the great circle used to compute this distance may not fall within
// the bounds of provided line segment.
func CrossTrackDistanceNm(startLat, startLon, endLat, endLon, posLat, posLon float64) float64 {

	distance := distanceInRadians(DistanceInNm(startLat, startLon, posLat, posLon))
	angle := toRadians(CourseInDegrees(startLat, startLon, posLat, posLon)) - toRadians(CourseInDegrees(startLat, startLon, endLat, endLon))

	return distanceInNm(math.Asin(math.Sin(distance) * math.Sin(angle)))
}

func AngleDifference(headingDegrees, headingDegrees0 float64) float64 {
	return angleDifference(headingDegrees - headingDegrees0)
}

// Computes the distance along the track (in nautical miles) from start point to end point and the provided position using the
// provided cross track distance instead of recomputing it internally.
//
// This value be negative if the point is prior to the start point of the segment along the great circle and will be positive
// if it occurs after.
//
// Note: If CTD is invalid (i.e. it is not the correct cross track distance for the {startPoint, endPoint, and p} then it may
// return NaN.
func AlongTrackDistanceNm(startLat, startLon, endLat, endLon, posLat, posLon, crossTrackDistanceNm float64) float64 {

	relativeAngle := AngleDifference(CourseInDegrees(startLat, startLon, endLat, endLon), CourseInDegrees(startLat, startLon, posLat, posLon))

	var sign float64
	if math.Abs(relativeAngle) > 90. {
		sign = -1.
	} else {
		sign = 1.
	}

	posDistance := DistanceInNm(startLat, startLon, posLat, posLon)
	cosCtd, cosPtd := math.Cos(distanceInRadians(crossTrackDistanceNm)), math.Cos(distanceInRadians(posDistance))

	// In rare cases this ratio can exceed 1 due to numeric error. We've seen cases (refer to the unit tests)
	// where ratio = 1.0000000000000002
	ratio := cosPtd / cosCtd
	if (-1.-tolerance) > ratio || (1.+tolerance) < ratio {
		panic(fmt.Sprintf("Cannot compute acos(%f). Inputs were: Start(%f, %f), End(%f, %f), Position(%f, %f), CTD(%f)",
			ratio, startLat, startLon, endLat, endLon, posLat, posLon, crossTrackDistanceNm))
	}

	// Clamp down the ratio from the tolerant range to the "valid acos range"
	ratio = math.Min(math.Max(ratio, -1.), 1.)
	return sign * distanceInNm(math.Acos(ratio))
}

func asinReal(x float64) float64 {
	return math.Asin(math.Max(-1., math.Min(1., x)))
}

func acosReal(x float64) float64 {
	return math.Acos(math.Max(-1., math.Min(1., x)))
}

func haversine(x float64) float64 {
	return (1. - math.Cos(x)) / 2.
}

func ahaversine(x float64) float64 {
	return 2. * math.Asin(math.Sqrt(x))
}

func toRadians(degrees float64) float64 {
	return degreesToRadians * degrees
}

func toDegrees(radians float64) float64 {
	return radiansToDegrees * radians
}

func mod(x, y float64) float64 {
	z := math.Remainder(x, y)
	if z < 0 {
		z += y
	}
	return z
}

// Convert a distance (in nautical miles) to the corresponding amount of "great circle radians"
func distanceInRadians(nauticalMiles float64) float64 {
	return (math.Pi / (180. * 60.)) * nauticalMiles
}

// Convert an amount of "great circle radians" to the corresponding number of nautical miles
func distanceInNm(radians float64) float64 {
	return ((180. * 60.) / math.Pi) * radians
}

func angleDifference(dz float64) float64 {
	if dz > 180. {
		return dz - 360.
	} else if dz < -180. {
		return dz + 360.
	} else {
		return dz
	}
}

func alongTrackDistanceNm(startLat, startLon, endLat, endLon, posLat, posLon float64) float64 {
	ctd := CrossTrackDistanceNm(startLat, startLon, endLat, endLon, posLat, posLon)
	return AlongTrackDistanceNm(startLat, startLon, endLat, endLon, posLat, posLon, ctd)
}
