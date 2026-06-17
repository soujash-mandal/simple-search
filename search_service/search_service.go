package searchservice

import (
	analyticsservice "simple-search/analytics_service"
	internal_model "simple-search/internal/model"
	searchengine "simple-search/search_engine"
	"simple-search/search_service/model"
	"simple-search/search_service/repository"
)

type SearchService struct {
	repository        repository.SearchRepository
	engine            searchengine.SearchEngine
	analytics_service analyticsservice.AnalyticsService
}

func NewSearchService(
	repository repository.SearchRepository,
	engine searchengine.SearchEngine,
	analytics_service analyticsservice.AnalyticsService,
) *SearchService {
	return &SearchService{
		repository:        repository,
		engine:            engine,
		analytics_service: analytics_service,
	}
}

func (s *SearchService) AddDocument(doc model.CreateDocumentRequest) {
	new_doc, _ := s.repository.Save(model.CreateDocumentRequest{
		Title:   doc.Title,
		Content: doc.Content,
	})
	s.engine.AddDocument(new_doc)
}

func (s *SearchService) LoadDocuments() {
	docs, _ := s.repository.GetAll()
	for _, doc := range docs {
		s.engine.AddDocument(doc)
	}

}

func (s *SearchService) GetAllDocuments() []internal_model.Document {
	results, _ := s.repository.GetAll()
	return results
}

func (s *SearchService) Search(query string) []internal_model.Document {
	results := s.engine.Search(query)
	s.analytics_service.RecordQuery(query)
	return results
}

func (s *SearchService) AutoComplete(prefix string) []string {
	return s.engine.AutoComplete(prefix)
}

func (s *SearchService) Suggest(query string) []string {
	return s.engine.Suggest(query, 2)
}

func (s *SearchService) PhraseSearch(
	phrase string,
) []internal_model.Document {
	return s.engine.PhraseSearch(phrase)
}
