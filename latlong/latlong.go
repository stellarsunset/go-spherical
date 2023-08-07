package latlong

import (
	"errors"
	"fmt"
	sph "stellarsunset/spherical"
	crs "stellarsunset/spherical/course"
	dist "stellarsunset/spherical/distance"
)

type LatLong struct {
	latitude  float64
	longitude float64
}

func checkLatitude(latitude float64) (float64, error) {
	if latitude <= -90 || 90. <= latitude {
		return latitude, errors.New(fmt.Sprintf("Latitude is out of range (-90, 90): %f", latitude))
	}
	return latitude, nil
}

func checkLongitude(longitude float64) (float64, error) {
	if longitude <= -180 || 180. <= longitude {
		return longitude, errors.New(fmt.Sprintf("Longitude is out of range (-180, 180): %f", longitude))
	}
	return longitude, nil
}

// Creates a new LatLong struct from the provided latitude and longitude values in degrees, panicking with an error code if the
// provided latitude or longitude fall outside the accepted ranges (-90, 90), (-180, 180).
func NewLatLong(latitude, longitude float64) *LatLong {

	_, laterr := checkLatitude(latitude)
	_, lonerr := checkLongitude(longitude)

	if laterr != nil || lonerr != nil {
		panic(errors.Join(laterr, lonerr))
	}

	return &LatLong{latitude, longitude}
}

func (this *LatLong) Latitude() float64 {
	return this.latitude
}

func (this *LatLong) Longitude() float64 {
	return this.longitude
}

func (this *LatLong) DistanceTo(that *LatLong) *dist.Distance {
	return dist.OfNauticalMiles(this.DistanceInNm(that))
}

func (this *LatLong) DistanceInNm(that *LatLong) float64 {
	return sph.DistanceInNm(this.latitude, this.longitude, that.latitude, that.longitude)
}

func (this *LatLong) CourseTo(that *LatLong) *crs.Course {
	return crs.OfDegrees(this.CourseInDegrees(that))
}

func (this *LatLong) CourseInDegrees(that *LatLong) float64 {
	return sph.CourseInDegrees(this.latitude, this.longitude, that.latitude, that.longitude)
}

// Return true if this LatLong is within (<=) the given distance of the provided LatLong
func (this *LatLong) IsWithin(distance *dist.Distance, that *LatLong) bool {
	return this.DistanceTo(that).IsLessThanOrEqualTo(distance)
}

func (this *LatLong) project(course *crs.Course, distance *dist.Distance) *LatLong {
	return this.ProjectOut(course.InDegrees(), distance.InNauticalMiles())
}

// Projects outwards from this LatLong along the provided course (in degrees) the provided distance (in nm) returning a new LL.
func (this *LatLong) ProjectOut(direction, distance float64) *LatLong {
	latitude, longitude := sph.ProjectOut(this.latitude, this.longitude, direction, distance)
	return &LatLong{latitude, longitude}
}

func (this *LatLong) CrossTrackDistanceTo(start, end *LatLong) *dist.Distance {
	return dist.OfNauticalMiles(this.CrossTrackDistanceNm(start, end))
}

func (this *LatLong) CrossTrackDistanceNm(start, end *LatLong) float64 {
	return sph.CrossTrackDistanceNm(start.latitude, start.longitude, end.latitude, end.longitude, this.latitude, this.longitude)
}

func (this *LatLong) AlongTrackDistanceTo(start, end *LatLong, crossTrack *dist.Distance) *dist.Distance {
	return dist.OfNauticalMiles(this.AlongTrackDistanceNm(start, end, crossTrack.InNauticalMiles()))
}

func (this *LatLong) AlongTrackDistanceNm(start, end *LatLong, crossTrackDistanceNm float64) float64 {
	return sph.AlongTrackDistanceNm(start.latitude, start.longitude, end.latitude, end.longitude, this.latitude, this.longitude, crossTrackDistanceNm)
}
