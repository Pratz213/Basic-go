package main

import "fmt"

func main() {
	defer fmt.Println(one())

	fmt.Println("Hello")
	defer fmt.Println(two())

	fmt.Println("world")
	defer fmt.Println(three())
	num()
}

func two() string {
	return "two"

}

func three() string {
	return "three"
}

func one() string {
	return "one"
}

func num() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
