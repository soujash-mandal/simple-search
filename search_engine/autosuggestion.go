package searchengine

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
