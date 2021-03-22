package main

import (
	"fmt"
	"os"
	"sort"
)

var PEPTIDES_MASSES = map[string]int{
	"G": 57, "A": 71, "S": 87, "P": 97, "V": 99, "T": 101, "C": 103,
	"I": 113, "L": 113, "N": 114, "D": 115, "K": 128, "Q": 128, "E": 129,
	"M": 131, "H": 137, "F": 147, "R": 156, "Y": 163, "W": 186,
}

func accumulate(arr []int) int {
	var res int = 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func reducePeptides(masses []int, count int) []int {
	var res []int = []int{}
	var queue []int
	if count+1 != len(masses) {
		for i := 0; i < len(masses); i++ {
			if (i + count) > len(masses)-1 {
				queue = masses[i:]
				queue = append(queue, masses[0:count+1-len(queue)]...)
			} else {
				queue = masses[i : i+count+1]
			}
			res = append(res, accumulate(queue))
		}
	} else {
		res = append(res, accumulate(masses))
	}

	return res
}

func main() {
	var peptide string
	fmt.Fscanf(os.Stdin, "%s", &peptide)

	var (
		masses = []int{}
		res    = []int{0}
	)

	for i := 0; i < len(peptide); i++ {
		masses = append(masses, PEPTIDES_MASSES[string(peptide[i])])
	}

	for i := 0; i < len(peptide); i++ {
		var reduced []int = reducePeptides(masses, i)
		for j := 0; j < len(reduced); j++ {
			res = append(res, reduced[j])
		}
	}

	sort.Ints(res)
	for i, mass := range res {
		if i == len(res)-1 {
			fmt.Print(mass)
		} else {
			fmt.Print(mass, " ")
		}
	}
}
