package main

import (
	"fmt"
)

type Animal struct {
	name string
	sound string
}

func (animal Animal) MakeSound() {
	fmt.Printf("%s is making sound: %s", animal.name, animal.sound)
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
}

func main() {
	c := Cat{Animal{
		name:  "Cat",
		sound: "Meow Meow",
	}}

	d := Dog{Animal{
		name:  "Dog",
		sound: "Bow Bow",
	}}

	c.MakeSound()
	d.MakeSound()
}