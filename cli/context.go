package cli

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/azbshiri/parking-lot/parking"
)

// Context represents a command-line interface which will operate on a garage using given commands
type Context struct {
	Garage parking.Garage
	Cmd    []string
	Out io.Writer
}

func New(garage parking.Garage, cmd []string, out io.Writer) *Context {
	return &Context{garage, cmd, out}
}

func (ctx *Context) Exec() parking.Garage {
	switch name := ctx.Cmd[0]; name {
	case "create_parking_lot":
		ctx.createParkingLot()
	case "park":
		ctx.park()
	case "leave":
		ctx.leave()
	case "slot_numbers_for_cars_with_colour":
		ctx.indexOfCarsWithColor()
	case "slot_number_for_registration_number":
		ctx.indexOfCarWithPlateNumber()
	case "registration_numbers_for_cars_with_colour":
		ctx.findCarsWithColor()
	case "status":
		ctx.status()
	case "exit":
		os.Exit(1)
	default:
		fmt.Fprintf(ctx.Out, "%s: command not found\n", name)
		os.Exit(1)
	}

	return ctx.Garage
}

func (ctx *Context) findCarsWithColor() {
	color := ctx.Cmd[1]
	cars := ctx.Garage.FindCarsWithColor(color)
	var plateNumbers []string
	for _, car := range cars {
		plateNumbers = append(plateNumbers, car.PlateNumber)
	}
	fmt.Fprintf(ctx.Out, "%s\n", strings.Join(plateNumbers, ", "))
}

func (ctx *Context) indexOfCarWithPlateNumber() {
	color := ctx.Cmd[1]
	lotNumber := ctx.Garage.IndexOfCarWithPlateNumber(color)
	if lotNumber == -1 {
		fmt.Fprintln(ctx.Out, "Not found")
		return
	}
	fmt.Fprintln(ctx.Out, lotNumber)
}

func (ctx *Context) indexOfCarsWithColor() {
	color := ctx.Cmd[1]
	lotNumbers := ctx.Garage.IndexOfCarsWithColor(color)
	fmt.Fprintln(ctx.Out, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(lotNumbers)), ", "), "[]"))
}

func (ctx *Context) leave() {
	lotNumber, err := strconv.Atoi(ctx.Cmd[1])
	if err != nil {
		fmt.Fprintf(ctx.Out, "Cannot convert string to int: %q\n", err)
		os.Exit(1)
	}
	ctx.Garage.Lots = ctx.Garage.Leave(lotNumber)
	fmt.Fprintf(ctx.Out, "Slot number %d is free\n", lotNumber)
}

func (ctx *Context) park() {
	var lotNumber int
	plateNumber, color := ctx.Cmd[1], ctx.Cmd[2]
	car := parking.Car{PlateNumber: plateNumber, Color: color}
	ctx.Garage.Lots, lotNumber = ctx.Garage.Park(car)
	if lotNumber == -1 {
		fmt.Printf("Sorry, parking lot is full\n")
		return
	}
	fmt.Fprintf(ctx.Out, "Allocated slot number: %d\n", lotNumber)
}

func (ctx *Context) createParkingLot() {
	capacity, err := strconv.Atoi(ctx.Cmd[1])
	if err != nil {
		fmt.Fprint(ctx.Out, "Cannot convert string to int: %q\n", err)
	}
	ctx.Garage.Lots = make(parking.Lots, capacity)
	for i, _ := range ctx.Garage.Lots {
		ctx.Garage.Lots[i] = &parking.Lot{}
	}
	fmt.Fprintf(ctx.Out, "Created a parking lot with %d slots\n", capacity)
}

func (ctx *Context) status() {
	fmt.Printf("Slot No.\tRegisteration No\tColour\n")
	for i, lot := range ctx.Garage.Lots {
		if lot.Car == nil {
			continue
		}
		fmt.Fprintf(ctx.Out, "%d\t\t%s\t\t%s\n", i+1, lot.Car.PlateNumber, lot.Car.Color)
	}
}
