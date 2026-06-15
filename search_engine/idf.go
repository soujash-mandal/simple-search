package searchengine

import "math"

func (s *SearchEngine) idf(term string) float64 {
	df := len(s.Index[term])
	if df == 0 {
		return 0
	}
	N := len(s.Documents)
	return math.Log(1 + (float64(N)-float64(df)+0.5)/(float64(df)+0.5))
}