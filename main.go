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

	var texts Texts

	texts.init("./sample_texts")

	fmt.Println(texts)
}
