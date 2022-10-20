package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/calissaf/five_clique/words"
)

func main() {
	files := []string{
		"wordle-allowed-guesses.txt",
		"wordle-answers-alphabetical.txt",
		"words_alpha.txt",
	}

	allWords := loadWords(files)
	newLineWords := strings.Split(allWords, "\n")
	filtered := words.FilterWords(newLineWords)
	uniqueWords := words.CheckUnique(filtered)
	reverse(uniqueWords)
	anagrams := words.GenerateAnagrams(uniqueWords)
	fmt.Printf("words: %d | anagrams %d\n", len(uniqueWords), len(anagrams))
	fmt.Printf("%d words saved\n", len(uniqueWords)-len(anagrams))

	//  var test = []string{"brick", "brick", "glent", "jumpy", "vozhd", "waqfs"}
	// var test = []string{"brick", "jumpy", "brick", "jumpy", "snail", "trunk", "glent"}

	uniqueWordList := words.WordList(anagrams)

	fmt.Println(len(uniqueWordList))

	for _, word := range uniqueWordList {
		println(word)
	}
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

// loadWords will read file of words into string array
func loadWords(files []string) string {
	var w []string
	for _, f := range files {
		file, err := os.ReadFile(f)
		w = append(w, string(file))

		if err != nil {
			log.Fatal(fmt.Sprintf("error reading file: %v", err))
		}
	}

	return strings.Join(w, "")
}
