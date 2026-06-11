package main

import (
	"log"
)

func main() {
	err := ConnectMongo("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")

	docs, err := LoadDocumentsMongo()
	if err != nil {
		log.Fatal(err)
	}

	for _, doc := range docs {
		engine.AddDocument(doc)
	}

	log.Printf(
		"Loaded %d documents into search index",
		len(docs),
	)

	analytics, err := LoadAnalytics()
	if err == nil {
		engine.QueryCounts = analytics
	}

	StartServer()
}