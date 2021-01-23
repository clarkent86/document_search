package main

import (
	"strings"
	"time"
)

func (search *Search) stringMatchSearch(term string) {

	start := time.Now()

	for i, text := range search.texts {
		splitStrings := strings.Split(text.content, " ")
		for _, word := range splitStrings {
			if word == term {
				search.texts[i].relevancy++
				search.total_relevancy++
			}
		}
	}

	executionTime := time.Since(start)

	search.executionTime = executionTime

}
