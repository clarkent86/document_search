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

	var sampleTexts SampleTexts

	sampleTexts.init("./sample_text")

	fmt.Println(sampleTexts)
}
