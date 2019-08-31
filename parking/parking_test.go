package parking

import (
	"reflect"
	"testing"
)

func TestLots_Park(t *testing.T) {
	lots := make(Lots, 3)
	type args struct {
		lotNumber int
		car       Car
	}
	tests := []struct {
		name    string
		lots    Lots
		args    args
		want    *Lot
		wantErr bool
	}{
		{"park a car in given lot 0", lots, args{0, Car{}}, &lots[0], false},
		{"park a car in given lot 1", lots, args{1, Car{}}, &lots[1], false},
		{"park a car in given lot 2", lots, args{2, Car{}}, &lots[1], false},
		{"return an error if lot is taken", lots, args{0, Car{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.lots.Park(tt.args.lotNumber, tt.args.car)
			if (err != nil) != tt.wantErr {
				t.Errorf("Park() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Park() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLots_Leave(t *testing.T) {
	lots := make(Lots, 3)
	_, err := lots.Park(0, Car{})
	if err != nil {
		t.Fatal(err)
	}

	_, err = lots.Park(1, Car{})
	if err != nil {
		t.Fatal(err)
	}

	_, err = lots.Park(2, Car{})
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		lotNumber int
	}
	tests := []struct {
		name string
		lots Lots
		args args
		want *Lot
	}{
		{"remove a car from given lot 0", lots, args{0}, &lots[0]},
		{"remove a car from given lot 1", lots, args{1}, &lots[1]},
		{"remove a car from given lot 2", lots, args{2}, &lots[2]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lots.Leave(tt.args.lotNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Leave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want Lots
	}{
		{"create 2 empty parking lots", args{2}, Lots{{0, nil},{0, nil},}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLots_FindCarsWithColor(t *testing.T) {
	lots := make(Lots, 6)
	lots.Park(0, Car{"1", "white"})
	lots.Park(1, Car{"2", "white"})
	lots.Park(2, Car{"3", "blue"})
	lots.Park(3, Car{"4", "pink"})
	lots.Park(4, Car{"5", "pink"})
	lots.Park(5, Car{"6", "black"})

	type args struct {
		name string
	}
	tests := []struct {
		name string
		lots Lots
		args args
		want int
	}{
		{"find cars with color white", lots, args{"white"}, 2},
		{"find cars with color blue", lots, args{"blue"}, 1},
		{"find cars with color pink", lots, args{"pink"}, 2},
		{"find cars with color black", lots, args{"black"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lots.FindCarsWithColor(tt.args.name); len(*got) != tt.want {
				t.Errorf("FindCarsWithColor() = %v, want %v", len(*got), tt.want)
			}
		})
	}
}

func TestLots_IndexOfCarsWithColor(t *testing.T) {
	lots := make(Lots, 6)
	lots.Park(0, Car{"1", "white"})
	lots.Park(1, Car{"2", "white"})
	lots.Park(2, Car{"3", "blue"})
	lots.Park(3, Car{"4", "pink"})
	lots.Park(4, Car{"5", "pink"})
	lots.Park(5, Car{"6", "black"})

	type args struct {
		color string
	}
	tests := []struct {
		name string
		lots Lots
		args args
		want []int
	}{
		{"return lot numbers of cars with color white", lots, args{"white"}, []int{0, 1}},
		{"return lot numbers of cars with color pink", lots, args{"pink"}, []int{3, 4}},
		{"return lot numbers of cars with color black", lots, args{"black"}, []int{5}},
		{"return empty array when cannot find cars", lots, args{"purple"}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lots.IndexOfCarsWithColor(tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexOfCarsWithColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLots_IndexOfCarWithPlateNumber(t *testing.T) {
	lots := make(Lots, 2)
	lots.Park(0, Car{"1", "white"})
	lots.Park(1, Car{"2", "black"})

	type args struct {
		plateNumber string
	}
	tests := []struct {
		name string
		lots Lots
		args args
		want int
	}{
		{"return lot number of car with plate number 1", lots, args{"1"}, 0},
		{"return lot number of car with plate number 2", lots, args{"2"}, 1},
		{"return -1 when cannot find any cars", lots, args{"white"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lots.IndexOfCarWithPlateNumber(tt.args.plateNumber); got != tt.want {
				t.Errorf("IndexOfCarWithPlateNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}