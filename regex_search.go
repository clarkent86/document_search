package main

func (search *Search) regexMatchSearch(text *Text) {
	matches := search.regexTerm.FindAllStringIndex(text.content, -1)
	text.relevancy = len(matches)
	search.totalRelevancy = search.totalRelevancy + len(matches)
}
