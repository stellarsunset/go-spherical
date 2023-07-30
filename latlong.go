package spherical

type Positioner interface {
	LatLong() *latlong
}

type latlong struct {
	latitude  float64
	longitude float64
}

func NewLatLong(latitude, longitude float64) *latlong {
	return &latlong{latitude: latitude, longitude: longitude}
}

func (this *latlong) DistanceInNm(that *latlong) float64 {
	return DistanceInNm(this.latitude, this.longitude, that.latitude, that.longitude)
}

func (this *latlong) CourseInDegrees(that *latlong) float64 {
	return CourseInDegrees(this.latitude, this.longitude, that.latitude, that.longitude)
}
