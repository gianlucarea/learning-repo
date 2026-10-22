package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var result []int
	var result2 []int
	var result3 []int
	for i := 0; i < 100; i++ {
		result = append(result, rand.Intn(100))
	}
	j := 0
	for j < 100 {
		result2 = append(result2, rand.Intn(100))
		j++
	}
	for _ = range 100 {
		result3 = append(result3, rand.Intn(100))
	}

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
}
