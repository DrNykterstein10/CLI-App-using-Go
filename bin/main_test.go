package main

import (
	"os"
	"testing"
)

func openTestFile(t *testing.T) *os.File {
	file, err := os.Open("../test.txt")
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	return file
}

func TestByteCount(t *testing.T) {
	file := openTestFile(t)
	defer file.Close()

	result := byteCount(file)
	expected := 21

	if result != expected {
		t.Errorf("Expected byte count: %v, but got: %v", expected, result)
	}
}

func TestLineCount(t *testing.T) {
	file := openTestFile(t)
	defer file.Close()

	result := lineCount(file)
	expected := 1

	if result != expected {
		t.Errorf("Expected line count: %v, but got: %v", expected, result)
	}
}

func TestWordCount(t *testing.T) {
	file := openTestFile(t)
	defer file.Close()

	result := wordCount(file)
	expected := 5

	if result != expected {
		t.Errorf("Expected word count: %v, but got: %v", expected, result)
	}
}

func TestCharacterCount(t *testing.T) {
	file := openTestFile(t)
	defer file.Close()

	result := characterCount(file)
	expected := 21

	if result != expected {
		t.Errorf("Expected character count: %v, but got: %v", expected, result)
	}
}
