package main

import (
	"fmt"

	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("Starting document search app")

	fmt.Println("\nEnter a search term or phrase (single token):")

	var search Search
	fmt.Scanln(&search.term)

	fmt.Println("\nEnter 1-3 for the following search methods:\n1. String Matching\n2. Regex Search\n3. Indexed Search")

	fmt.Scanln(&search.method)

	search.init("./sample_texts")

	search.executeSearch()

	fmt.Printf("\n%s Relevancy Results for '%s':\n", search.executionType, search.term)

	for _, text := range search.texts {
		fmt.Printf("%s's relevancy: %d\n", text.name, text.relevancy)
	}

	fmt.Printf("\nTotal relevancy across all texts: %d\n", search.totalRelevancy)

	fmt.Printf("Execution time: %v\n", search.executionTime)
}
