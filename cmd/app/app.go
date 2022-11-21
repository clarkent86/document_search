package main

import (
	"net/http"

	s "github.com/clarkent86/document_search/internal/search"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("Starting Gorilla/Mux document search app")

	var search s.Search

	search.Init("./sample_texts")
	// for i := 0; i < len(search.Texts); i++ {
	// 	index := make(s.Index)
	// 	countIndex := index.Add(&search.Texts[i])
	// }

	r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/string", search.StringMatchSearchHandler)
	// r.HandleFunc("/regex", search.RegexMatchSearch())
	// r.HandleFunc("/index", search.IndexSearch())
	http.Handle("/", r)

	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "127.0.0.1:8000",
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	sugar.Fatal(http.ListenAndServe(":8080", r))
}
