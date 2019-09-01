package parking

// NewGarage instantiates a new garage with n capacity of empty parking lots
func NewGarage(cap int) Garage {
	g := Garage{}
	g.Lots = make(Lots, cap)
	for i, _ := range g.Lots {
		g.Lots[i] = &Lot{}
	}
	return g
}

// Park put a car into an empty parking lot
func (g Garage) Park(car Car) (Lots, int) {
	lotNumber := -1
	// Check if there are any empty parking lots
	// NOTE: This will preserve order of leave/parking
	for i, lot := range g.Lots {
		if lot.Car == nil {
			g.Lots[i].Car = &car
			lotNumber = i + 1
			break
		}
	}

	return g.Lots, lotNumber
}

// Leave removes a car from a parking lot
func (g Garage) Leave(lotNumber int) Lots {
	g.Lots[lotNumber-1].Car = nil
	return g.Lots
}

// FindCarsWithColor returns cars with a specific color parked in lots
func (g Garage) FindCarsWithColor(color string) []Car {
	var cars []Car
	for _, lot := range g.Lots {
		if lot.Car != nil && lot.Car.Color == color {
			cars = append(cars, *lot.Car)
		}
	}

	return cars
}

// IndexOfCarsWithColor returns lot numbers of parked cars with specific color
func (g Garage) IndexOfCarsWithColor(color string) []int {
	var lotNumbers []int
	for i, lot := range g.Lots {
		if lot.Car != nil && lot.Car.Color == color {
			lotNumbers = append(lotNumbers, i+1)
		}
	}

	return lotNumbers
}

// IndexOfCarWithPlateNumber returns lot number of a parked car parked with specific plate number
func (g Garage) IndexOfCarWithPlateNumber(plateNumber string) int {
	for i, lot := range g.Lots {
		if lot.Car != nil && lot.Car.PlateNumber == plateNumber {
			return i + 1
		}
	}

	return -1
}
