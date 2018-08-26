package parser

import (
	"fmt"
)

// Parse parses given tokens
func Parse(tokens []string) {
	i := 0

	for i < len(tokens) {
		if tokens[i]+" "+tokens[i+1][0:6] == "print string" {
			fmt.Println(tokens[i+1][7:])
			i += 2
		}
	}
}
