package model


type CreateDocumentRequest struct {
	Title   string
	Content string
}

type Document struct {
	ID      string
	Title   string
	Content string
}
