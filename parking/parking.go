package parking

import (
	"fmt"
)

// Car represents an automobile which can be parked in a parking lot
type Car struct {
	PlateNumber string
	Color       string
}

// Lot represents a single parking lot with a number and car
type Lot struct {
	Number int
	Car    *Car
}

// Lots represents many parking lots
type Lots []Lot

// New creates n number of empty parking lots
func New(n int) Lots {
	return make(Lots, n)
}

// Park put a car into an empty parking lot
func (lots Lots) Park(lotNumber int, car Car) (*Lot, error) {
	if lots[lotNumber].Car != nil {
		return nil, fmt.Errorf("sorry but lot number %d is taken", lotNumber)
	}

	lots[lotNumber].Car = &car
	return &lots[lotNumber], nil
}

// Leave removes a car from a parking lot
func (lots Lots) Leave(lotNumber int) *Lot {
	lots[lotNumber].Car = nil
	return &lots[lotNumber]
}

// FindCarsWithColor returns cars with a specific color parked in lots
func (lots Lots) FindCarsWithColor(name string) *Lots {
	var lotsWithColor Lots
	for _, lot := range lots {
		if lot.Car.Color == name {
			lotsWithColor = append(lotsWithColor, lot)
		}
	}

	return &lotsWithColor
}

// IndexOfCarsWithColor returns lot numbers of parked cars with specific color
func (lots Lots) IndexOfCarsWithColor(color string) []int {
	lotNumbers := []int{}
	for i, lot := range lots {
		if lot.Car.Color == color {
			lotNumbers = append(lotNumbers, i)
		}
	}

	return lotNumbers
}

// IndexOfCarWithPlateNumber returns lot number of a parked car parked with specific plate number
func (lots Lots) IndexOfCarWithPlateNumber(plateNumber string) int {
	for i, lot := range lots {
		if lot.Car.PlateNumber == plateNumber {
			return i
		}
	}

	return -1
}
