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

func Test_Words_FindCollisions(t *testing.T) {
	type args struct {
		anagrams map[string][]string
	}
	type testCase struct {
		args args
		want map[string][]string
	}

	testCases := map[string]testCase{
		"returns a collision map with anagram keys and values that don't collide with the anagram": {
			args: args{
				anagrams: map[string][]string{
					"abcde": {"abcde", "edcba"},
					"fghij": {"fghij", "jihgf"},
					"abdde": {"abdde", "eddbe"},
				},
			},
			want: map[string][]string{
				"abcde": {"fghij"},
				"fghij": {"abcde", "abdde"},
				"abdde": {"fghij"},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := FindCollisions(tc.args.anagrams)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_Words_checkCharInWord(t *testing.T) {
	type args struct {
		originalWord   string
		comparisonWord string
	}
	type testCase struct {
		args args
		want bool
	}

	testCases := map[string]testCase{
		"returns true is any character matches between comparison word and original word": {
			args: args{
				originalWord:   "abcde",
				comparisonWord: "hiaok",
			},
			want: true,
		},
		"returns false if no characters match between comparison word and original word": {
			args: args{
				originalWord:   "abcde",
				comparisonWord: "fghij",
			},
			want: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := checkCharInWord(tc.args.originalWord, tc.args.comparisonWord)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_Words_checkWordInAnagrams(t *testing.T) {
	type args struct {
		anagram        string
		comparisonWord string
	}
	type testCase struct {
		args args
		want bool
	}

	testCases := map[string]testCase{
		"returns a true if word is in list of anagrams": {
			args: args{
				anagram:        "fghij",
				comparisonWord: "jihgf",
			},
			want: true,
		},
		"returns a false if word is not in list of anagrams": {
			args: args{
				anagram:        "abcde",
				comparisonWord: "lmnop",
			},
			want: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := checkWordInAnagram(tc.args.anagram, tc.args.comparisonWord)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_Words_WordList(t *testing.T) {
	type args struct {
		collisonMap  map[string][]string
		fullWordList []string
	}
	type testCase struct {
		args args
		want []string
	}
	testCases := map[string]testCase{
		"returns list of character distinct words": {
			args: args{
				collisonMap: map[string][]string{
					"lmnop": {"fghij"},
					"fghij": {"abcde", "abdde"},
					"abdde": {"fghij"},
					"abcde": {"efghi", "jklmn", "opqrs", "tuvwx"},
				},
				fullWordList: []string{
					"lmnop",
					"fghij",
					"fghij",
					"opqrs",
					"efghi",
					"ihgef",
					"abdde",
					"srqpo",
					"abdde",
					"fghij",
					"efghi",
					"jklmn",
					"opqrs",
					"tuvwx",
					"abcde",
				},
			},
			want: []string{
				"abcde",
				"efghi",
				"jklmn",
				"opqrs",
				"tuvwx",
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := WordList(tc.args.collisonMap, tc.args.fullWordList)
			assert.Equal(t, tc.want, got)
		})
	}
}
