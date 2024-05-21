package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func handleEscapeSequence(runes []rune, i *int, length int) (rune, error) {
	if *i+1 < length {
		nextChar := runes[*i+1]
		if unicode.IsDigit(nextChar) || nextChar == '\\' {
			*i++
			return nextChar, nil
		}
		return 0, errors.New("invalid escape sequence")
	}
	return 0, errors.New("trailing backslash")
}

func handleRepeatedChar(result *strings.Builder, char rune, runes []rune, i *int, length int) error {
	if *i+1 < length && unicode.IsDigit(runes[*i+1]) {
		count, err := strconv.Atoi(string(runes[*i+1]))
		if err != nil {
			return err
		}
		result.WriteString(strings.Repeat(string(char), count-1))
		*i++
	}
	return nil
}

func UnpackString(input string) (string, error) {
	var result strings.Builder
	runes := []rune(input)
	length := len(runes)

	for i := 0; i < length; i++ {
		char := runes[i]

		if char == '\\' {
			escapedChar, err := handleEscapeSequence(runes, &i, length)
			if err != nil {
				return "", err
			}
			result.WriteRune(escapedChar)
			if err := handleRepeatedChar(&result, escapedChar, runes, &i, length); err != nil {
				return "", err
			}
		} else if unicode.IsDigit(char) {
			return "", errors.New("invalid string format")
		} else {
			result.WriteRune(char)
			if err := handleRepeatedChar(&result, char, runes, &i, length); err != nil {
				return "", err
			}
		}
	}

	return result.String(), nil
}

func main() {
	examples := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}
	for _, example := range examples {
		result, err := UnpackString(example)
		if err != nil {
			log.Printf("Error unpacking string %q: %v\n", example, err)
		} else {
			fmt.Printf("Unpacked string %q: %q\n", example, result)
		}
	}
}
