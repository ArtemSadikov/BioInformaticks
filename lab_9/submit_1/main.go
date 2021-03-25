package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var PEPTIDES_MASSES = map[string]int{
	"G": 57, "A": 71, "S": 87, "P": 97, "V": 99, "T": 101, "C": 103,
	"I": 113, "L": 113, "N": 114, "D": 115, "K": 128, "Q": 128, "E": 129,
	"M": 131, "H": 137, "F": 147, "R": 156, "Y": 163, "W": 186,
}

func main() {
	var (
		peptidesMasses []int
		spectrumMasses []int
		inputSpectrum  string
		reader         = bufio.NewReader(os.Stdin)
	)

	inputSpectrum, _ = reader.ReadString('\n')
	separatedSpectrums := strings.Split(inputSpectrum[0:len(inputSpectrum)-1], " ")
	for i := 0; i < len(separatedSpectrums); i++ {
		var spectrum, _ = strconv.Atoi(separatedSpectrums[i])
		spectrumMasses = append(spectrumMasses, spectrum)
	}
}
