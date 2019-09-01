package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
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

	var stdin bool
	flag.BoolVar(&stdin, "stdin", false,
		"Execute parking lot instructions from standard streams")

	var interactive bool
	flag.BoolVar(&interactive, "i", false,
		"Execute parking lot instructions from interactive shell")

	flag.Parse()

	if len(filename) > 1 {
		file, err := os.Open(filename)
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
		var garage parking.Garage
		for _, cmd := range cmds {
			ctx := cli.New(garage, cmd, os.Stdout, os.Stderr)
			garage = ctx.Exec()
		}
	}

	if interactive {
		reader := bufio.NewReader(os.Stdin)
		var garage parking.Garage
		for {
			// if reading from standard streams avoid prompt
			if !stdin {
				fmt.Print("parking-lot> ")
			}

			raw, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			raw = strings.TrimSuffix(raw, "\n")
			cmd := strings.Split(raw, " ")

			ctx := cli.New(garage, cmd, os.Stdout, os.Stderr)
			garage = ctx.Exec()
		}
	}
}
