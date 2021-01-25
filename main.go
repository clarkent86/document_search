package main

import (
	"fmt"
	"time"

	"github.com/zhexuany/wordGenerator"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("Performance Testing Mode.")

	var search Search

	fmt.Println("Enter 1-3 for the following search methods:\n1. String Matching\n2. Regex Search\n3. Indexed Search")
	fmt.Scanln(&search.method)

	search.init("./sample_texts")

	start := time.Now()

	for i := 0; i < 2000000; i++ {
		search.term = wordGenerator.GetWord(20)
		search.executeSearch()
	}

	search.executionTime = time.Since(start)

	fmt.Printf("Execution time: %v", search.executionTime)
}
