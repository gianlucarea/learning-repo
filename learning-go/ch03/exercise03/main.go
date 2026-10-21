package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	first_employee := Employee{"John", "Doe", 1}
	second_employee := Employee{
		firstName: "Bob",
		lastName:  "Dylan",
		id:        2,
	}
	var third_emplyee Employee
	third_emplyee.firstName = "Frank"
	third_emplyee.lastName = "Sinatra"
	third_emplyee.id = 3

	fmt.Println(first_employee)
	fmt.Println(second_employee)
	fmt.Println(third_emplyee)
}
