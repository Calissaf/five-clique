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

	//  var test = []string{"brick", "brick", "glent", "jumpy", "vozhd", "waqfs"}
	// var test = []string{"brick", "jumpy", "brick", "jumpy", "snail", "trunk", "glent"}

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
	var finalUniqueWordList []string
	var maxListSize = 2
	var count = 0
	var completeList = false

	for completeList == false {
		finalUniqueWordList = append(finalUniqueWordList, words[count])
		for _, word := range words {
			charInWordList := checkCharInWord(finalUniqueWordList, word)
			if charInWordList == false {
				finalUniqueWordList = append(finalUniqueWordList, word)
			}
			if len(finalUniqueWordList) == maxListSize {
				break
			}
		}
		if len(finalUniqueWordList) < maxListSize {
			count++
			finalUniqueWordList = nil
		} else {
			completeList = true
		}
		if count == len(words) {
			break
		}
	}
	return finalUniqueWordList
}

// checkCharInWord will check if a character is in a word list
func checkCharInWord(wordList []string, newWord string) bool {
	for _, word := range wordList {
		for _, wordListCharacter := range word {
			for _, newWordCharacter := range newWord {
				if newWordCharacter == wordListCharacter {
					return true
				}
			}
		}
	}
	return false
}
