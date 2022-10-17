package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	files := []string{
		"wordle-allowed-guesses.txt",
		"wordle-answers-alphabetical.txt",
		"words_alpha.txt",
	}

	allWords := loadWords(files)
	words := strings.Split(allWords, "\n")
	filtered := filterWords(words)
	unique := checkUnique(filtered)

	for _, word := range unique {

		fmt.Println(word)
	}
}

func filterWords(words []string) []string {
	var fiveLetters []string

	for _, word := range words {
		runes := utf8.RuneCountInString(word)
		if runes == 6 {
			fiveLetters = append(fiveLetters, word)
		}
	}

	return fiveLetters
}

// checkUnique will find words without duplicate characters
func checkUnique(words []string) []string {
	var uniqueWord []string
	for _, word := range words {
		for i, char := range word {
			if strings.Count(word, string(char)) > 1 {
				break
			}
			if i == len(word)-1 {
				uniqueWord = append(uniqueWord, word)
			}
		}
	}
	return uniqueWord
}

func loadWords(files []string) string {
	var words []string
	for _, f := range files {
		file, err := os.ReadFile(f)
		words = append(words, string(file))

		if err != nil {
			log.Fatal(fmt.Sprintf("errors reading file: %v", err))
		}
	}

	return strings.Join(words, "")
}
