package cli

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/azbshiri/parking-lot/parking"
)

func TestContext_Exec_CreateParkingLot(t *testing.T) {
	g := parking.Garage{}
	buf := &bytes.Buffer{}
	type fields struct {
		garage parking.Garage
		cmd    []string
		out    io.Writer
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"creates 6 parking lot", fields{g, []string{"create_parking_lot", "6"}, buf}, "Created a parking lot with 6 slots\n"},
		{"creates 2 parking lot", fields{g, []string{"create_parking_lot", "2"}, buf}, "Created a parking lot with 2 slots\n"},
		{"invalid parking lot value", fields{g, []string{"create_parking_lot", "lf"}, buf}, "Invalid command value: \"lf\"\n"},
		{"creates 2 parking lot", fields{g, []string{"create_parking_lot", "2"}, buf}, "Created a parking lot with 2 slots\n"},
	}
	for _, tt := range tests {
		ctx := &Context{
			Garage: tt.fields.garage,
			Out:    tt.fields.out,
			OutErr: tt.fields.out,
			Cmd:    tt.fields.cmd,
		}
		t.Run(tt.name, func(t *testing.T) {
			if g = ctx.Exec(); !reflect.DeepEqual(buf.String(), tt.want) {
				t.Errorf("Context.Exec() = %v, want %v", buf.String(), tt.want)
			}
			buf.Reset()
		})
	}
}

func TestContext_Exec_Park(t *testing.T) {
	g := parking.NewGarage(2)
	buf := &bytes.Buffer{}
	type fields struct {
		garage parking.Garage
		cmd    []string
		out    io.Writer
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"park white car in slot 1", fields{g, []string{"park", "1", "white"}, buf}, "Allocated slot number: 1\n"},
		{"park black car in slot 2", fields{g, []string{"park", "2", "black"}, buf}, "Allocated slot number: 2\n"},
		{"apologize for inconvenience", fields{g, []string{"park", "3", "pink"}, buf}, "Sorry, parking lot is full\n"},
	}
	for _, tt := range tests {
		ctx := &Context{
			Garage: tt.fields.garage,
			Out:    tt.fields.out,
			OutErr: tt.fields.out,
			Cmd:    tt.fields.cmd,
		}
		t.Run(tt.name, func(t *testing.T) {
			if g = ctx.Exec(); !reflect.DeepEqual(buf.String(), tt.want) {
				t.Errorf("Context.Exec() = %v, want %v", buf.String(), tt.want)
			}
			buf.Reset()
		})
	}
}

func TestContext_Exec_Leave(t *testing.T) {
	g := parking.NewGarage(2)
	buf := &bytes.Buffer{}
	type fields struct {
		garage parking.Garage
		cmd    []string
		out    io.Writer
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"park white car in slot 1", fields{g, []string{"park", "1", "white"}, buf}, "Allocated slot number: 1\n"},
		{"park black car in slot 2", fields{g, []string{"park", "2", "black"}, buf}, "Allocated slot number: 2\n"},
		{"apologize for inconvenience", fields{g, []string{"park", "3", "pink"}, buf}, "Sorry, parking lot is full\n"},
		{"remove white car from slot 1", fields{g, []string{"leave", "1"}, buf}, "Slot number 1 is free\n"},
		{"park pink car in slot 1", fields{g, []string{"park", "2", "pink"}, buf}, "Allocated slot number: 1\n"},
	}
	for _, tt := range tests {
		ctx := &Context{
			Garage: tt.fields.garage,
			Out:    tt.fields.out,
			OutErr: tt.fields.out,
			Cmd:    tt.fields.cmd,
		}
		t.Run(tt.name, func(t *testing.T) {
			if g = ctx.Exec(); !reflect.DeepEqual(buf.String(), tt.want) {
				t.Errorf("Context.Exec() = %v, want %v", buf.String(), tt.want)
			}
			buf.Reset()
		})
	}
}

func TestContext_Exec_RegistrationNumbersForCarsWithColour(t *testing.T) {
	g := parking.NewGarage(3)
	buf := &bytes.Buffer{}
	type fields struct {
		garage parking.Garage
		cmd    []string
		out    io.Writer
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"park white car in slot 1", fields{g, []string{"park", "1", "white"}, buf}, "Allocated slot number: 1\n"},
		{"park black car in slot 2", fields{g, []string{"park", "2", "black"}, buf}, "Allocated slot number: 2\n"},
		{"park another white car in slot 3", fields{g, []string{"park", "3", "white"}, buf}, "Allocated slot number: 3\n"},
		{"registration numbers for white car", fields{g, []string{"registration_numbers_for_cars_with_colour", "white"}, buf}, "1, 3\n"},
	}
	for _, tt := range tests {
		ctx := &Context{
			Garage: tt.fields.garage,
			Out:    tt.fields.out,
			OutErr: tt.fields.out,
			Cmd:    tt.fields.cmd,
		}
		t.Run(tt.name, func(t *testing.T) {
			if g = ctx.Exec(); !reflect.DeepEqual(buf.String(), tt.want) {
				t.Errorf("Context.Exec() = %v, want %v", buf.String(), tt.want)
			}
			buf.Reset()
		})
	}
}

