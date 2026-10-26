package main

import "fmt"

func UpdateSlice(input []string, lastElement string) {
	input[len(input)-1] = lastElement
	fmt.Println(input) // Prints [Ciao Miao Hello]
}

func GrowSlice(input []string, lastElement string) {
	input = append(input, lastElement)
	fmt.Println(input) // Prints [Ciao Miao Bau Hello]
}
func main() {
	var input1 []string = []string{"Ciao", "Miao", "Bau"}
	var input2 []string = []string{"Ciao", "Miao", "Bau"}
	fmt.Println(input1) // Prints [Ciao Miao Bau]
	fmt.Println(input2) // Prints [Ciao Miao Bau]
	UpdateSlice(input1, "Hello")
	GrowSlice(input2, "Hello")
	fmt.Println(input1) // Prints [Ciao Miao Hello]
	fmt.Println(input2) // Prints [Ciao Miao Bau]
}
