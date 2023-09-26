/*
	package : go.core/assignments

	program reads a file provided as a command-line argument, processes each line to count the occurrences of words
	(excluding common words and single-letter words and ignoring punctuations),and then prints out the word count along with their frequencies.
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

	filename := getFileName()

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

	//loop that continues to read till end of input file.
	for scanner.Scan() {
		text := removePunctuation(scanner.Text())
		if len(text) > 1 && !ignoreWord(text) {
			wordCount[text]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	return wordCount
}

func getFileName() string {
	// Check if a filename is provided as a command-line argument
	if len(os.Args) != 2 {
		fmt.Println("invalid Number of input: ")
		os.Exit(1)
	}
	filename := os.Args[1]
	return filename
}

// ignoreWord checks if a given word should be ignored based on a predefined list.
func ignoreWord(word string) bool {
	ignoredWords := []string{"the", "if", "an", "in", "or", "is", "and", "to", "of", "do", "if", "is"}

	for _, ignored := range ignoredWords {
		if word == ignored {
			return true
		}
	}
	return false
}

func removePunctuation(word string) string {
	var result strings.Builder
	for _, char := range word {
		if !unicode. IsPunct(char) {
			result.WriteRune(char)
		}
		fmt.Print("\n", result.String())
	}
	return result.String()
}
 