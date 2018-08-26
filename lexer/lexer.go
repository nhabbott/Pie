package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

var tokens = []string{}                    // String array of parsed tokens
var exprOps = []string{"+", "-", "/", "*"} // String array of potential expression operators

// If a is a string in list then true is returned
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Lex accepts a file ([]byte) and tokenizes it
func Lex(file []byte) []string {
	contents := strings.Split(string(file), "") // Array of the file's individual characters
	contents = append(contents, "<EOF>")        // Add <EOF> to the end of the file

	token := "" // Temp. variable to store tokens
	state := 0  // State to tell whether or not a string is being lex'd
	s := ""     // Temp. variable to store strings
	expr := ""  // Temp. variable to store expressions

	// Loop through all file contents and tokenize them
	for _, char := range contents {
		token += char
		isNumber, _ := regexp.MatchString("[0-9]", token) // Is true if the current token is a number (0-9)

		if token == " " {
			if state == 0 {
				token = ""
			} else {
				token = " "
			}
		} else if token == "\r\n" || token == "\n" || token == "<EOF>" {
			if expr != "" {
				fmt.Println(expr)
				expr = ""
			}
			token = ""
		} else if token == "print" {
			tokens = append(tokens, "print")
			token = ""
		} else if isNumber && state != 1 || stringInSlice(token, exprOps) {
			expr += token
			token = ""
		} else if token == "\"" {
			if state == 0 {
				state = 1
			} else if state == 1 {
				tokens = append(tokens, "string:"+s+"\"")
				s = ""
				state = 0
				token = ""
			}
		} else if state == 1 {
			s += token
			token = ""
		}
	}
	//fmt.Println(tokens)

	return tokens
}
