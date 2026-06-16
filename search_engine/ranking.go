package searchengine

const (
	TitleBoost   = 3.0
	ContentBoost = 1.0
)

func (s *SearchEngine) scoreToken(
	token string,
	docID string,
	avgDocLength float64,
) float64 {

	postings, exists := s.Index[token]
	if !exists {
		return 0
	}

	tf, exists := postings[docID]
	if !exists {
		return 0
	}

	titleTf := 0

	if titlePostings, exists := s.TitleIndex[token]; exists {
		titleTf = titlePostings[docID]
	}

	boostedTf :=
		float64(tf)*ContentBoost +
			float64(titleTf)*TitleBoost

	return s.bm25(
		boostedTf,
		s.DocLengths[docID],
		avgDocLength,
		s.idf(token),
	)
}