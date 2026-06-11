package main

import (
	"encoding/json"
	"os"
)

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