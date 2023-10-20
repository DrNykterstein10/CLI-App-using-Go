/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 4 {
		option := args[2]
		filePath := args[3]

		file, err := os.Open(filePath)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		switch option {
		case "-c":
			numBytes := byteCount(file)
			fmt.Println(numBytes, filePath)

		case "-l":
			numLines := lineCount(file)
			fmt.Println(numLines, filePath)

		case "-w":
			numWords := wordCount(file)
			fmt.Println(numWords, filePath)

		case "-m":
			numCharacters := characterCount(file)
			fmt.Println(numCharacters, filePath)

		default:
			panic("Invalid option selected")
		}

	}

	if len(args) == 3 {
		filePath := args[2]

		file, err := os.Open(filePath)

		if err != nil {
			fmt.Println(err)
			return
		}

		numBytes := byteCount(file)
		numLines := lineCount(file)
		numWords := wordCount(file)

		fmt.Println(numLines, numWords, numBytes, filePath)

	}
}

func byteCount(file *os.File) int {
	fileBytes, err := file.Stat()

	if err != nil {
		fmt.Println("Could not read file")
	}

	return int(fileBytes.Size())
}

func lineCount(file *os.File) int {
	reader := bufio.NewReader(file)

	numLines := 0

	for {
		_, _, err := reader.ReadLine()

		if err != nil {
			break
		}
		numLines++
	}

	return numLines
}

func wordCount(file *os.File) int {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("Could not seek file")
		return 0
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)
	defer file.Close()

	numWords := 0
	for fileScanner.Scan() {
		numWords++
	}

	return numWords
}

func characterCount(file *os.File) int {
	data := make([]byte, 1024)
	defer file.Close()

	numCharacters := 0

	for {
		n, err := file.Read(data)

		if err != nil && err != io.EOF {
			fmt.Printf("Error reading from the file: %v", err)
			return 0
		}

		if n == 0 {
			break
		}

		numCharacters += n
	}

	return numCharacters
}
