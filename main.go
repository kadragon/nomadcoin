package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) seyHello() {
	fmt.Printf("Hello! My name is %s and I'm %d\n", p.name, p.age)
}

func (p person) seyKoreaAge() {
	fmt.Printf("My Korea Age is %d\n", p.age+1)
}

func main() {
	nico := person{name: "nico", age: 12}
	nico.seyHello()
	nico.seyKoreaAge()
}
