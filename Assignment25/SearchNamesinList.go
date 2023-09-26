package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var candidates []string

	// Accept candidate names until "done" is entered
	fmt.Println("Enter a candidate name or enter 'done' : ")
	for {
		var name string
		fmt.Scanln(&name)

		if strings.ToLower(name) == "done" {
			break
		}
		candidates = append(candidates, name)
	}

	var searchPattern string
	fmt.Print("Enter a search pattern : ")
	fmt.Scanln(&searchPattern)

	// Search and prints candidate names
	for _, candidate := range candidates {
		matched, _ := regexp.MatchString(searchPattern, candidate)
		if matched {
			fmt.Println(candidate)
		}
	}
}
