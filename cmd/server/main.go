package main

import (
	"log"
	"net/http"
	"simple-search/container"
	"simple-search/handler"
	"simple-search/router"
)

func main() {

	c, err := container.Build()
	if err != nil {
		log.Fatal(err)
	}

	searchHandler := handler.NewSearchHandler(
		c.SearchService,
	)
	analyticsHandler:=handler.NewAnalyticsHandler(
		c.AnalyticsService,
	)

	router.RegisterRoutes(searchHandler,analyticsHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
