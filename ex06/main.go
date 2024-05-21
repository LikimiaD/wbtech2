package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	fields    = flag.String("f", "", "поля для вывода (например, '1,3')")
	delimiter = flag.String("d", ",", "используемый разделитель (по умолчанию - запятая)")
	separated = flag.Bool("s", false, "печатать только строки с разделителем")
)

func readFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		cols := strings.Split(line, *delimiter)
		fieldIndexes := parseFields(*fields)

		var result []string
		for _, index := range fieldIndexes {
			if index > 0 && index <= len(cols) {
				result = append(result, cols[index-1])
			}
		}

		fmt.Println(strings.Join(result, *delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %v\n", err)
	}
}

func parseFields(fields string) []int {
	var indexes []int
	if fields != "" {
		parts := strings.Split(fields, ",")
		for _, part := range parts {
			var index int
			if _, err := fmt.Sscanf(part, "%d", &index); err == nil {
				indexes = append(indexes, index)
			} else {
				fmt.Printf("error while reading parts: %s\n", err.Error())
			}
		}
	}
	return indexes
}

func openFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening file: %s\n", err.Error())
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("error while closing file: %s\n", err.Error())
		}
	}()

	readFile(file)
}

func main() {
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("no file specified")
		return
	}

	for _, fileName := range files {
		openFile(fileName)
	}
}
