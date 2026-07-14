package main

import (
	"fmt"
	"io"
	"os"

	"tetris-optimizer/internal/solver"
	"tetris-optimizer/internal/tetromino"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Println("ERROR")
	}
}

func run(args []string, output io.Writer) error {
	if len(args) != 1 {
		return tetromino.ErrInvalidFormat
	}

	data, err := os.ReadFile(args[0])
	if err != nil {
		return err
	}
	pieces, err := tetromino.Parse(string(data))
	if err != nil {
		return err
	}

	fmt.Fprintln(output, solver.Solve(pieces))
	return nil
}
