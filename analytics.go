package main

import (
	"sort"
)

type QueryStat struct {
	Query string `json:"query"`
	Count int    `json:"count"`
}

func (s *SearchEngine) RecordQuery(query string) {
	if query == "" {
		return
	}
	s.QueryCounts[query]++
}

func (s *SearchEngine) TopQueries(limit int) []QueryStat {
	stats := make([]QueryStat, 0, len(s.QueryCounts))
	for query, count := range s.QueryCounts {
		stats = append(stats, QueryStat{
			Query: query,
			Count: count,
		})
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[j].Count
	})
	if limit > 0 && len(stats) > limit {
		stats = stats[:limit]
	}
	return stats
}
