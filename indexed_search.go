package main

import (
	"strings"
	"unicode"
)

type index map[string][]int

func (index index) add(text *Text) map[string]int {
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

func analyze(content string) []string {
	words := tokenize(content)
	words = toLowerFilter(words)
	return words
}

func toLowerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		// split on unexpected word chacters (space, newline, etc)
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsPunct(r)
	})
}
