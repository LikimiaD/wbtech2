package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	column       = flag.Int("k", 0, "указание колонки для сортировки")
	numeric      = flag.Bool("n", false, "сортировать по числовому значению")
	reverse      = flag.Bool("r", false, "сортировать в обратном порядке")
	unique       = flag.Bool("u", false, "не выводить повторяющиеся строки")
	monthSort    = flag.Bool("M", false, "сортировать по названию месяца")
	ignoreSpaces = flag.Bool("b", false, "игнорировать хвостовые пробелы")
	checkSorted  = flag.Bool("c", false, "проверять отсортированы ли данные")
	humanSort    = flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("usage: sort [options] inputfile outputfile")
		return
	}

	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	lines, err := readLines(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	if *ignoreSpaces {
		for i, line := range lines {
			lines[i] = strings.TrimRight(line, " \t")
		}
	}

	if *unique {
		lines = uniqueLines(lines)
	}

	if *checkSorted {
		if isSorted(lines, *numeric, *reverse, *column, *monthSort, *humanSort) {
			fmt.Println("The file is sorted.")
		} else {
			fmt.Println("The file is not sorted.")
		}
		return
	}

	sortLines(lines, *numeric, *reverse, *column, *monthSort, *humanSort)

	err = writeLines(lines, outputFile)
	if err != nil {
		fmt.Println("Error writing output file:", err)
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

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

func uniqueLines(lines []string) []string {
	uniqueMap := make(map[string]struct{})
	var uniqueLines []string
	for _, line := range lines {
		if _, found := uniqueMap[line]; !found {
			uniqueMap[line] = struct{}{}
			uniqueLines = append(uniqueLines, line)
		}
	}
	return uniqueLines
}

func isSorted(lines []string, numeric, reverse bool, column int, monthSort, humanSort bool) bool {
	sortedLines := make([]string, len(lines))
	copy(sortedLines, lines)
	sortLines(sortedLines, numeric, reverse, column, monthSort, humanSort)
	for i, line := range lines {
		if line != sortedLines[i] {
			return false
		}
	}
	return true
}

func sortLines(lines []string, numeric, reverse bool, column int, monthSort, humanSort bool) {
	sort.SliceStable(lines, func(i, j int) bool {
		a, b := lines[i], lines[j]

		a = getColumnValue(a, column)
		b = getColumnValue(b, column)

		if numeric && compareNumeric(&a, &b, reverse) {
			return a < b
		}

		if monthSort && compareMonths(&a, &b, reverse) {
			return a < b
		}

		if humanSort && compareHumanNumbers(&a, &b, reverse) {
			return a < b
		}

		if reverse {
			return a > b
		}
		return a < b
	})
}

func getColumnValue(line string, column int) string {
	if column > 0 {
		words := strings.Fields(line)
		if column-1 < len(words) {
			return words[column-1]
		}
	}
	return line
}

func compareNumeric(a, b *string, reverse bool) bool {
	numA, errA := strconv.ParseFloat(*a, 64)
	numB, errB := strconv.ParseFloat(*b, 64)
	if errA == nil && errB == nil {
		if reverse {
			*a, *b = strconv.FormatFloat(numA, 'f', -1, 64), strconv.FormatFloat(numB, 'f', -1, 64)
			return numA > numB
		}
		*a, *b = strconv.FormatFloat(numA, 'f', -1, 64), strconv.FormatFloat(numB, 'f', -1, 64)
		return numA < numB
	}
	return false
}

func compareMonths(a, b *string, reverse bool) bool {
	months := map[string]int{
		"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4, "May": 5, "Jun": 6,
		"Jul": 7, "Aug": 8, "Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
	}
	numA, okA := months[*a]
	numB, okB := months[*b]
	if okA && okB {
		if reverse {
			*a, *b = strconv.Itoa(numA), strconv.Itoa(numB)
			return numA > numB
		}
		*a, *b = strconv.Itoa(numA), strconv.Itoa(numB)
		return numA < numB
	}
	return false
}

func compareHumanNumbers(a, b *string, reverse bool) bool {
	numA, unitA := parseHumanNumber(*a)
	numB, unitB := parseHumanNumber(*b)
	if numA != numB {
		if reverse {
			*a, *b = fmt.Sprintf("%f%s", numA, unitA), fmt.Sprintf("%f%s", numB, unitB)
			return numA > numB
		}
		*a, *b = fmt.Sprintf("%f%s", numA, unitA), fmt.Sprintf("%f%s", numB, unitB)
		return numA < numB
	}
	if unitA != unitB {
		if reverse {
			*a, *b = fmt.Sprintf("%f%s", numA, unitA), fmt.Sprintf("%f%s", numB, unitB)
			return unitA > unitB
		}
		*a, *b = fmt.Sprintf("%f%s", numA, unitA), fmt.Sprintf("%f%s", numB, unitB)
		return unitA < unitB
	}
	return false
}

func parseHumanNumber(s string) (float64, string) {
	var numStr string
	for i, r := range s {
		if (r < '0' || r > '9') && r != '.' {
			num, err := strconv.ParseFloat(s[:i], 64)
			if err != nil {
				return 0, ""
			}
			return num, s[i:]
		}
		numStr += string(r)
	}
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, ""
	}
	return num, ""
}
