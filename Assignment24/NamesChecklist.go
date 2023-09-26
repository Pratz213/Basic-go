package main

import (
	"fmt"
	"strings"
)

func main() {
	//declaring map (where key is string and value is int)
	names := make(map[string]int)

	// Accept names as input until "done" is entered
	fmt.Println("Enter a name or 'done' : ")
	for {
		var name string
		fmt.Scanln(&name)

		if strings.ToLower(name) == "done" {
			break
		}
		names[strings.ToLower(name)] = 0
	}

	// Accept a name to check if it exists in the hash map
	fmt.Print("Enter a name to check: ")
	var checkName string
	fmt.Scanln(&checkName)

	_, exists := names[strings.ToLower(checkName)] //returns value and present in map or not

	// Check if the entered name exists in the hash map
	if exists {
		fmt.Printf("name %s found in list\n", checkName)
	} else {
		fmt.Printf("name %s not found in list\n", checkName)
	}

}
