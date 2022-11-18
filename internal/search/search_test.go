package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteSearch(t *testing.T) {
	tests := []struct {
		name           string
		search         Search
		text           Text
		searchTerm     string
		expectedResult int
	}{{
		name: "simple",
		search: Search{
			term:  "the",
			texts: []Text{{content: `the`}},
		},
		expectedResult: 1,
	}, {
		name: "case insensitive",
		search: Search{
			term:  "the",
			texts: []Text{{content: `The`}},
		},
		expectedResult: 1,
	}, {
		name: "in the midde of another word",
		search: Search{
			term:  "the",
			texts: []Text{{content: `other`}},
		},
		expectedResult: 0,
	}, {
		name: "beginning another word",
		search: Search{
			term:  "the",
			texts: []Text{{content: `there`}},
		},
		expectedResult: 0,
	}, {
		name: "end of another word",
		search: Search{
			term:  "the",
			texts: []Text{{content: `breathe`}},
		},
		expectedResult: 0,
	}, { // the following tests are scenarios I might clarify with a customer/product owner on desired functionality. Since the case study defines exact documentContent matching I left these out but would be happy to talk through stategies to support matching these cases.

		name: "ends with .",
		search: Search{
			term:  "the",
			texts: []Text{{content: `the.`}},
		},
		expectedResult: 0,
	}, {
		name: "ends with ,",
		search: Search{
			term:  "the",
			texts: []Text{{content: `the.`}},
		},
		expectedResult: 0,
	}, {
		name: "starting with quotes",
		search: Search{
			term:  "the",
			texts: []Text{{content: `"the`}},
		},
		expectedResult: 0,
	}, {
		name: "ending with quotes",
		search: Search{
			term:  "the",
			texts: []Text{{content: `the"`}},
		},
		expectedResult: 0,
	}, {
		name: "starting with (",
		search: Search{
			term:  "the",
			texts: []Text{{content: `(the`}},
		},
		expectedResult: 0,
	}, {
		name: "ending with )",
		search: Search{
			term:  "the",
			texts: []Text{{content: `the)`}},
		},
		expectedResult: 0,
	}, {
		name: "starting with [",
		search: Search{
			term:  "the",
			texts: []Text{{content: `[the`}},
		},
		expectedResult: 0,
	}, {
		name: "ending with ]",
		search: Search{
			term:  "the",
			texts: []Text{{content: `the]`}},
		},
		expectedResult: 0,
	}, {
		name: "starting with {",
		search: Search{
			term:  "the",
			texts: []Text{{content: `{the`}},
		},
		expectedResult: 0,
	}, {
		name: "ending with }",
		search: Search{
			term:  "the",
			texts: []Text{{content: `the}`}},
		},
		expectedResult: 0,
	}, {
		name: "starting with ellipses",
		search: Search{
			term:  "the",
			texts: []Text{{content: `...the`}},
		},
		expectedResult: 0,
	}, {
		name: "ending with ellipses",
		search: Search{
			term:  "the",
			texts: []Text{{content: `the...`}},
		},
		expectedResult: 0,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//tests string matching search
			test.search.stringMatchSearch(&test.search.Texts[0])
			assert.Equal(t, test.expectedResult, test.search.Texts[0].relevancy)

			//tests regex search
			test.search.regexMatchSearch(&test.search.Texts[0])
			assert.Equal(t, test.expectedResult, test.search.Texts[0].relevancy)

			//tests indexed search
			index := make(index)
			test.search.texts[0].id = 0
			countIndex := index.add(&test.search.Texts[0])
			fmt.Println(countIndex[test.search.Term])
			assert.Equal(t, test.expectedResult, countIndex[test.search.Term])
		})
	}
}
