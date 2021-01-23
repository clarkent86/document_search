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
	var term string
	fmt.Scanln(&term)

	fmt.Println("\nEnter 1-3 for the following search methods:\n1. String Matching\n2. Regex Search\n3. Indexed Search")
	var method int
	fmt.Scanln(&method)

	var search Search

	search.init("./sample_texts")

	switch methodChoice := method; methodChoice {
	case 1:
		search.executionType = "String Matching"
		search.stringMatchSearch(term)
	default:
		fmt.Println("A valid search method was not detected. Please enter an int 1-3.")
	}

	fmt.Printf("\n%s Relevancy Results for '%s':\n", search.executionType, term)

	for _, text := range search.texts {
		fmt.Printf("%s's relevancy: %d\n", text.name, text.relevancy)
	}

	fmt.Printf("\nTotal relevancy across all texts: %d\n", search.total_relevancy)

	fmt.Printf("Execution time: %v\n", search.executionTime)
}
