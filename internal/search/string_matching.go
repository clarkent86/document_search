package search

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// response string for /api/string
type stringMatchResults struct {
	Results        map[string]int
	TotalRelevancy int
	ExecutionTime  time.Duration
}

// StringMatchSearch is the standalone string matching function
func (search *Search) StringMatchSearch(text *Text) {
	splitStrings := strings.Fields(text.content)
	for _, word := range splitStrings {
		if strings.EqualFold(word, search.Term) {
			text.Relevancy++
			search.TotalRelevancy++
		}
	}
}

// StringMatchSearchHandler is the API handler for string matching
func (search *Search) StringMatchSearchHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	results := make(map[string]int)
	for _, text := range search.Texts {
		splitStrings := strings.Fields(text.content)
		results[text.Name] = 0
		for _, word := range splitStrings {
			if strings.EqualFold(word, r.URL.Query().Get("term")) {
				results[text.Name]++
				search.TotalRelevancy++
			}
		}
	}

	json.NewEncoder(w).Encode(stringMatchResults{
		Results:        results,
		TotalRelevancy: search.TotalRelevancy,
		ExecutionTime:  time.Since(start),
	})
}
