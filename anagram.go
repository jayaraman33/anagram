package anagram

import (
	"reflect"
	"strings"
	"unicode"
)

func countLetters(word string) map[rune]int {
	letters := make(map[rune]int)

	for _, letter := range strings.ToUpper(word) {
		if unicode.IsLetter(letter) {
			letters[letter]++
		}
	}
	return letters
}

func Detect(subject string, list []string) []string {
	var result []string

	letters := countLetters(subject)

	subject = strings.ToUpper(subject)

	for _, word := range list {
		target := countLetters(word)
		if subject != strings.ToUpper(word) && reflect.DeepEqual(letters, target) {
			result = append(result, word)
		}
	}

	return result
}
