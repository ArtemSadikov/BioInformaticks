package main

import (
	"fmt"
	"os"
)

var REVERSABLE = map[string]string{
	"A": "T",
	"T": "A",
	"G": "C",
	"C": "G",
}

func main() {
	var (
		inputGenome  string
		resultGenome string
	)

	fmt.Fscanf(os.Stdin, "%s", &inputGenome)

	for i := len(inputGenome) - 1; i > -1; i-- {
		resultGenome += REVERSABLE[string(inputGenome[i])]
	}

	fmt.Print(resultGenome)
}
