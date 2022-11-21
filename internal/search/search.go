package search

import (
	"os"
	"regexp"
	"time"
)

// Search is struct to contain the files and their contents
type Search struct {
	countIndex     map[string]int
	ExecutionTime  time.Duration
	ExecutionType  string
	Method         int
	Term           string
	Texts          []Text
	TotalRelevancy int
	RegexTerm      *regexp.Regexp
}

// Text is a struct to contain individual document metrics
type Text struct {
	content   string
	Name      string
	path      string
	Relevancy int
}

func (search *Search) Init(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
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
