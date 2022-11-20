package main

import (
	"fmt"

	s "github.com/clarkent86/document_search/internal/search"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("Starting document search app")

	var search s.Search

	err := search.InitEnv()
	if err != nil {
		sugar.Fatal(err.Error())
	}

	search.ExecuteSearch()

	fmt.Printf("\n%sRelevancy Results for '%s':\n", search.ExecutionType, search.Term)

	for _, text := range search.Texts {
		fmt.Printf("%s's relevancy: %d\n", text.Name, text.Relevancy)
	}

	fmt.Printf("\nTotal relevancy across all texts: %d\n", search.TotalRelevancy)

	fmt.Printf("Execution time: %v\n", search.ExecutionTime)
}
