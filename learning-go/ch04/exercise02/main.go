package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var random_slice []int
	for _ = range 100 {
		random_slice = append(random_slice, rand.Intn(100))
	}
	fmt.Println(random_slice)
	for _, v := range random_slice {
		if v%3 == 0 && v%2 == 0 {
			fmt.Println("Six")
			continue
		}
		if v%3 == 0 {
			fmt.Println("Three")
		} else if v%2 == 0 {
			fmt.Println("Two")
		} else {
			fmt.Println("Never Mind")
		}
	}
}
