package main

import "fmt"

func main() {
	var a int = 2
	var b int = a
	var c *int = &a

	a = 12

	fmt.Println(b)
	fmt.Println(*c)
}
