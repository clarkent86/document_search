package search

import (
	"strings"
	"unicode"
)

// Index is a text's associated map of strings and their count
type Index map[string][]int

// Add will process a text and resolve it's index
func (index Index) Add(text *Text) map[string]int {
	var countIndex = make(map[string]int)
	for _, word := range analyze(text.content) {
		_, ok := countIndex[word]
		if ok {
			countIndex[word] = countIndex[word] + 1
		} else {
			countIndex[word] = 1
		}
	}
	return countIndex
}

// analyze is a helper function to normalize a text's strings
func analyze(content string) []string {
	words := tokenize(content)
	words = toLowerFilter(words)
	return words
}

// toLowerFilter is a helper function to lower case a string
func toLowerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

// tokenize is a helper function to process a texts content including special characters
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		// split on unexpected word chacters (space, newline, etc)
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsPunct(r)
	})
}
