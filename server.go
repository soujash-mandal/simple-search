package main

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/documents", addDocumentHandler)
	http.HandleFunc("/search", searchHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}