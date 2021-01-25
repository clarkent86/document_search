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

	search.method = "3"

	search.init("./sample_texts")

	start := time.Now()

	for i := 0; i < 2000000; i++ {
		if i%20000 == 0 {
			fmt.Printf("%d%%", i/20000)
		}
		search.term = wordGenerator.GetWord(20)
		search.executeSearch()
	}

	search.executionTime = time.Since(start)

	fmt.Printf("Execution time: %v\n", search.executionTime)
}
