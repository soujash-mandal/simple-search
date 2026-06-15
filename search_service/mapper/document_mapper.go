package mapper

import search_service_model "simple-search/search_service/model"
import search_engine_model "simple-search/search_engine/model"

func DocumentEngineToService(results []search_engine_model.Document) []search_service_model.Document {
	documents := make([]search_service_model.Document, len(results))
	for i, result := range results {
		documents[i] = search_service_model.Document{
			ID:      result.ID,
			Title:   result.Title,
			Content: result.Content,
		}
	}
	return documents
}

func DocumentServiceToEngine(results []search_service_model.Document) []search_engine_model.Document {
	documents := make([]search_engine_model.Document, len(results))
	for i, result := range results {
		documents[i] = search_engine_model.Document{
			ID:      result.ID,
			Title:   result.Title,
			Content: result.Content,
		}
	}
	return documents
}
