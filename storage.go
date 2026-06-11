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

const analyticsFile = "data/query_analytics.json"

func SaveAnalytics(data map[string]int) error {
	file, err := os.Create(analyticsFile)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func LoadAnalytics() (map[string]int, error) {
	file, err := os.Open(analyticsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var data map[string]int
	err = json.NewDecoder(file).Decode(&data)
	return data, err
}