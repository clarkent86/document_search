package main

import (
	"fmt"

	"go.uber.org/zap"
	s "github.com/clarkent86/document_search/internal/search"
)

var (
	logger *zap.SugaredLogger
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("Starting document search app")

	var search s.Search

	fmt.Println("\nEnter a search term or phrase (single token):")
	fmt.Scanln(&search.Term)

	fmt.Println("\nEnter 1-3 for the following search methods:\n1. String Matching\n2. Regex Search\n3. Indexed Search")
	fmt.Scanln(&search.Method)

	search.Init("./sample_texts")

	search.ExecuteSearch()

	fmt.Printf("\n%s Relevancy Results for '%s':\n", search.ExecutionType, search.Term)

	for _, text := range search.Texts {
		fmt.Printf("%s's relevancy: %d\n", text.Name, text.Relevancy)
	}

	fmt.Printf("\nTotal relevancy across all texts: %d\n", search.TotalRelevancy)

	fmt.Printf("Execution time: %v\n", search.ExecutionTime)
}
