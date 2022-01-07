package main

import "fmt"

func main() {
	// slice
	foods := []string{"potato", "pizza", "pasta"}

	// array
	// foods := [3]string{"potato", "pizza", "pasta"}
	for _, food := range foods {
		fmt.Println(food)
	}

	foods = append(foods, "tomato")
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
}
