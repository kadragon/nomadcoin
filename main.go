package main

import (
	"log"

	"github.com/kadragon/nomadcoin/person"
)

func main() {
	nico := person.Person{}

	nico.SetDetails("nico", 12)
	log.Println(nico)

	nico.SetDetailsWithPointer("nico", 12)
	log.Println(nico)
}
