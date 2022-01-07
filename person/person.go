package person

import "log"

type Person struct {
	name string
	age  int
}

func (p Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age

	log.Printf("SetDetails %v\n", p)
}

func (p *Person) SetDetailsWithPointer(name string, age int) {
	p.name = name
	p.age = age

	log.Printf("SetDetails Pointer %v\n", p)
}
