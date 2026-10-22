package main

import "fmt"

func main() {
	var total int
	for i := range 10 {
		// total := total + i  The outside print will be 0 because this shadow the total
		total = total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}
