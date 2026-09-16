package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// processCharacters separates vowels and consonants while preserving their original order.
func processCharacters(input string) (vowels, consonants string) {
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))

	var vowelOrder, consonantOrder []rune
	characterCounts := make(map[rune]int)

	isVowel := func(r rune) bool {
		return strings.ContainsRune("aeiou", r)
	}

	for _, character := range input {
		// Record unique characters to preserve appearance order
		if characterCounts[character] == 0 {
			if isVowel(character) {
				vowelOrder = append(vowelOrder, character)
			} else {
				consonantOrder = append(consonantOrder, character)
			}
		}

		characterCounts[character]++
	}

	// Using String Builder for efficient string concatenation
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input one line of word (S): ")
	inputStr, _ := reader.ReadString('\n')

	inputStr = strings.TrimSpace(inputStr)

	vowels, consonants := processCharacters(inputStr)

	fmt.Printf("Vowel Characters: %s\n", vowels)
	fmt.Printf("Consonant Characters: %s\n", consonants)
}
