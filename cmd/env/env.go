package main

import (
	"fmt"
	"regexp"
	"time"

	s "github.com/clarkent86/document_search/internal/search"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("Starting Environment Based Document Search App")

	var search s.Search

	err := search.InitEnv()
	if err != nil {
		sugar.Fatal(err.Error())
	}

	start := time.Now()
	for i := 0; i < len(search.Texts); i++ {
		switch methodChoice := search.Method; methodChoice {
		case 1:
			search.ExecutionType = "String Matching"
			search.StringMatchSearch(&search.Texts[i])
		case 2:
			search.ExecutionType = "Regex Matching"
			search.RegexTerm = regexp.MustCompile(`(?:\A|\z|\s)(?i)` + search.Term + `(?:\A|\z|\s)`)
			search.RegexMatchSearch(&search.Texts[i])
		case 3:
			search.ExecutionType = "Indexed Search"
			index := make(s.Index)
			countIndex := index.Add(&search.Texts[i])
			search.Texts[i].Relevancy = countIndex[search.Term]
			search.TotalRelevancy = search.TotalRelevancy + search.Texts[i].Relevancy
		}
		search.ExecutionTime = time.Since(start)
	}

	fmt.Printf("\n%s Relevancy Results for '%s':\n", search.ExecutionType, search.Term)

	for _, text := range search.Texts {
		fmt.Printf("%s's relevancy: %d\n", text.Name, text.Relevancy)
	}

	fmt.Printf("\nTotal relevancy across all texts: %d\n", search.TotalRelevancy)

	fmt.Printf("Execution time: %v\n", search.ExecutionTime)
}
