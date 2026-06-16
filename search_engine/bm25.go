package searchengine

const (
	k1 = 1.2
	b  = 0.75
)

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

func (s *SearchEngine) bm25(tf float64, docLength int, avgDocLength float64, idf float64) float64 {
	numerator := float64(tf) * (k1 + 1)
	denominator := float64(tf) +
		k1*(1-b+b*(float64(docLength)/avgDocLength))
	return idf * (numerator / denominator)
}