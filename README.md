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

# Running document_search

Requirements:
- go version go1.11.6

Usage:
```bash
go run .
```

# Performance testing

I've taken results from [my performanceTesting branch](https://github.com/clarkent86/document_search/tree/performanceTesting) where I've alrtered the output of the program to execute 20M random searches on a chosen method. The key differences in real searches being that the randomly generated "words" are not english, but expected performance with real data would yield similar results.

Results of running the performance testing with 2M searches:

|  Search Method  |      3 character strings      |  20 character strings |
|----------|:-------------:|------:|
| String matching: |  1m44.4954022s | 1m49.5259472s |
| Regex matching: |   13m58.9412604s    |   13m33.6419126s |
|Index Search: | 319.6756ms |   2.0906441s |

 
  (it helped to do do a modulus derived printout of the randomly generated terms to verify progress was made)
 

Regex takes the longest amount of time here by far. This is because of the function:

```golang
regexp.MustCompile(`(?:\A|\z|\s)(?i)` + search.term + `(?:\A|\z|\s)`)
```

This regex must compile 2 million times, one for each randomly generated search term, in order to perform the regex search method.

String searching came in second place. It was much more performant, but it still needed to search each document for every randomly generated token.

Finally we have the indexed search. Since we only needed to index each document once it became a very quick map lookup for each randomly generated token vs a more expensive string match or even more regex compilation & match. Indexing is clearly the more performant search when it comes to a large number of random terms.

# Scaling

In order to scale this solution for massive content and/or very large request volumes I would continue to use the indexed search.

## Scaling for massive content

### Processing Tokens
In order to scale for massive content I imagine we are considering a lot more than 3 sample texts to search across. I would stick with my index based search method and move twoards splitting my application into microservices. With a microservice that will handle all of my index based search tokenization adds I can then queue my add requests with a queueing tool and put the adds behind a load balancer. Ultimately the add service will upload it's finished tokeziation to a backend. I could then move my search fucntion to a separate microservice that will query that backend.

### Processing Queries
This doesn't necessarily counteract the massive amount of tokens in this backend. Another thing I may need to scale for is how the tokens are stored. Based on many different keys, I could hash the data to get tokenized in a predictive way in order for the query microservice to be able to target specific backends in the same way. This way different backend requests will balance the load. Both the backend and the query microservice will need to be able to derive the way these tokens are stored in order for this to work.

## Scaling for massive amounts of requests
If I'm implemented the other improvements scaling for more requests is now probably bottlenecked on the query processing micro-service. Using a load balancer infront of the query microservice I would be able to scale that service up to handle as many queries as the backend can take. If I then notice the bottleneck being the backend, we could mirror the backend and but a load balancer between the query and backend as well.