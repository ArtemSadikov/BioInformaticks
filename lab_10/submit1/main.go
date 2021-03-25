package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

var PEPTIDES_MASSES = map[string]int{
	"G": 57, "A": 71, "S": 87, "P": 97, "V": 99, "T": 101, "C": 103,
	"I": 113, "L": 113, "N": 114, "D": 115, "K": 128, "Q": 128, "E": 129,
	"M": 131, "H": 137, "F": 147, "R": 156, "Y": 163, "W": 186,
}

func accumulate(arr []int) int {
	var res int = 0
	for _, value := range arr {
		res += value
	}
	return res
}

// allCount finds all repeated elements
func allCount(elem interface{}, arr []int) int {
	var res int = 0
	for _, value := range arr {
		if elem == value {
			res++
		}
	}
	return res
}

func reducePeptides(masses []int, count int) []int {
	var res []int = []int{}
	var queue []int
	if count+1 != len(masses) {
		for i := range masses {
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
	var (
		peptide      string
		experemental []int
		reader       = bufio.NewReader(os.Stdin)
		masses       = []int{}
		aminoAcid    = []int{0}
		res          = 0
	)
	fmt.Fscanf(os.Stdin, "%s", &peptide)
	inputExperemental, _ := reader.ReadString('\n')

	if runtime.GOOS == "darwin" {
		inputExperemental = inputExperemental[:len(inputExperemental)-1]
	}

	parsedInputExperemental := strings.Split(inputExperemental, " ")

	for _, parsedInput := range parsedInputExperemental {
		value, _ := strconv.Atoi(parsedInput)
		experemental = append(experemental, value)
	}

	for _, amino := range peptide {
		masses = append(masses, PEPTIDES_MASSES[string(amino)])
	}

	for i := range peptide {
		reduced := reducePeptides(masses, i)
		for _, value := range reduced {
			aminoAcid = append(aminoAcid, value)
		}
	}

	sort.Ints(aminoAcid)

	var cachedAmino int = -1
	for _, amino := range aminoAcid {
		if cachedAmino != amino {
			aminoRepeats, experementalRepeats := allCount(amino, aminoAcid), allCount(amino, experemental)
			cachedAmino = amino
			if experementalRepeats != 0 {
				if aminoRepeats > experementalRepeats {
					res += experementalRepeats
				} else {
					res += aminoRepeats
				}
			}
		}
	}
	fmt.Print(res)
}
