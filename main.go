package main

import (
	"fmt"
	"os"

	"tdd_kata_go/string_calculator"
)

func main() {
	expr := ""
	args := os.Args[1:]
	if len(args) > 0 {
		expr = args[0]
	}
	r := string_calculator.Add(expr)
	fmt.Printf("Add(%s): %d", expr, r)
}
