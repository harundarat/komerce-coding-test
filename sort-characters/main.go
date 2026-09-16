package main

import (
	"fmt"
	"strings"
)

func processCharacters(input string) (vowels, consonants string) {
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))

	var vowelOrder, consonantOrder []rune
	characterCounts := make(map[rune]int)

	isVowel := func(r rune) bool {
		return strings.ContainsRune("aeiou", r)
	}

	for _, character := range input {
		if characterCounts[character] == 0 {
			if isVowel(character) {
				vowelOrder = append(vowelOrder, character)
			} else {
				consonantOrder = append(consonantOrder, character)
			}
		}

		characterCounts[character]++
	}

	var vowelBuilder, consonantBuilder strings.Builder

	for _, v := range vowelOrder {
		vowelBuilder.WriteString(strings.Repeat(string(v), characterCounts[v]))
	}

	for _, c := range consonantOrder {
		consonantBuilder.WriteString(strings.Repeat(string(c), characterCounts[c]))
	}

	return vowelBuilder.String(), consonantBuilder.String()
}

func main() {
	testCases := []string{
		"Sample Case",
		"Next Case",
		"Makanan Enak",
	}

	for _, tc := range testCases {
		fmt.Printf("Input: %s \n", tc)
		vowels, consonants := processCharacters(tc)
		fmt.Printf("Vowel Characters: %s \n", vowels)
		fmt.Printf("Consonants Characters: %s \n\n", consonants)
	}
}
