package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Search is struct to contain the files and their contents
type Search struct {
	executionTime   time.Duration
	executionType   string
	texts           []Text
	total_relevancy int
}

type Text struct {
	content   string
	name      string
	path      string
	relevancy int
}

func (search *Search) init(path string) error {

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	files, _ := file.Readdirnames(0) // 0 to read all files and folders

	var text Text

	for _, name := range files {
		text.path = path + "/" + name
		text.content, _ = readInFile(text.path)
		text.name = name
		search.texts = append(search.texts, text)
	}

	return nil
}

func readInFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content), err
}
