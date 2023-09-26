/*
	package : go.core/assignments

	program reads a file provided as a command-line argument, processes each line to count the occurrences of words
	and then prints out the word count along with their frequencies.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	filename := getInput()

	wordCount := openAndReadFile(filename)
	// Print the word count
	fmt.Println("Word Count :")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func openAndReadFile(filename string) map[string]int {

	// open the file
	file, err := os.Open(filename)
	wordCount := make(map[string]int)

	// Check for the error that occurred during the opening of the file
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := removePunctuation(scanner.Text())
		wordCount[strings.ToLower(word)]++
	}
	return wordCount
}

func getInput() string {
	// Check if a filename is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("invalid Number of input: ")
		os.Exit(1)
	}

	filename := os.Args[1]
	return filename
}

func removePunctuation(word string) string {
	var result strings.Builder
	for _, char := range word {
		if !unicode.IsPunct(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}
