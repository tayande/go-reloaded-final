package main

import (
	"os"
	"strconv"
	"strings"
	"fmt"
)
func base(sentence string) string {
	words := strings.Fields(sentence)

	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(hex)":
			value, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error parsing value to decimal")
				os.Exit(1)
			}
			words[i-1] = strconv.FormatInt(value, 10)
			words = append(words[:i], words[i+1:]...)
			i--
		case "(bin)":
			value, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error parsing value to decimal")
			}
			words[i-1] = strconv.FormatInt(value, 10)
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}