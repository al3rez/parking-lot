package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/azbshiri/parking-lot/cli"
	"github.com/azbshiri/parking-lot/parking"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "",
		"Execute parking lot instructions from given file")

	var interactive bool
	flag.BoolVar(&interactive, "i", false,
		"Execute parking lot instructions from interactive shell")

	flag.Parse()

	file, err := os.Open("file_inputs.txt")
	if err != nil {
		log.Fatalf("Cannot open file: %q", err)
	}

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1

	cmds, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Cannot parse file: %q", err)
	}

	// read ticketing commands from file name
	if len(filename) > 1 {
		var garage parking.Garage
		for _, cmd := range cmds {
			ctx := cli.New(garage, cmd)
			garage = ctx.Exec()
		}
	}

	// read ticketing commands from interactive shell
	if interactive {
		reader := bufio.NewReader(os.Stdin)
		var garage parking.Garage
		for {
			fmt.Print("parking-lot> ")
			raw, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Cannot read string: %q\n", err)
			}
			raw = strings.TrimSuffix(raw, "\n")
			cmd := strings.Split(raw, " ")
			fmt.Println(cmd)

			ctx := cli.New(garage, cmd)
			garage = ctx.Exec()
		}
	}
}
