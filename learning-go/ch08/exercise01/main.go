package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func doubleDown[T Number](value T) T {
	return value * 2
}

func main() {
	fmt.Println(doubleDown(10))
	fmt.Println(doubleDown(20.20))
}
