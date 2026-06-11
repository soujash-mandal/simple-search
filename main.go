package main

import "log"

func main() {
	analytics, err := LoadAnalytics()
	if err == nil {
		engine.QueryCounts = analytics
	}

	docs, err := LoadDocuments()
	if err != nil {
		log.Println("No existing documents found")
	} else {
		for _, doc := range docs {
			engine.AddDocument(doc)
		}
	}

	StartServer()
}
