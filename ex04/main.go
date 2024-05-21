package main

import (
	"fmt"
	"sort"
	"strings"
)

func normalize(word string) string {
	word = strings.ToLower(word)
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	normalizedWords := make(map[string]string)

	for _, word := range words {
		loweredWord := strings.ToLower(word)
		normalized := normalize(loweredWord)

		if original, exists := normalizedWords[normalized]; exists {
			anagrams[original] = append(anagrams[original], loweredWord)
		} else {
			normalizedWords[normalized] = loweredWord
			anagrams[loweredWord] = []string{loweredWord}
		}
	}

	for key, group := range anagrams {
		if len(group) <= 1 {
			delete(anagrams, key)
		} else {
			sort.Strings(group)
		}
	}

	return anagrams
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "апельсин", "спаниель", "пенсил", "cat", "tac", "act"}
	anagrams := findAnagrams(words)
	for key, group := range anagrams {
		fmt.Printf("Key: %s, Anagrams: %v\n", key, group)
	}
}
