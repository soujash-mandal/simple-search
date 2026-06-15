package container

import (
	"log"
	"os"
	analyticsservice "simple-search/analytics_service"
	"simple-search/database"
	searchengine "simple-search/search_engine"
	searchservice "simple-search/search_service"
	searchadapters "simple-search/search_service/adapters"

	"github.com/joho/godotenv"
)

type Container struct {
	SearchService    *searchservice.SearchService
	AnalyticsService *analyticsservice.AnalyticsService
}

func Build() (*Container, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	log.Println("Loaded .env file")

	err = database.ConnectMongo(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	_search_document_collection := database.DocumentCollection()
	_search_repository := searchadapters.NewMongoSearchRepository(
		_search_document_collection,
	)
	_search_engine := *searchengine.NewSearchEngine()
	analytics_service_di := analyticsservice.NewAnalyticsService()

	search_service_di := searchservice.NewSearchService(
		_search_repository,
		_search_engine,
		*analytics_service_di,
	)

	search_service_di.LoadDocuments()
	log.Println("Loaded Documents for search")

	return &Container{
		SearchService: search_service_di,
		AnalyticsService: analytics_service_di,
	}, nil
}
