package main

import (
	"strings"
)

func (search *Search) stringMatchSearch(text *Text) {
	splitStrings := strings.Fields(text.content)
	for _, word := range splitStrings {
		if checkToken(word, search.term) {
			text.relevancy++
			search.totalRelevancy++
		}
	}
}

func checkToken(token string, searchTerm string) bool {
	return strings.EqualFold(token, searchTerm)
}
