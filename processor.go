package main

import (
	
	"strings"
)
func processor(sentence string) string {
	words := strings.Split(sentence, "\n")

	processed := []string{}
	for i := 0; i < len(words); i++ {
		words[i] = Case(words[i])
		words[i] = base(words[i])
		words[i] = articles(words[i])
		words[i] = punctuations(words[i])
		words[i] = quotes(words[i])
		words[i] = double(words[i])
		processed = append(processed, words[i])
	}
	return strings.Join(processed, "\n")
}
