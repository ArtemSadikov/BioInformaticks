package main

import (
	"fmt"
	"os"
)

type GenomeRepeats struct {
	genome  string
	repeats int
}

func isExist(genomeRepeats []GenomeRepeats, finding string) int {
	if len(genomeRepeats) == 0 {
		return -1
	}

	for i := 0; i < len(genomeRepeats); i++ {
		if genomeRepeats[i].genome == finding {
			return i
		}
	}
	return -1
}

func getRepeats(genomeRepeats *[]GenomeRepeats, genome string, finding string) {
	var findingLen int = len(finding)
	var res GenomeRepeats = GenomeRepeats{genome: finding, repeats: 0}
	var findingIndex int = isExist(*genomeRepeats, finding)
	if findingIndex == -1 {
		for i := 0; i < len(genome)-findingLen+1; i++ {
			if genome[i:i+findingLen] == finding {
				res.repeats++
			}
		}
		*genomeRepeats = append(*genomeRepeats, res)
	}
}

func main() {
	var (
		genome        string
		textLen       int
		genomeRepeats []GenomeRepeats
	)

	fmt.Fscanf(os.Stdin, "%s", &genome)
	fmt.Fscanf(os.Stdin, "%d", &textLen)

	for i := 0; i < len(genome)-textLen; i++ {
		getRepeats(&genomeRepeats, genome, genome[i:i+textLen])
	}

	var maxRepeatsCount int = 0
	for i := 0; i < len(genomeRepeats); i++ {
		if genomeRepeats[i].repeats > maxRepeatsCount {
			maxRepeatsCount = genomeRepeats[i].repeats
		}
	}

	for i := 0; i < len(genomeRepeats); i++ {
		if genomeRepeats[i].repeats == maxRepeatsCount {
			fmt.Print(genomeRepeats[i].genome + " ")
		}
	}
}
