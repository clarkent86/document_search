package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
)

// Search is struct to contain the files and their contents
type Search struct {
	executionTime  time.Duration
	executionType  string
	method         string
	term           string
	texts          []Text
	totalRelevancy int
	regexTerm      *regexp.Regexp
	countIndex     map[string]int
}

// Text is a struct to contain individual document metrics
type Text struct {
	content   string
	id        int
	name      string
	path      string
	relevancy int
	tokens    []string
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

func (search *Search) executeSearch() {
	index := make(index)
	search.regexTerm = regexp.MustCompile(`(?:\A|\z|\s)(?i)` + search.term + `(?:\A|\z|\s)`)
	start := time.Now()
	for i := 0; i < len(search.texts); i++ {
		search.texts[i].id = i
		switch methodChoice := search.method; methodChoice {
		case "1":
			search.stringMatchSearch(&search.texts[i])
		case "2":
			search.regexMatchSearch(&search.texts[i])
		case "3":
			search.countIndex = index.add(&search.texts[i])
		default:
			fmt.Println("A valid search method was not detected. Please enter an int 1-3.")
		}
		search.executionTime = time.Since(start)
	}
}
