package parser

import (
	"fmt"
)

// Prints given string and removes anything added by tokenization
func doPrint(toPrint string) {
	if toPrint[0:6] == "string" {
		toPrint = toPrint[8 : len(toPrint)-1]
	} else if toPrint[0:3] == "num" {
		toPrint = toPrint[4:]
	} else if toPrint[0:4] == "expr" {
		toPrint = toPrint[5:]
	}
	fmt.Println(toPrint)
}

// Supposed to eval the numbers in a string; still haven't figured this out
func eval(expr string) {

}

// Parse parses given tokens
func Parse(tokens []string) {
	i := 0

	// Checks for print command and prints accordingly
	for i < len(tokens) {
		if tokens[i] == "print" {
			if tokens[i+1][0:6] == "string" {
				doPrint(tokens[i+1])
			} else if tokens[i+1][0:3] == "num" {
				doPrint(tokens[i+1])
			} else if tokens[i+1][0:4] == "expr" {
				doPrint(tokens[i+1])
				eval(tokens[i+1])
			}
			i += 2
		}
	}
}
