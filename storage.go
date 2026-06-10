package main

import (
	"encoding/json"
	"os"
)

const documentsFile = "data/documents.json"

func SaveDocuments(docs []Document) error {
	file, err := os.Create(documentsFile)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(docs)
}

func LoadDocuments() ([]Document, error) {
	file, err := os.Open(documentsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var docs []Document
	err = json.NewDecoder(file).Decode(&docs)
	return docs, err
}