// Package textprocessor
// Defines a text processor
package textprocessor

import (
	"slices"
	"strings"
)

// Processor defines the functionality for processing text
type Processor interface {
	// ReverseWords reverses the order of words in a given string.
	// Words are separated by spaces. Punctuation is part of the word.
	// Example: "hello, world!" becomes "world! hello,"
	ReverseWords(input string) string

	// CountVowels returns the number of vowels (a,e,i,o,u case-insentive)
	CountVowels(input string) int
}

type MyProcessor struct{}

func (m MyProcessor) ReverseWords(input string) string {
	empty := " "
	s := strings.Split(input, empty)
	l := len(s)
	var sb strings.Builder
	sb.Grow(len(input))
	for i := range s {
		r := l - i - 1
		v := s[r]
		if v != "" {
			sb.WriteString(v)
			if i < (l - 1) {
				sb.WriteString(empty)
			}
		}
	}
	return sb.String()
}

func (m MyProcessor) CountVowels(input string) int {
	t := strings.ToLower(input)
	count := 0
	vowels := []rune("aeiou")
	for _, v := range t {
		if slices.Contains(vowels, v) {
			count++
		}
	}
	return count
}
