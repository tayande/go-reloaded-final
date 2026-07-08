package main

import (
	"testing"
)
func TestArticles(t *testing.T) {
	input := articles("There is no greater agony than bearing a untold story inside you.")
	output := articles("There is no greater agony than bearing an untold story inside you.")
	result := articles(input)
	if result != output {
		t.Errorf("got: %s\nexpected: %s\n", result, output)
	}
}
func TestBase(t *testing.T) {
	input := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	output := "Simply add 66 and 2 and you will see the result is 68."
	result := base(input)
	if result != output {
		t.Errorf("got: %s\nexpected: %s\n", result, output)
	}
}
func TestCase(t *testing.T) {
	input := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	output := "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."
	result := Case(input)
	if result != output {
		//t.Errorf("got: %s\nexpected: %s\n", result, output)
		t.Errorf("got: %q\n expected: %q\n", result, output)
	}
}
func TestPunctuation(t *testing.T) {
	input := "Punctuation tests are ... kinda boring ,what do you think ?"
	output := "Punctuation tests are... kinda boring, what do you think?"
	result := punctuations(input)
	if result != output {
		t.Errorf("got: %s\nexpected: %s\n", result, output)
	}
}
func TestQuotes(t *testing.T) {
	input := "I am exactly how they describe me: ' awesome '"
	output := "I am exactly how they describe me: 'awesome'"
	result := quotes(input)
	if result != output {
		t.Errorf("got: %s\nexpected: %s\n", result, output)
	}

}