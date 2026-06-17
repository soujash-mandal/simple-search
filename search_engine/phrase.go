package searchengine

import internal_model "simple-search/internal/model"


func (s *SearchEngine) PhraseSearch(
	phrase string,
) []internal_model.Document {

	tokens := tokenize(phrase)

	if len(tokens) == 0 {
		return nil
	}

	firstPosting, exists :=
		s.PositionalIndex[tokens[0]]

	if !exists {
		return nil
	}

	var results []internal_model.Document

	for docID, firstPositions := range firstPosting {

		found := false

		for _, startPos := range firstPositions {

			match := true

			for i := 1; i < len(tokens); i++ {

				posting, exists :=
					s.PositionalIndex[tokens[i]]

				if !exists {
					match = false
					break
				}

				positions :=
					posting[docID]

				expected :=
					startPos + i

				ok := false

				for _, pos := range positions {
					if pos == expected {
						ok = true
						break
					}
				}

				if !ok {
					match = false
					break
				}
			}

			if match {
				found = true
				break
			}
		}

		if found {
			results = append(
				results,
				s.Documents[docID],
			)
		}
	}

	return results
}