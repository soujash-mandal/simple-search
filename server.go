package main

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/documents", documentsHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/autocomplete", autocompleteHandler)
	http.HandleFunc("/suggest", suggestHandler)
	http.HandleFunc("/phrase", phraseSearchHandler)
	http.HandleFunc("/analytics/top", analyticsHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