func TestContext_Exec_SlotNumberForRegistrationNumber(t *testing.T) {
	g := parking.NewGarage(3)
	buf := &bytes.Buffer{}
	type fields struct {
		garage parking.Garage
		cmd    []string
		out    io.Writer
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"park white car in slot 1", fields{g, []string{"park", "1", "white"}, buf}, "Allocated slot number: 1\n"},
		{"park black car in slot 2", fields{g, []string{"park", "KAH", "black"}, buf}, "Allocated slot number: 2\n"},
		{"park another white car in slot 3", fields{g, []string{"park", "3", "white"}, buf}, "Allocated slot number: 3\n"},
		{"slot number for car with registration number 1", fields{g, []string{"slot_number_for_registration_number", "KAH"}, buf}, "2\n"},
	}
	for _, tt := range tests {
		ctx := &Context{
			Garage: tt.fields.garage,
			Out:    tt.fields.out,
			OutErr: tt.fields.out,
			Cmd:    tt.fields.cmd,
		}
		t.Run(tt.name, func(t *testing.T) {
			if g = ctx.Exec(); !reflect.DeepEqual(buf.String(), tt.want) {
				t.Errorf("Context.Exec() = %v, want %v", buf.String(), tt.want)
			}
			buf.Reset()
		})
	}
}

func TestContext_Exec_SlotNumbersForCarsWithColour(t *testing.T) {
	g := parking.NewGarage(3)
	buf := &bytes.Buffer{}
	type fields struct {
		garage parking.Garage
		cmd    []string
		out    io.Writer
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"park white car in slot 1", fields{g, []string{"park", "1", "white"}, buf}, "Allocated slot number: 1\n"},
		{"park black car in slot 2", fields{g, []string{"park", "KAH", "black"}, buf}, "Allocated slot number: 2\n"},
		{"park another white car in slot 3", fields{g, []string{"park", "3", "white"}, buf}, "Allocated slot number: 3\n"},
		{"slot numbers for white cars", fields{g, []string{"slot_numbers_for_cars_with_colour", "white"}, buf}, "1, 3\n"},
	}
	for _, tt := range tests {
		ctx := &Context{
			Garage: tt.fields.garage,
			Out:    tt.fields.out,
			OutErr: tt.fields.out,
			Cmd:    tt.fields.cmd,
		}
		t.Run(tt.name, func(t *testing.T) {
			if g = ctx.Exec(); !reflect.DeepEqual(buf.String(), tt.want) {
				t.Errorf("Context.Exec() = %v, want %v", buf.String(), tt.want)
			}
			buf.Reset()
		})
	}
}

func TestContext_Exec_Status(t *testing.T) {
	status :=
"Slot No.    Registration No    Colour\n"+
"1           KMH                white\n"+
"2           KAH                black\n"

	g := parking.NewGarage(3)
	buf := &bytes.Buffer{}
	type fields struct {
		garage parking.Garage
		cmd    []string
		out    io.Writer
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"park white car in slot 1", fields{g, []string{"park", "KMH", "white"}, buf}, "Allocated slot number: 1\n"},
		{"park black car in slot 2", fields{g, []string{"park", "KAH", "black"}, buf}, "Allocated slot number: 2\n"},
		{"status", fields{g, []string{"status"}, buf}, status},
	}
	for _, tt := range tests {
		ctx := &Context{
			Garage: tt.fields.garage,
			Out:    tt.fields.out,
			OutErr: tt.fields.out,
			Cmd:    tt.fields.cmd,
		}
		t.Run(tt.name, func(t *testing.T) {
			if g = ctx.Exec(); !reflect.DeepEqual(buf.String(), tt.want) {
				t.Errorf("Context.Exec() = %v, want %v", buf.String(), tt.want)
			}
			buf.Reset()
		})
	}
}

func TestContext_Exec_NotFound(t *testing.T) {
	g := parking.NewGarage(3)
	buf := &bytes.Buffer{}
	type fields struct {
		garage parking.Garage
		cmd    []string
		out    io.Writer
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"invalid command 1", fields{g, []string{"park_left"}, buf}, "park_left: command not found\n"},
		{"invalid command 2", fields{g, []string{"park_right"}, buf}, "park_right: command not found\n"},
	}
	for _, tt := range tests {
		ctx := &Context{
			Out:    tt.fields.out,
			OutErr: tt.fields.out,
			Cmd:    tt.fields.cmd,
		}
		t.Run(tt.name, func(t *testing.T) {
			if _ = ctx.Exec(); !reflect.DeepEqual(buf.String(), tt.want) {
				t.Errorf("Context.Exec() = %v, want %v", buf.String(), tt.want)
			}
			buf.Reset()
		})
	}
}
