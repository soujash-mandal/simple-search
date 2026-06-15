package searchservice

import (
	analyticsservice "simple-search/analytics_service"
	searchengine "simple-search/search_engine"
	search_engine_model "simple-search/search_engine/model"
	"simple-search/search_service/mapper"
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
	s.engine.AddDocument(search_engine_model.Document{
		ID:      new_doc.ID,
		Title:   doc.Title,
		Content: doc.Content,
	})
}

func (s *SearchService) LoadDocuments() {
	repo_docs, _ := s.repository.GetAll()
	engine_docs := mapper.DocumentServiceToEngine(repo_docs)
	for _, doc := range engine_docs {
		s.engine.AddDocument(doc)
	}

}

func (s *SearchService) GetAllDocuments() []model.Document {
	results, _ := s.repository.GetAll()
	return results
}

func (s *SearchService) Search(query string) []model.Document {
	results := s.engine.Search(query)
	s.analytics_service.RecordQuery(query)
	return mapper.DocumentEngineToService(results)
}

func (s *SearchService) AutoComplete(prefix string) []string {
	return s.engine.AutoComplete(prefix)
}

func (s *SearchService) Suggest(query string) []string {
	return s.engine.Suggest(query, 2)
}

func (s *SearchService) PhraseSearch(
	phrase string,
) []model.Document {
	engine_docs := s.engine.PhraseSearch(phrase)
	return mapper.DocumentEngineToService(engine_docs)
}
