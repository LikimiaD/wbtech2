package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var (
	after   = flag.Int("A", 0, "указание сколько строк выводить после совпадения")
	before  = flag.Int("B", 0, "указание сколько строк выводить до совпадения")
	context = flag.Int("C", 0, "указание сколько строк выводить до и после совпадения")
	count   = flag.Bool("c", false, "вывести количество строк совпадения")
	ignore  = flag.Bool("i", false, "игнорировать регистр")
	invert  = flag.Bool("v", false, "вместо совпадения исключать")
	fixed   = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum = flag.Bool("n", false, "напечатать номер строки")
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("usage: grep [options] pattern inputfile")
		return
	}

	pattern := flag.Arg(0)
	inputFile := flag.Arg(1)

	if *context > 0 {
		*after = *context
		*before = *context
	}

	lines, err := readLines(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	matches := grep(lines, pattern)
	if *count {
		fmt.Println(len(matches))
	} else {
		for _, match := range matches {
			fmt.Println(match)
		}
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func compilePattern(pattern string) (*regexp.Regexp, error) {
	if *fixed {
		pattern = regexp.QuoteMeta(pattern)
	} else if *ignore {
		pattern = "(?i)" + pattern
	}
	return regexp.Compile(pattern)
}

func matchesPattern(line string, re *regexp.Regexp) bool {
	return re.MatchString(line)
}

func getContext(lines []string, start, end int, lineNum bool, added map[int]bool) []string {
	var result []string
	for i := start; i < end; i++ {
		if !added[i] {
			if lineNum {
				result = append(result, fmt.Sprintf("%d:%s", i+1, lines[i]))
			} else {
				result = append(result, lines[i])
			}
			added[i] = true
		}
	}
	return result
}

func grep(lines []string, pattern string) []string {
	var results []string
	re, err := compilePattern(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil
	}

	matchCount := 0
	added := make(map[int]bool)
	for i := 0; i < len(lines); i++ {
		matched := matchesPattern(lines[i], re)
		if (matched && !*invert) || (!matched && *invert) {
			matchCount++
			if !*count {
				start := max(0, i-*before)
				end := min(len(lines), i+*after+1)
				context := getContext(lines, start, end, *lineNum, added)
				results = append(results, context...)
			}
		}
	}
	if *count {
		results = append(results, fmt.Sprintf("%d", matchCount))
	}

	return results
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
