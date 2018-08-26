package lexer

import (
	"regexp"
	"strings"
)

var tokens = []string{}                              // String array of parsed tokens
var exprOps = []string{"+", "-", "/", "*", "(", ")"} // String array of potential expression operators

// If 'a' is a string in 'list' then true is returned
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

	token := ""   // Temp. variable to store tokens
	isString := 0 // State to tell whether or not a string is being tokenized
	s := ""       // Temp. variable to store strings
	expr := ""    // Temp. variable to store expressions
	isExpr := 0   // State to tell whether or not an expression is being tokenized

	// Loop through all file contents and tokenize them
	for _, char := range contents {
		token += char
		isNumber, _ := regexp.MatchString("[0-9]", token) // Is true if the current token is a number (0-9)

		if token == " " {
			if isString == 0 {
				token = ""
			} else {
				token = " "
			}
		} else if token == "\r\n" || token == "\n" || token == "<EOF>" {
			if expr != "" && isExpr == 1 {
				tokens = append(tokens, "expr:"+expr)
				expr = ""
				isExpr = 0
			} else if expr != "" && isExpr == 0 {
				tokens = append(tokens, "num:"+expr)
				expr = ""
			}
			token = ""
		} else if token == "print" {
			tokens = append(tokens, "print")
			token = ""
		} else if isNumber && isString != 1 || stringInSlice(token, exprOps) {
			if stringInSlice(token, exprOps) {
				isExpr = 1
			}
			expr += token
			token = ""
		} else if token == "\"" {
			if isString == 0 {
				isString = 1
			} else if isString == 1 {
				tokens = append(tokens, "string:"+s+"\"")
				s = ""
				isString = 0
				token = ""
			}
		} else if isString == 1 {
			s += token
			token = ""
		}
	}
	//fmt.Println(tokens)
	return tokens
}
