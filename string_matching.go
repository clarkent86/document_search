package main

import (
	"strings"
)

func (search *Search) stringMatchSearch(text *Text) {
	splitStrings := strings.Split(text.content, " ")
	for _, word := range splitStrings {
		if strings.EqualFold(word, search.term) {
			text.relevancy++
			search.totalRelevancy++
		}
	}
}
