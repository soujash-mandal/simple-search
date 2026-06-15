package router

import (
	"net/http"
	"simple-search/handler"
)

func RegisterRoutes(
	searchHandler *handler.SearchHandler,
	analyticsHandler *handler.AnalyticsHandler,
) {

	// Documents
	http.HandleFunc(
		"/documents",
		searchHandler.Document,
	)

	// Search
	http.HandleFunc(
		"/search",
		searchHandler.Search,
	)

	// Phrase Search
	http.HandleFunc(
		"/phrase-search",
		searchHandler.PhraseSearch,
	)

	// Autocomplete
	http.HandleFunc(
		"/autocomplete",
		searchHandler.Autocomplete,
	)

	// Typo Suggestions
	http.HandleFunc(
		"/suggest",
		searchHandler.Suggest,
	)

	// Analytics
	http.HandleFunc(
		"/top-queries",
		analyticsHandler.TopQueries,
	)
}