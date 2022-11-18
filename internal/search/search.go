package search

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
	countIndex     map[string]int
	ExecutionTime  time.Duration
	ExecutionType  string
	Method         string
	Term           string
	Texts          []Text
	TotalRelevancy int
	regexTerm      *regexp.Regexp
}

// Text is a struct to contain individual document metrics
type Text struct {
	content   string
	Name      string
	path      string
	Relevancy int
	tokens    []string
}

func (search *Search) Init(path string) error {

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
		text.Name = name
		search.Texts = append(search.Texts, text)
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

func (search *Search) ExecuteSearch() {
	index := make(index)
	search.regexTerm = regexp.MustCompile(`(?:\A|\z|\s)(?i)` + search.Term + `(?:\A|\z|\s)`)
	start := time.Now()
	for i := 0; i < len(search.Texts); i++ {
		switch methodChoice := search.Method; methodChoice {
		case "1":
			search.stringMatchSearch(&search.Texts[i])
		case "2":
			search.regexMatchSearch(&search.Texts[i])
		case "3":
			countIndex := index.add(&search.Texts[i])
			search.Texts[i].Relevancy = countIndex[search.Term]
			search.TotalRelevancy = search.TotalRelevancy + search.Texts[i].Relevancy
		default:
			fmt.Println("A valid search method was not detected. Please enter an int 1-3.")
		}
		search.ExecutionTime = time.Since(start)
	}
}
