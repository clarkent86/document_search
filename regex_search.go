package main

import (
	"regexp"
)

func (search *Search) regexMatchSearch(text *Text) {
	// term := regexp.MustCompile("(?i)" + search.term)
	term := regexp.MustCompile(`(?:\A|\z|\s|\"|\[)(?i)` + search.term + `(?:\A|\z|\s|\"|\])`)
	// (?:\A|\z|\s)
	matches := term.FindAllStringIndex(text.content, -1)
	text.relevancy = len(matches)
	search.totalRelevancy = search.totalRelevancy + len(matches)
}
