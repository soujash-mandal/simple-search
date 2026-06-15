package repository

import model "simple-search/search_service/model"


type SearchRepository interface {
	Save(doc model.CreateDocumentRequest) (model.Document, error)
	GetAll() ([]model.Document, error)
}