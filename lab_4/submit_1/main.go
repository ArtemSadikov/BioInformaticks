package main

import (
	"fmt"
	"os"
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

func main() {
	var RNA string
	fmt.Fscanf(os.Stdin, "%s\n", &RNA)
	for i := 0; i < len(RNA); i += 3 {
		var res = string(singleletter[RNA_codon_table[RNA[i:i+3]]])
		if res != "*" {
			fmt.Printf(res)
		}
	}
}
