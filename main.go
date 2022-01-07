package main

import "fmt"

func main() {
	var x int = 405940594059
	fmt.Printf("%b\n", x)
	fmt.Printf("%o\n", x)
	fmt.Printf("%x\n", x)

	xAsBinary := fmt.Sprintf("%b\n", x)
	fmt.Println(x, xAsBinary)
}
