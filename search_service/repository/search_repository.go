package repository

import (
	internal_model "simple-search/internal/model"
	model "simple-search/search_service/model"
)

type SearchRepository interface {
	Save(doc model.CreateDocumentRequest) (internal_model.Document, error)
	GetAll() ([]internal_model.Document, error)
}
