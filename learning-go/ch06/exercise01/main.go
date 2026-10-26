package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	p1 := MakePerson("Francesco", "Totti", 50)
	p2 := MakePersonPointer("Daniele", "De Rossi", 50)
	fmt.Println(p1.FirstName, p1.LastName, p1.Age)
	fmt.Println(p2.FirstName, p2.LastName, p2.Age)
}
