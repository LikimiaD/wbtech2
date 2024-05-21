package main

import (
	"testing"
)

func resetOptions() {
	*after = 0
	*before = 0
	*context = 0
	*count = false
	*ignore = false
	*invert = false
	*fixed = false
	*lineNum = false
}

func setOptions(options map[string]interface{}) {
	if val, ok := options["after"]; ok {
		*after = val.(int)
	}
	if val, ok := options["before"]; ok {
		*before = val.(int)
	}
	if val, ok := options["context"]; ok {
		*context = val.(int)
	}
	if val, ok := options["count"]; ok {
		*count = val.(bool)
	}
	if val, ok := options["ignore"]; ok {
		*ignore = val.(bool)
	}
	if val, ok := options["invert"]; ok {
		*invert = val.(bool)
	}
	if val, ok := options["fixed"]; ok {
		*fixed = val.(bool)
	}
	if val, ok := options["lineNum"]; ok {
		*lineNum = val.(bool)
	}
}

func runTest(t *testing.T, name string, lines []string, pattern string, options map[string]interface{}, expected []string) {
	t.Run(name, func(t *testing.T) {
		resetOptions()
		setOptions(options)

		result := grep(lines, pattern)
		if !equal(result, expected) {
			t.Errorf("got %v, want %v", result, expected)
		}
	})
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGrep(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		pattern  string
		options  map[string]interface{}
		expected []string
	}{
		{
			name:    "BasicMatch",
			lines:   []string{"hello world", "goodbye world", "hello again"},
			pattern: "hello",
			options: map[string]interface{}{},
			expected: []string{
				"hello world",
				"hello again",
			},
		},
		{
			name:    "AfterMatch",
			lines:   []string{"line1", "line2", "hello world", "line4", "line5"},
			pattern: "hello",
			options: map[string]interface{}{"after": 2},
			expected: []string{
				"hello world",
				"line4",
				"line5",
			},
		},
		{
			name:    "BeforeMatch",
			lines:   []string{"line1", "line2", "hello world", "line4", "line5"},
			pattern: "hello",
			options: map[string]interface{}{"before": 2},
			expected: []string{
				"line1",
				"line2",
				"hello world",
			},
		},
		{
			name:    "IgnoreCase",
			lines:   []string{"Hello world", "goodbye world", "hello again"},
			pattern: "hello",
			options: map[string]interface{}{"ignore": true},
			expected: []string{
				"Hello world",
				"hello again",
			},
		},
		{
			name:    "InvertMatch",
			lines:   []string{"hello world", "goodbye world", "hello again"},
			pattern: "hello",
			options: map[string]interface{}{"invert": true},
			expected: []string{
				"goodbye world",
			},
		},
		{
			name:    "FixedString",
			lines:   []string{"hello world", "hello*world", "hello again"},
			pattern: "hello*world",
			options: map[string]interface{}{"fixed": true},
			expected: []string{
				"hello*world",
			},
		},
		{
			name:    "LineNum",
			lines:   []string{"hello world", "goodbye world", "hello again"},
			pattern: "hello",
			options: map[string]interface{}{"lineNum": true},
			expected: []string{
				"1:hello world",
				"3:hello again",
			},
		},
		{
			name:    "CountMatch",
			lines:   []string{"hello world", "goodbye world", "hello again"},
			pattern: "hello",
			options: map[string]interface{}{"count": true},
			expected: []string{
				"2",
			},
		},
	}

	for _, tt := range tests {
		runTest(t, tt.name, tt.lines, tt.pattern, tt.options, tt.expected)
	}
}
