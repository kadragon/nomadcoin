package main

import "fmt"

const outname string = "nico"

func main() {
	var name string = "nico"
	name = "12"

	// in function only
	newName := "nico"

	fmt.Println(name, newName, outname)
}
