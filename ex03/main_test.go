package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSort(t *testing.T) {
	input := "4\n2\n3\n1\n"
	expected := "1\n2\n3\n4\n"
	runTest(t, input, expected, []string{"-n"})

	input = "banana\napple\ncherry\n"
	expected = "apple\nbanana\ncherry\n"
	runTest(t, input, expected, []string{})

	input = "Jan\nFeb\nMar\n"
	expected = "Feb\nJan\nMar\n"
	runTest(t, input, expected, []string{"-M"})

	input = "5K\n3K\n10K\n"
	expected = "10K\n3K\n5K\n"
	runTest(t, input, expected, []string{"-h"})
}

func runTest(t *testing.T, input, expected string, args []string) {
	inputFile, err := ioutil.TempFile("", "input")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(inputFile.Name())

	outputFile, err := ioutil.TempFile("", "output")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(outputFile.Name())

	if _, err := inputFile.WriteString(input); err != nil {
		t.Fatal(err)
	}
	inputFile.Close()

	origArgs := os.Args
	defer func() { os.Args = origArgs }()
	os.Args = append([]string{os.Args[0], inputFile.Name(), outputFile.Name()}, args...)

	main()

	outputBytes, err := ioutil.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	output := string(outputBytes)
	if output != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s\n", expected, output)
	}
}
