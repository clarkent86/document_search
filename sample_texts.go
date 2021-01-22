package main

import (
	"io/ioutil"
	"log"
	"os"
)

// SampleTexts is struct to contain the files and their contents
type SampleTexts struct {
	directory string
	files     []string
	contents  []string
}

func (sampleTexts *SampleTexts) init(path string) error {

	sampleTexts.directory = path

	file, err := os.Open(sampleTexts.directory)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	files, _ := file.Readdirnames(0) // 0 to read all files and folders

	contents := make([]string, 0)

	for _, name := range files {
		content, _ := readInFile(path, name)
		contents = append(contents, content)
	}

	sampleTexts.files = files
	sampleTexts.contents = contents

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
