package search

func (search *Search) RegexMatchSearch(text *Text) {
	matches := search.RegexTerm.FindAllStringIndex(text.content, -1)
	text.Relevancy = len(matches)
	search.TotalRelevancy = search.TotalRelevancy + len(matches)
}
