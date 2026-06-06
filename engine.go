package main

type SearchEngine struct {
	Documents map[int]Document
	Index     map[string]map[int]struct{}
}

func NewSearchEngine() *SearchEngine {
	return &SearchEngine{
		Documents: make(map[int]Document),
		Index:     make(map[string]map[int]struct{}),
	}
}