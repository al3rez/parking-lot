package parking

import (
	"reflect"
	"testing"
)

func emptyLots(lots Lots) (len int) {
	for _, lot := range lots {
		if lot.Car == nil {
			len += 1
		}
	}

	return
}

func TestGarage_Park(t *testing.T) {
	g := Garage{
		Lots: Lots{&Lot{}, &Lot{}, &Lot{}},
	}

	type args struct {
		car Car
	}
	tests := []struct {
		name   string
		args   args
		want   Lots
		want1  int
	}{
		{"park a car in slot 1", args{Car{}}, g.Lots, 1},
		{"park a car in slot 2 ", args{Car{}}, g.Lots, 2},
		{"park a car in slot 3 ", args{Car{}}, g.Lots, 3},
		{"error when parking is full", args{Car{}}, g.Lots, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := g.Park(tt.args.car)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Park() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Park() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGarage_Leave(t *testing.T) {
	g := Garage{
		Lots: Lots{&Lot{}, &Lot{}, &Lot{}},
	}
	g.Park(Car{})
	g.Park(Car{})
	g.Park(Car{})

	type args struct {
		lotNumber int
	}
	tests := []struct {
		name   string
		args   args
		want   int
	}{
		{"remove car from slot 1", args{1}, 1},
		{"remove car from slot 2", args{2}, 2},
		{"remove car from slot 3", args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.Leave(tt.args.lotNumber); emptyLots(got) != tt.want {
				t.Errorf("Leave() = %v, want %v", emptyLots(got), tt.want)
			}
		})
	}
}

func TestGarage_FindCarsWithColor(t *testing.T) {
	g := Garage{
		Lots: Lots{&Lot{}, &Lot{}, &Lot{}},
	}
	g.Park(Car{"1", "White"})
	g.Park(Car{"2", "Black"})
	g.Park(Car{"3", "White"})

	type args struct {
		color string
	}
	tests := []struct {
		name   string
		args   args
		want   int
	}{
		{"find 2 cars with color white", args{"White"}, 2},
		{"find 1 car with color black", args{"Black"}, 1},
		{"find 0 car with color pink", args{"Pink"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.FindCarsWithColor(tt.args.color); len(got) != tt.want {
				t.Errorf("FindCarsWithColor() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestGarage_IndexOfCarsWithColor(t *testing.T) {
	g := Garage{
		Lots: Lots{&Lot{}, &Lot{}, &Lot{}},
	}
	g.Park(Car{"1", "White"})
	g.Park(Car{"2", "Black"})
	g.Park(Car{"3", "White"})

	type args struct {
		color string
	}
	tests := []struct {
		name   string
		args   args
		want   int
	}{
		{"index 2 cars with color white", args{"White"}, 2},
		{"index 1 cars with color black", args{"Black"}, 1},
		{"index 0 cars with color pink", args{"Pink"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.IndexOfCarsWithColor(tt.args.color); len(got) != tt.want {
				t.Errorf("IndexOfCarsWithColor() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestGarage_IndexOfCarWithPlateNumber(t *testing.T) {
	g := Garage{
		Lots: Lots{&Lot{}, &Lot{}, &Lot{}},
	}
	g.Park(Car{"1", "White"})
	g.Park(Car{"2", "Black"})
	g.Park(Car{"3", "White"})

	type fields struct {
		Lots Lots
	}
	type args struct {
		plateNumber string
	}
	tests := []struct {
		name   string
		args   args
		want   int
	}{
		{"index car with plate number 1", args{"1"}, 1},
		{"index car with plate number 2", args{"2"}, 2},
		{"index car with plate number 3", args{"3"}, 3},
		{"index -1 when cannot find", args{"4"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.IndexOfCarWithPlateNumber(tt.args.plateNumber); got != tt.want {
				t.Errorf("IndexOfCarWithPlateNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}