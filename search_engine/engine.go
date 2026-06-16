package searchengine

import (
	"simple-search/search_engine/model"
	"sort"
)

type SearchEngine struct {
	Documents       map[string]model.Document
	Index           map[string]map[string]int
	TitleIndex      map[string]map[string]int
	PositionalIndex map[string]map[string][]int
	DocLengths      map[string]int
	Trie            *Trie
}

func NewSearchEngine() *SearchEngine {
	return &SearchEngine{
		Documents:       make(map[string]model.Document),
		Index:           make(map[string]map[string]int),
		TitleIndex:      make(map[string]map[string]int),
		PositionalIndex: make(map[string]map[string][]int),
		DocLengths:      make(map[string]int),
		Trie:            NewTrie(),
	}
}

func (s *SearchEngine) AddDocument(doc model.Document) {
	s.Documents[doc.ID] = doc
	titleTokens := tokenize(doc.Title)
	contentTokens := tokenize(doc.Content)
	tokens := append(titleTokens, contentTokens...)
	s.DocLengths[doc.ID] = len(tokens)

	for _, token := range titleTokens {
		if _, exists := s.TitleIndex[token]; !exists {
			s.TitleIndex[token] = make(map[string]int)
		}
		s.TitleIndex[token][doc.ID]++
	}

	for pos, token := range tokens {
		if _, exists := s.Index[token]; !exists {
			s.Index[token] = make(map[string]int)
		}
		s.Index[token][doc.ID]++
		if _, exists := s.PositionalIndex[token]; !exists {
			s.PositionalIndex[token] = make(map[string][]int)
		}
		s.PositionalIndex[token][doc.ID] = append(s.PositionalIndex[token][doc.ID], pos)
		s.Trie.Insert(token)
	}
}

func (s *SearchEngine) GetAllDocuments() []model.Document {
	docs := make([]model.Document, 0, len(s.Documents))
	for _, doc := range s.Documents {
		docs = append(docs, doc)
	}
	return docs
}

func (s *SearchEngine) Search(query string) []model.Document {
	tokens := tokenize(query)
	if len(tokens) == 0 {
		return nil
	}
	avgDocLength := s.avgDocLength()
	scores := make(map[string]float64)
	for _, token := range tokens {
		postings, exists := s.Index[token]
		if !exists {
			continue
		}
		for docID := range postings {
			scores[docID] += s.scoreToken(
				token,
				docID,
				avgDocLength,
			)
		}
	}
	if len(scores) == 0 {
		return nil
	}
	ids := make([]string, 0, len(scores))
	for id := range scores {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool {
		return scores[ids[i]] > scores[ids[j]]
	})
	results := make([]model.Document, 0, len(ids))
	for _, id := range ids {
		results = append(results, s.Documents[id])
	}
	return results
}
