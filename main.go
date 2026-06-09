package main

import "fmt"

func main() {
	engine := NewSearchEngine()

	engine.AddDocument(Document{
		ID:      1,
		Title:   "Redis Guide",
		Content: "Redis is an in memory database",
	})

	engine.AddDocument(Document{
		ID:      2,
		Title:   "MongoDB Basics",
		Content: "MongoDB is a document database",
	})

	results := engine.Search("redis database")
	fmt.Println("Results:", results)
}
