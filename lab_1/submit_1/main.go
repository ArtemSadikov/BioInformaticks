package main

import (
	"fmt"
	"os"
)

func main() {
	var pattern, genome string
	var result, patternLen int = 0, 0

	fmt.Fscan(os.Stdin, &pattern)
	fmt.Fscan(os.Stdin, &genome)

	patternLen = len(pattern)

	for i := 0; i < len(genome)-patternLen-1; i++ {
		if genome[i] == pattern[0] {
			if genome[i:i+patternLen] == pattern {
				result++
			}
		}
	}

	fmt.Print(result)
}
