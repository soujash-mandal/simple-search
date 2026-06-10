package main

import "log"

func main() {
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