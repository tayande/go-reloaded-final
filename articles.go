package main

import (

	"strings"
)
func articles(sentence string) string {
	words := strings.Fields(sentence)

	for i := 0; i < len(words)-1; i++ {
		vowel := strings.ContainsAny("AEIOUHaeiouh", string(words[i+1][0]))

		switch words[i] {
		case "A":
			if vowel {
				words[i] = "An"
			}
		case "a":
			if vowel {
				words[i] = "an"
			}
		case "An":
			if !vowel {
				words[i] = "A"
			}
		case "an":
			if !vowel {
				words[i] = "a"
			}
		}
	}
	return strings.Join(words, " ")
}