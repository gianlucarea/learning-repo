package main

import "fmt"

const value = 20

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	fmt.Printf("Byte + 1: %d\n", b+1)
	fmt.Printf("Small Int + 1: %d\n", smallI+1)
	fmt.Printf("Big Int + 1: %d\n", bigI+1)
}
