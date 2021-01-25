package main

import (
	"fmt"
	"regexp"
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

	fmt.Println("\nEnter 1-3 for the following search methods:\n1. String Matching\n2. Regex Search\n3. Indexed Search")
	fmt.Scanln(&search.method)

	search.init("./sample_texts")

	start := time.Now()

	switch methodChoice := search.method; methodChoice {
	case "1":
		for i := 0; i < 2000000; i++ {
			search.term = wordGenerator.GetWord(20)
			search.executeSearch()
		}
	case "2":
		for i := 0; i < 2000000; i++ {
			search.term = wordGenerator.GetWord(20)
			search.regexTerm = regexp.MustCompile(`(?:\A|\z|\s)(?i)` + search.term + `(?:\A|\z|\s)`)
			search.executeSearch()
		}
	case "3":
		search.executeSearch()
		for i := 0; i < 2000000; i++ {
			search.term = wordGenerator.GetWord(20)
			for p := 0; p < len(search.texts); p++ {
				search.texts[p].relevancy = search.countIndex[search.term]
			}
		}
	default:
		fmt.Println("A valid search method was not detected. Please enter an int 1-3.")
	}

	search.executionTime = time.Since(start)

	fmt.Printf("Execution time: %v\n", search.executionTime)
}
