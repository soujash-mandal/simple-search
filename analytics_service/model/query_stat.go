package model

type QueryStat struct {
	Query string `json:"query"`
	Count int    `json:"count"`
}