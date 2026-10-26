package main

import "fmt"

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func main() {
	fmt.Println("Hello")
}
