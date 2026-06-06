package main

import "strings"

func tokenize(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
