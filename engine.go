package main

import "sort"

type SearchEngine struct {
	Documents map[int]Document
	Index     map[string]map[int]int
}

func NewSearchEngine() *SearchEngine {
	return &SearchEngine{
		Documents: make(map[int]Document),
		Index:     make(map[string]map[int]int),
	}
}

func (s *SearchEngine) AddDocument(doc Document) {
	s.Documents[doc.ID] = doc
	titleTokens := tokenize(doc.Title)
	contentTokens := tokenize(doc.Content)
	tokens := append(titleTokens, contentTokens...)
	for _, token := range tokens {
		if _, exists := s.Index[token]; !exists {
			s.Index[token] = make(map[int]int)
		}
		s.Index[token][doc.ID]++
	}
}

func (s *SearchEngine) Search(query string) []Document {
	tokens := tokenize(query)
	if len(tokens) == 0 {
		return nil
	}
	scores := make(map[int]int)
	for _, token := range tokens {
		docFreqs, exists := s.Index[token]
		if !exists {
			continue
		}
		for id, tf := range docFreqs {
			scores[id] += tf
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
	results := make([]Document, 0, len(ids))
	for _, id := range ids {
		results = append(results, s.Documents[id])
	}
	return results
}