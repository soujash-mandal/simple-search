package analyticsservice

import (
	"simple-search/analytics_service/model"
	"sort"
)



type AnalyticsService struct {
	QueryCounts map[string]int
}

func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{
		QueryCounts: make(map[string]int),
	}
}

func (s *AnalyticsService) RecordQuery(query string) {
	if query == "" {
		return
	}
	s.QueryCounts[query]++
}

func (s *AnalyticsService) TopQueries(limit int) []model.QueryStat {
	stats := make([]model.QueryStat, 0, len(s.QueryCounts))
	for query, count := range s.QueryCounts {
		stats = append(stats, model.QueryStat{
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
