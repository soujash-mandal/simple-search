package handler

import (
	"encoding/json"
	"net/http"
	analyticsservice "simple-search/analytics_service"
)

type AnalyticsHandler struct {
	analyticsService *analyticsservice.AnalyticsService
}

func NewAnalyticsHandler(
	analyticsService *analyticsservice.AnalyticsService,
) *AnalyticsHandler {
	return &AnalyticsHandler{
		analyticsService: analyticsService,
	}
}

func (a *AnalyticsHandler) TopQueries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a.analyticsService.TopQueries(10))
}