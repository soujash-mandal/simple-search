package main

import "sort"

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

func (s *SearchEngine) AddDocument(doc Document) {
	s.Documents[doc.ID] = doc
	titleTokens := tokenize(doc.Title)
	contentTokens := tokenize(doc.Content)
	tokens := append(titleTokens, contentTokens...)
	for _, token := range tokens {
		if _, exists := s.Index[token]; !exists {
			s.Index[token] = make(map[int]struct{})
		}
		s.Index[token][doc.ID] = struct{}{}
	}
}


func (s *SearchEngine) Search(query string) []Document {
	tokens := tokenize(query)
	if len(tokens) == 0 {
		return nil
	}
	scores := make(map[int]int)
	for _, token := range tokens {
		for id := range s.Index[token] {
			scores[id]++
		}
	}
	if len(scores) == 0 {
		return nil
	}
	ids := make([]int, 0, len(scores))
	for id := range scores {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool {
		return scores[ids[i]] > scores[ids[j]]
	})
	results := make([]Document, len(ids))
	for i, id := range ids {
		results[i] = s.Documents[id]
	}
	return results
}