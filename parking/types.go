package parking

// Car represents an automobile which can be parked in a parking lot
type Car struct {
	PlateNumber string
	Color       string
}

// Lot represents a single parking lot with a number and car
type Lot struct {
	Car *Car
}

// Lots represents many parking lots
type Lots []*Lot

// Garage represents a building which holds parking lots
type Garage struct {
	Capacity int
	Lots     Lots
}
