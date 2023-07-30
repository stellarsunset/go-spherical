package distance

type distance struct {
	amount float64
	unit   Unit
}

func OfNauticalMiles(amount float64) *distance {
	return &distance{amount: amount, unit: NauticalMiles}
}

func OfFeet(amount float64) *distance {
	return &distance{amount: amount, unit: Feet}
}

func OfMeters(amount float64) *distance {
	return &distance{amount: amount, unit: Meters}
}

func OfKilometers(amount float64) *distance {
	return &distance{amount: amount, unit: Kilometers}
}

func OfMiles(amount float64) *distance {
	return &distance{amount: amount, unit: Miles}
}
