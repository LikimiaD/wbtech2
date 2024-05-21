package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "апельсин", "спаниель", "пенсил", "cat", "tac", "act"}
	expected := map[string][]string{
		"пятак":    {"пятак", "пятка", "тяпка"},
		"листок":   {"листок", "слиток", "столик"},
		"апельсин": {"апельсин", "спаниель"},
		"cat":      {"act", "cat", "tac"},
	}

	result := findAnagrams(words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
