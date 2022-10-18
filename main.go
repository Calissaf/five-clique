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
	uniqueWords := checkUnique(filtered)
	uniqueWordList := wordList(uniqueWords)

	fmt.Println(len(uniqueWordList))

	for _, word := range uniqueWordList {
		println(word)
	}
}

// filterWords will find words with character length equal to 5
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

// loadWords will read file of words into string array
func loadWords(files []string) string {
	var words []string
	for _, f := range files {
		file, err := os.ReadFile(f)
		words = append(words, string(file))

		if err != nil {
			log.Fatal(fmt.Sprintf("error reading file: %v", err))
		}
	}

	return strings.Join(words, "")
}

// wordList will find a combination of words where every letter is unique across all the words
func wordList(words []string) []string {
	var uniqueWordList []string
	var maxListSize = 5
	var count = 0

	for len(uniqueWordList) < maxListSize {
		uniqueWordList = append(uniqueWordList, words[count])
		for _, word := range words {
			for i, char := range word {
				uniqueChar := checkCharInWord(uniqueWordList, char)
				if uniqueChar == true {
					break
				}
				if i == len(word)-1 {
					uniqueWordList = append(uniqueWordList, word)
				}
			}
		}
		count++
		uniqueWordList = nil
		if count == len(words) {
			break
		}
	}
	return uniqueWordList
}

// checkCharInWord will check if a character is in a word list
func checkCharInWord(words []string, character rune) bool {
	for _, word := range words {
		for _, char := range word {
			if character == char {
				return true
			}
		}
	}
	return false
}
