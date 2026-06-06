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

func (s *SearchEngine) AddDocument(doc Document) {
	// Store the document
	s.Documents[doc.ID] = doc

	// Tokenize title and content
	titleTokens := tokenize(doc.Title)
	contentTokens := tokenize(doc.Content)

	// Combine all tokens
	tokens := append(titleTokens, contentTokens...)

	// Build inverted index
	for _, token := range tokens {
		// Create a set for the token if it doesn't exist
		if _, exists := s.Index[token]; !exists {
			s.Index[token] = make(map[int]struct{})
		}
		// Add document ID to the token's posting list
		s.Index[token][doc.ID] = struct{}{}
	}
}
