package distance

type Kilometers float64
type Miles float64

// Add methods to Miles
// (m Miles) is the receiver telling us what type
// to add this function on
func (m Miles) ToKilometers() Kilometers {
	return Kilometers(m * 1.60934)
}

func (km Kilometers) ToMiles() Miles {
	return Miles(km * 0.621371)
}
