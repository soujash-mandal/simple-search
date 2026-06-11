package main

import (
	"math"
	"sort"
)

const (
	k1 = 1.2
	b  = 0.75
)

type SearchEngine struct {
	Documents       map[int]Document
	Index           map[string]map[int]int
	PositionalIndex map[string]map[int][]int
	DocLengths      map[int]int
	Trie            *Trie
}

func NewSearchEngine() *SearchEngine {
	return &SearchEngine{
		Documents:       make(map[int]Document),
		Index:           make(map[string]map[int]int),
		PositionalIndex: make(map[string]map[int][]int),
		DocLengths:      make(map[int]int),
		Trie:            NewTrie(),
	}
}

func (s *SearchEngine) AddDocument(doc Document) {
	s.Documents[doc.ID] = doc
	titleTokens := tokenize(doc.Title)
	contentTokens := tokenize(doc.Content)
	tokens := append(titleTokens, contentTokens...)
	s.DocLengths[doc.ID] = len(tokens)
	for pos, token := range tokens {
		if _, exists := s.Index[token]; !exists {
			s.Index[token] = make(map[int]int)
		}
		s.Index[token][doc.ID]++
		if _, exists := s.PositionalIndex[token]; !exists {
			s.PositionalIndex[token] = make(map[int][]int)
		}
		s.PositionalIndex[token][doc.ID] = append(s.PositionalIndex[token][doc.ID], pos)
		s.Trie.Insert(token)
	}
}

func (s *SearchEngine) GetAllDocuments() []Document {
	docs := make([]Document, 0, len(s.Documents))
	for _, doc := range s.Documents {
		docs = append(docs, doc)
	}
	return docs
}

func (s *SearchEngine) avgDocLength() float64 {
	if len(s.DocLengths) == 0 {
		return 0
	}
	total := 0
	for _, length := range s.DocLengths {
		total += length
	}
	return float64(total) / float64(len(s.DocLengths))
}

func (s *SearchEngine) idf(term string) float64 {
	df := len(s.Index[term])
	if df == 0 {
		return 0
	}
	N := len(s.Documents)
	return math.Log(1 + (float64(N)-float64(df)+0.5)/(float64(df)+0.5))
}

func (s *SearchEngine) bm25(tf int, docLength int, avgDocLength float64, idf float64) float64 {
	numerator := float64(tf) * (k1 + 1)
	denominator := float64(tf) +
		k1*(1-b+b*(float64(docLength)/avgDocLength))
	return idf * (numerator / denominator)
}

func (s *SearchEngine) Search(query string) []Document {
	tokens := tokenize(query)
	if len(tokens) == 0 {
		return nil
	}
	avgDocLength := s.avgDocLength()
	scores := make(map[int]float64)
	for _, token := range tokens {
		postings, exists := s.Index[token]
		if !exists {
			continue
		}
		idf := s.idf(token)
		for docID, tf := range postings {
			docLength := s.DocLengths[docID]
			scores[docID] += s.bm25(
				tf,
				docLength,
				avgDocLength,
				idf,
			)
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

func (s *SearchEngine) AutoComplete(prefix string) []string {
	return s.Trie.AutoComplete(prefix)
}

func (s *SearchEngine) Suggest(query string, maxDistance int) []string {
	words := s.Trie.Words()
	var suggestions []string
	for _, word := range words {
		distance := Levenshtein(
			query,
			word,
		)
		if distance <= maxDistance {
			suggestions = append(
				suggestions,
				word,
			)
		}
	}
	return suggestions
}
