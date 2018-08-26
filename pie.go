package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	lexer "../pie/lexer"
	parser "../pie/parser"
)

// Reads the entire given file to memory
func openFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("File '%s' could not be found\n", filename)
		os.Exit(0)
	}
	return b
}

// Prints time elapsed in milliseconds
func timer(endTime int64) {
	begin := time.Now().UnixNano()
	diff := (begin - endTime) / 1000000
	fmt.Printf("Compiled in %d millisecond(s)\n", diff)
}

func main() {
	var verbose bool // Boolean used for flags

	fileFlag := flag.String("f", "", "Specify a .pie file to compile")     // Look for -f flag
	flag.BoolVar(&verbose, "v", false, "See specific compilation details") // If -v flag is given set verbose to true
	flag.Parse()                                                           // Read all given flags

	// If now file is given print message and exit
	if *fileFlag == "" {
		fmt.Println("Please specify a .pie file to compile")
		os.Exit(0)
	}

	// If -v flag is given do stuff
	if verbose {
		defer timer(time.Now().UnixNano()) // Wait till lexer & parser are finished to send endTime
	}

	data := openFile(*fileFlag) // Load specified file to memory
	tokens := lexer.Lex(data)   // Send file bytes to lexer
	parser.Parse(tokens)        // Send tokens to parser
}
