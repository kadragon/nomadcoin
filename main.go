package main

import "fmt"

func plus(a, b int, name string) (int, string) {
	return a + b, name
}

func multiplus(a ...int) int {
	var total int = 0
	for _, val := range a {
		total += val
	}

	return total
}

func main() {
	result, name := plus(2, 2, "nico")
	fmt.Println(result, name)

	result = multiplus(2, 2, 4, 5, 1, 2, 3)
	fmt.Println(result)

	var names string = "Nicolas!! is my name"
	for idx, letter := range names {
		fmt.Println(idx, letter, string(letter))
	}
}
