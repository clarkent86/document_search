package search

import (
	"strings"
	"time"
)

type stringMatchResults struct {
	results        map[string]int
	TotalRelevancy int
	executionTime  time.Time
}

func (search *Search) StringMatchSearch(text *Text) {
	splitStrings := strings.Fields(text.content)
	for _, word := range splitStrings {
		if checkToken(word, search.Term) {
			text.Relevancy++
			search.TotalRelevancy++
		}
	}
}

func checkToken(token string, searchTerm string) bool {
	return strings.EqualFold(token, searchTerm)
}
