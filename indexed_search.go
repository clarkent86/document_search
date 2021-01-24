package main

import (
	"strings"
	"unicode"
)

type index map[string][]int

func (index index) add(text Text) map[string]int {
	var countIndex = make(map[string]int)
	for _, word := range analyze(text.content) {
		ids := index[word]
		if ids != nil && ids[len(ids)-1] == text.id {
			// Don't add same ID twice & check if counter has been initialized
			_, ok := countIndex[word]
			if ok {
				countIndex[word] = countIndex[word] + 1
			} else {
				countIndex[word] = 1
				continue
			}
		}
		index[word] = append(ids, text.id)
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
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func (idx index) search(text string) [][]int {
	var r [][]int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			r = append(r, ids)
		}
	}
	return r
}
