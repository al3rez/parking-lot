package cli

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/azbshiri/parking-lot/parking"
)

// Context is a type that is passed through to
// each parking action in a cli application.
type Context struct {
	Garage parking.Garage
	Cmd    []string
	Out    io.Writer
	OutErr io.Writer
}

func New(garage parking.Garage, cmd []string, out io.Writer, outerr io.Writer) *Context {
	return &Context{garage, cmd, out, outerr}
}

// Exec runs commands against parking lots
// e.g. parking, leaving, searching
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
		fmt.Fprintf(ctx.OutErr, "%s: command not found\n", name)
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
	if _, err := fmt.Fprintf(ctx.Out, "%s\n", strings.Join(plateNumbers, ", ")); err != nil {
		log.Fatal(err)
	}
}

func (ctx *Context) indexOfCarWithPlateNumber() {
	color := ctx.Cmd[1]
	lotNumber := ctx.Garage.IndexOfCarWithPlateNumber(color)
	if lotNumber == -1 {
		if _, err := fmt.Fprintln(ctx.Out, "Not found"); err != nil {
			log.Fatal(err)
		}
		return
	}
	if _, err := fmt.Fprintln(ctx.Out, lotNumber); err != nil {
		log.Fatal(err)
	}
}

func (ctx *Context) indexOfCarsWithColor() {
	color := ctx.Cmd[1]
	lotNumbers := ctx.Garage.IndexOfCarsWithColor(color)
	if _, err := fmt.Fprintln(ctx.Out, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(lotNumbers)), ", "), "[]")); err != nil {
		log.Fatal(err)
	}
}

func (ctx *Context) leave() {
	lotNumber, err := strconv.Atoi(ctx.Cmd[1])
	if err != nil {
		if _, err := fmt.Fprintf(ctx.Out, "Cannot convert string to int: %q\n", err); err != nil {
			log.Fatal(err)
		}
	}
	ctx.Garage.Lots = ctx.Garage.Leave(lotNumber)
	if _, err := fmt.Fprintf(ctx.Out, "Slot number %d is free\n", lotNumber); err != nil {
		log.Fatal(err)
	}
}

func (ctx *Context) park() {
	var lotNumber int
	plateNumber, color := ctx.Cmd[1], ctx.Cmd[2]
	car := parking.Car{PlateNumber: plateNumber, Color: color}
	ctx.Garage.Lots, lotNumber = ctx.Garage.Park(car)
	if lotNumber == -1 {
		if _, err := fmt.Fprintf(ctx.Out, "Sorry, parking lot is full\n"); err != nil {
			log.Fatal(err)
		}
		return
	}

	if _, err := fmt.Fprintf(ctx.Out, "Allocated slot number: %d\n", lotNumber); err != nil {
		log.Fatal(err)
	}
}

func (ctx *Context) createParkingLot() {
	capacity, err := strconv.Atoi(ctx.Cmd[1])
	if err != nil {
		if _, err := fmt.Fprintf(ctx.Out, "Invalid command value: %q\n", ctx.Cmd[1]); err != nil {
			log.Fatal(err)
		}
		return
	}
	ctx.Garage = parking.NewGarage(capacity)
	if _, err := fmt.Fprintf(ctx.Out, "Created a parking lot with %d slots\n", capacity); err != nil {
		log.Fatal(err)
	}
}

func (ctx *Context) status() {
	if _, err := fmt.Fprintln(ctx.Out, "Slot No.\tRegistration No\tColour"); err != nil {
		log.Fatal(err)
	}

	for i, lot := range ctx.Garage.Lots {
		if lot.Car == nil {
			continue
		}
		if _, err := fmt.Fprintf(ctx.Out, "%d\t\t%s\t\t%s\n", i+1, lot.Car.PlateNumber, lot.Car.Color); err != nil {
			log.Fatal(err)
		}
	}
}
