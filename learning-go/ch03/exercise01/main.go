package main

import "fmt"

func main() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	first_two_greetings := greetings[0:2]
	middle_greetings := greetings[1:4]
	last_two_greetings := greetings[3:5]
	fmt.Println(greetings)
	fmt.Println(first_two_greetings)
	greetings[3] = "Ciao" // Modifying the main greetings will modify the subslice because it share memory!
	fmt.Println(middle_greetings)
	fmt.Println(last_two_greetings)

}
