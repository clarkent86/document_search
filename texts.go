package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Texts is struct to contain the files and their contents
type Texts struct {
	directory       string
	executionTime   time.Time
	executionType   string
	files           []Document
	total_relevancy int
}

type Document struct {
	path      string
	content   string
	relevancy int
}

func (Texts *Texts) init(path string) error {

	Texts.directory = path

	file, err := os.Open(Texts.directory)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	files, _ := file.Readdirnames(0) // 0 to read all files and folders

	var document Document

	for _, name := range files {
		document.path = path + "/" + name
		document.content, _ = readInFile(path, name)
		Texts.files = append(Texts.files, document)
	}

	return nil
}

func readInFile(directory string, file string) (string, error) {
	content, err := ioutil.ReadFile(directory + "/" + file)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	return text, err
}
