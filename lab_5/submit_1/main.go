package main

import (
	"fmt"
	"os"
	"strings"
)

var RNA_codon_table = map[string]string{
	"UUU": "Phe", "UCU": "Ser", "UAU": "Tyr", "UGU": "Cys",
	"UUC": "Phe", "UCC": "Ser", "UAC": "Tyr", "UGC": "Cys",
	"UUA": "Leu", "UCA": "Ser", "UAA": "---", "UGA": "---",
	"UUG": "Leu", "UCG": "Ser", "UAG": "---", "UGG": "Trp",

	"CUU": "Leu", "CCU": "Pro", "CAU": "His", "CGU": "Arg",
	"CUC": "Leu", "CCC": "Pro", "CAC": "His", "CGC": "Arg",
	"CUA": "Leu", "CCA": "Pro", "CAA": "Gln", "CGA": "Arg",
	"CUG": "Leu", "CCG": "Pro", "CAG": "Gln", "CGG": "Arg",

	"AUU": "Ile", "ACU": "Thr", "AAU": "Asn", "AGU": "Ser",
	"AUC": "Ile", "ACC": "Thr", "AAC": "Asn", "AGC": "Ser",
	"AUA": "Ile", "ACA": "Thr", "AAA": "Lys", "AGA": "Arg",
	"AUG": "Met", "ACG": "Thr", "AAG": "Lys", "AGG": "Arg",

	"GUU": "Val", "GCU": "Ala", "GAU": "Asp", "GGU": "Gly",
	"GUC": "Val", "GCC": "Ala", "GAC": "Asp", "GGC": "Gly",
	"GUA": "Val", "GCA": "Ala", "GAA": "Glu", "GGA": "Gly",
	"GUG": "Val", "GCG": "Ala", "GAG": "Glu", "GGG": "Gly",
}

var singleletter = map[string]string{
	"Cys": "C", "Asp": "D", "Ser": "S", "Gln": "Q", "Lys": "K",
	"Trp": "W", "Asn": "N", "Pro": "P", "Thr": "T", "Phe": "F", "Ala": "A",
	"Gly": "G", "Ile": "I", "Leu": "L", "His": "H", "Arg": "R", "Met": "M",
	"Val": "V", "Glu": "E", "Tyr": "Y", "---": "*",
}

var REVERSABLE = map[string]string{
	"A": "T",
	"T": "A",
	"G": "C",
	"C": "G",
}

func main() {
	var (
		rna          string
		pattern      string
		patternWords []string
	)

	fmt.Fscanf(os.Stdin, "%s", &rna)
	fmt.Fscanf(os.Stdin, "%s", &pattern)

	rna = strings.Replace(rna, "T", "U", -1)

	for i := 0; i < len(pattern); i++ {
		patternWords = append(patternWords, string(pattern[i]))
	}

	for i := 0; i < len(rna)-len(patternWords)*3+1; i++ {
		var resWord string = ""
		var word string = rna[i : i+len(patternWords)*3]
		for j := 0; j < len(patternWords); j++ {
			if singleletter[RNA_codon_table[string(word[j*3:j*3+3])]] == patternWords[j] {
				resWord += word[j*3 : j*3+3]
			} else {
				resWord = ""
				break
			}
		}

		if resWord == "" {
			var backConvert = strings.Replace(word, "U", "T", -1)
			var reversedWord string = ""
			for j := len(backConvert) - 1; j >= 0; j-- {
				reversedWord += REVERSABLE[string(backConvert[j])]
			}
			reversedWord = strings.Replace(reversedWord, "T", "U", -1)
			for j := 0; j < len(patternWords); j++ {
				if singleletter[RNA_codon_table[string(reversedWord[j*3:j*3+3])]] == patternWords[j] {
					resWord += word[j*3 : j*3+3]
				} else {
					resWord = ""
					break
				}
			}
		}

		if resWord != "" {
			fmt.Print(strings.Replace(resWord, "U", "T", -1))
			fmt.Println()
		}
	}
}
