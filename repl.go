package main

import (
	"strings"
)

func cleanInput(text string) []string {
	word := ""
	words := []string{}
	for _, char := range text {
		if char == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += strings.ToLower(string(char))
		}
	}
	return words
}