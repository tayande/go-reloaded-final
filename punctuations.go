package main

import (
    "regexp"
)
func punctuations(word string) string {
    result := regexp.MustCompile(`\s+([,.?:;!])`)
    word = result.ReplaceAllString(word, "$1")

    result1 := regexp.MustCompile(`([,.?;:!])(\s*)(\w)`)
    word = result1.ReplaceAllString(word, "$1 $3")

    return word
}