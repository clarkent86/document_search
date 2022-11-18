package search

func (search *Search) regexMatchSearch(text *Text) {
	matches := search.regexTerm.FindAllStringIndex(text.content, -1)
	text.Relevancy = len(matches)
	search.TotalRelevancy = search.TotalRelevancy + len(matches)
}
