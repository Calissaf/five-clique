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

/*
abcde
cabde
badce


"abcde": []string{
	"abcde",
	"cabde",
	"badce",
}

collisions = abcde ** bdefg
abcde {


}

map[string][]string{
	"abcde": []string{
		...
		"zplio",
		...
	}
	"zplio": []string{
		"ghest",
	}
	"ghest": []string{
		"yuqwz",
		"yuqwm",
	}
}

map of collisions for every anagram
in WordList:
	disregard all collisions with length < maxWordList
	for each collision list > maxWordList compare with original words import

*/

// wordList will find a combination of words where every letter is unique across all the words
func WordList(wordMap map[string][]string) []string {
	var words []string
	var anagrams []string
	const maxWordsList = 5
	//  key, value :=
	for key, _ := range wordMap {
		anagrams = append(anagrams, key)
	}

	sort.Strings(anagrams)

	for i, j := 0, len(anagrams)-1; i < len(anagrams)/2; i, j = i+1, j-1 {
		beginning := anagrams[i]
		end := anagrams[j]

		if !checkCharInWord(words, beginning) && len(words) < maxWordsList {
			words = append(words, beginning)
		}
		if !checkCharInWord(words, end) && len(words) < maxWordsList {
			words = append(words, end)
		}
		if len(words) == maxWordsList {
			break
		}

	}

	return words
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
