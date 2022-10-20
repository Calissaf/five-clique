package words

import (
	"sort"
	"strings"
	"unicode/utf8"
)

// filterWords will find words with character length equal to 5
func FilterWords(words []string) []string {
	var fiveLetters []string

	for _, word := range words {
		trimmed := clean([]byte(word))
		runes := utf8.RuneCountInString(trimmed)
		if runes == 5 {
			fiveLetters = append(fiveLetters, word)
		}
	}

	return fiveLetters
}

func clean(s []byte) string {
	j := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			s[j] = b
			j++
		}
	}
	return string(s[:j])
}

// checkUnique will find words without duplicate characters
func CheckUnique(words []string) []string {
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

func GenerateAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	for _, w := range words {
		s := strings.Split(w, "")
		sort.Strings(s)
		sorted := strings.Join(s, "")
		anagrams[sorted] = append(anagrams[sorted], w)
	}

	return anagrams
}

/*
map of collisions for every anagram
in WordList:
	disregard all collisions with length < maxWordList
	for each collision list > maxWordList compare with original words import

*/

func FindCollisions(anagrams map[string][]string) map[string][]string {
	collisions := make(map[string][]string)
	for anagramKey, _ := range anagrams {
		for collisionKey, _ := range anagrams {
			if !checkCharInWord(anagramKey, collisionKey) {
				collisions[anagramKey] = append(collisions[anagramKey], collisionKey)
			}

		}
	}
	return collisions
}

func checkCharInWord(originalWord string, comparsionWord string) bool {
	for _, originalChar := range originalWord {
		for _, compcomparsionChar := range comparsionWord {
			if compcomparsionChar == originalChar {
				return true
			}
		}
	}

	return false
}

// wordList will find a combination of words where every letter is unique across all the words
func WordList(collisonMap map[string][]string, fullWordList []string) []string {
	var distinctWords []string
	const maxWordsList = 5 // REMEBER TO CHANGE TEST AND VARIABLE TO 5
	//  key, value :=

	for key, values := range collisonMap {
		if len(values)+1 == maxWordsList {
			for _, word := range fullWordList {
				if len(distinctWords) == 0 && checkWordInAnagram(key, word) {
					distinctWords = append(distinctWords, word)
					break
				}
			}
			for _, value := range values {
				for _, word := range fullWordList {
					if len(distinctWords) > 0 && checkWordInAnagram(value, word) {
						distinctWords = append(distinctWords, word)
						break
					}
				}
			}
		}
	}

	return distinctWords
}

// checkWordInAnagramList will check if a word is in a list of anagrams
func checkWordInAnagram(anagram string, comparisonWord string) bool {
	var countMatches = 0
	for _, anagramCharacter := range anagram {
		for _, newWordCharacter := range comparisonWord {
			if newWordCharacter == anagramCharacter {
				countMatches++
				break
			}
		}
	}
	if countMatches == len(anagram) {
		return true
	} else {
		return false
	}
}
