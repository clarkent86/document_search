package main

import (
	"encoding/json"
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

	r := mux.NewRouter()

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	r.HandleFunc("/api/string", search.StringMatchSearchHandler)

	// TODO: Build out regex & index endpoints
	// r.HandleFunc("/api/regex", search.RegexMatchHandler)
	// r.HandleFunc("/api/index", search.IndexSearchHandler)

	// TODO: PUT endpoint to add new texts while app is running
	http.Handle("/", r)

	// TODO: Swagger documentation

	sugar.Fatal(http.ListenAndServe(":8080", r))
}
