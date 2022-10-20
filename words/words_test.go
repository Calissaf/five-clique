package words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Words_FilterWords(t *testing.T) {
	type args struct {
		wordList []string
	}

	type testCase struct {
		args args
		want []string
	}

	testCases := map[string]testCase{
		"when given a list of words of variable length returns only words with 5 characters": {
			args: args{
				wordList: []string{
					"fooba",
					"barba",
					"bazba",
					"wowba",
					"yooba",
					"man",
					"manbara",
				},
			},
			want: []string{
				"fooba",
				"barba",
				"bazba",
				"wowba",
				"yooba",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := FilterWords(tc.args.wordList)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_Words_CheckUnique(t *testing.T) {
	type args struct {
		wordList []string
	}

	type testCase struct {
		args args
		want []string
	}

	testCases := map[string]testCase{
		"when given a list of words only returns words with distinct characters": {
			args: args{
				wordList: []string{
					"fooba",
					"barba",
					"those",
					"bazba",
					"brick",
					"jumpy",
				},
			},
			want: []string{
				"those",
				"brick",
				"jumpy",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := CheckUnique(tc.args.wordList)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_Words_checkCharInWord(t *testing.T) {
	type args struct {
		wordList []string
		newWord  string
	}

	type testCase struct {
		args args
		want bool
	}

	testCases := map[string]testCase{
		"returns true is any character matches between new word and word list": {
			args: args{
				wordList: []string{
					"those",
					"brick",
					"jumpy",
				},
				newWord: "chunk",
			},
			want: true,
		},
		"returns false if no character matches between new word and word list": {
			args: args{
				wordList: []string{
					"those",
					"brick",
					"jumpy",
				},
				newWord: "gland",
			},
			want: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := checkCharInWord(tc.args.wordList, tc.args.newWord)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_Words_WordList(t *testing.T) {
	type args struct {
		words map[string][]string
	}
	type testCase struct {
		args args
		want []string
	}

	testCases := map[string]testCase{
		"returns list of words with all distinct characters from word list": {
			args: args{
				words: map[string][]string{
					"abcde": {"abcde", "edcba"},
					"fghij": {"fghij", "jihgf"},
					// duplicate words
					"abdde": {"abdde", "eddbe"},
					"klghj": {"klghj"},
					//
					"klmno": {"klmmo", "onmlk"},
					"pqrst": {"pqrst", "tsrqp"},
				},
			},
			want: []string{
				"abcde",
				"pqrst",
				"klmno",
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := WordList(tc.args.words)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_Words_GenerateAnagrams(t *testing.T) {
	type args struct {
		wordList []string
	}

	type testCase struct {
		args args
		want map[string][]string
	}

	testCases := map[string]testCase{
		"maps anagrams of words": testCase{
			args: args{
				[]string{
					"abcde",
					"edcba",
					"cadbe",
					"badce",
					"lkjhg",
					"ghjkl",
				},
			},
			want: map[string][]string{
				"abcde": []string{
					"abcde",
					"edcba",
					"cadbe",
					"badce",
				},
				"ghjkl": []string{
					"lkjhg",
					"ghjkl",
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := GenerateAnagrams(tc.args.wordList)

			assert.Equal(t, tc.want, got)
		})
	}
}
