package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)
func title(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
}
func Case(sentence string) string {
	words := strings.Fields(sentence)

	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		case "(up,":
			words[i+1] = strings.Trim(words[i+1], ")")
			value, err := strconv.Atoi(words[i+1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error converting number to integer")
				os.Exit(1)
			}
			for j := 1; j <= value; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		case "(low,":
			words[i+1] = strings.Trim(words[i+1], ")")
			value, err := strconv.Atoi(words[i+1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error converting number to integer")
				os.Exit(1)
			}
			for j := 1; j <= value; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		case "(cap)":
			words[i-1] = title(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		case "(cap,":
			words[i+1] = strings.Trim(words[i+1], ")")
			value, err := strconv.Atoi(words[i+1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error converting number to integer")
				os.Exit(1)
			}
			for j := 1; j <= value; j++ {
				words[i-j] = title(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}
	return punctuations(strings.Join(words, " "))
}

