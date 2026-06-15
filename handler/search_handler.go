package handler

import (
	"encoding/json"
	"net/http"
	searchservice "simple-search/search_service"
	searchservicemodel "simple-search/search_service/model"
)

type SearchHandler struct {
	searchService *searchservice.SearchService
}

func NewSearchHandler(
	searchService *searchservice.SearchService,
) *SearchHandler {
	return &SearchHandler{
		searchService: searchService,
	}
}



func (h *SearchHandler) Document(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var doc searchservicemodel.CreateDocumentRequest
		if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.searchService.AddDocument(doc)
		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			h.searchService.GetAllDocuments(),
		)
	default:
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results := h.searchService.Search(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}



func (h *SearchHandler) Autocomplete(w http.ResponseWriter, r *http.Request) {
	prefix := r.URL.Query().Get("q")
	results := h.searchService.AutoComplete(prefix)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (h *SearchHandler) Suggest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results := h.searchService.Suggest(query)
	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	json.NewEncoder(w).Encode(
		results,
	)
}

func (h *SearchHandler) PhraseSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	results := h.searchService.PhraseSearch(query)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}