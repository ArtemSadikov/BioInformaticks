package main

import (
	"fmt"
	"os"
)

var PEPTIDES_MASSES = []int{
	57, 71, 87, 97, 99, 101, 103,
	113, 114, 115, 128, 129,
	131, 137, 147, 156, 163, 186,
}

func main() {
	var (
		n     int
		count []int
	)

	fmt.Fscanf(os.Stdin, "%d", &n)
	count = make([]int, n+1)
	count[0] = 1

	var rangeArray = make([]int, n+1-PEPTIDES_MASSES[0])

	for i := 0; i < len(rangeArray); i++ {
		rangeArray[i] = PEPTIDES_MASSES[0] + i
	}

	for _, i := range rangeArray {
		for _, m := range PEPTIDES_MASSES {
			if i-m >= 0 {
				count[i] += count[i-m]
			}
		}
	}
	fmt.Print(count[n])
}
