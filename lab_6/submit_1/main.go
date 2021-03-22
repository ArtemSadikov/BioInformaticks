package main

import (
	"fmt"
	"os"
)

func main() {
	var peptideLength int
	fmt.Fscanf(os.Stdin, "%d", &peptideLength)
	fmt.Print(peptideLength * (peptideLength - 1))
}
