# document_search

# Original Prompt:
The goal of this exercise is to create a working program to search a set of documents for the given search term or phrase (single token), and return results in order of relevance. 
Relevancy is defined as number of times the exact term or phrase appears in the document. 
Create three methods for searching the documents: 
Simple string matching
Text search using regular expressions
Preprocess the content and then search the index
Prompt the user to enter a search term and search method, execute the search, and return results. For instance:
 
Three files have been provided for you to read and use as sample search content.
Run a performance test that does 2M searches with random search terms, and measures execution time. Which approach is fastest? Why?
Provide some thoughts on what you would do on the software or hardware side to make this program scale to handle massive content and/or very large request volume (5000 requests/second or more)

# Performance testing

I have a branch where I've alrtered the output of the program to execute 20M random searches on a chosen method. Key differences in real searches being that the randomly generated "words" are not english, but expected performance with real data would yield similar results.

Results of running the performance testing with 2M searches:

String matching: 1m22.6363796s
Regex matching: 13m33.6419126s
Index Search: 2.0906441s

Note: For the regex matching I had a slight modification where I added a % complete in a parallel running perfomance test to the main one just to make sure the one without the percentage printing running would eventually return.

Regex takes the longest amount of time here by far. This is because of the function:

```golang
regexp.MustCompile(`(?:\A|\z|\s)(?i)` + search.term + `(?:\A|\z|\s)`)
```

This regex must compile 2 million times, one for each randomly generated search term, in order to perform the regex search method.

String searching came in second place. It was much more performant, but it still needed to search the each document for every randomly generated token.

Finally we have the indexed search.