package main

import (
	"fmt"
	"os"
)
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run . input.txt output.txt")
		os.Exit(2)
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if inputFile == outputFile {
		fmt.Fprintln(os.Stderr, "Cannot use the same input file as output file")
		os.Exit(1)
	}

	info, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error. File does not exist")
		os.Exit(1)
	}
	if len(info) == 0 {
		fmt.Fprintln(os.Stderr, "File is empty")
		os.Exit(1)
	}
	data := string(info)
	processed := processor(data)
	newProcessed := []byte(processed)

	err = os.WriteFile(outputFile, newProcessed, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing into output file")
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Done. Output written into %s\n", outputFile)

}