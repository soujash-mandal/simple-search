package main

import (
	"encoding/json"
	"net/http"
)

var engine = NewSearchEngine()

func getDocumentsHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		engine.GetAllDocuments(),
	)
}

func addDocumentHandler(w http.ResponseWriter, r *http.Request) {
	var doc Document
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	if err := SaveDocumentMongo(doc); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	engine.AddDocument(doc)
	w.WriteHeader(http.StatusCreated)
}

func documentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addDocumentHandler(w, r)
	case http.MethodGet:
		getDocumentsHandler(w)
	default:
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	engine.RecordQuery(query)
	_ = SaveAnalytics(engine.QueryCounts)
	results := engine.Search(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	prefix := r.URL.Query().Get("q")
	results := engine.AutoComplete(prefix)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func suggestHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results := engine.Suggest(query, 2)
	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	json.NewEncoder(w).Encode(
		results,
	)
}

func phraseSearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results := engine.PhraseSearch(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func analyticsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(engine.TopQueries(10))
}
