package main

import (
	"strings"
)
func quotes(sentence string) string {
	words := strings.Split(sentence, "'")

	for i := 1; i < len(words); i+=2 {
		words[i] = strings.TrimSpace(words[i])
	}
	return strings.Join(words, "'")
}
func double(sentence string) string {
	words := strings.Split(sentence, `"`)

	for i := 1; i < len(words); i+=2 {
		words[i] = strings.TrimSpace(words[i])
	}
	return strings.Join(words, `"`)
}

//qotes using the regexp package 

/*func here(s string)string{
	re := regexp.MustCompile(`'\s+(.*?)\s+'`)
	s = re.ReplaceAllString(s, "'$1'")
	return s
}

*/