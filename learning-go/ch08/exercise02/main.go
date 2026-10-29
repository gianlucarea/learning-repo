package main

import (
	"fmt"
)

// A type satisfies this only if it has Print() and it is either a int or float
type Printable interface {
	fmt.Stringer
	~int | ~float64
}

// Named type int and float
type PrintInt int
type PrintFloat float64

// We define String() for both named type and so we satisfy the interface
func (pi PrintInt) String() string {
	return fmt.Sprintf("%d", pi)
}
func (pf PrintFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

// Generic Function
func PrintIt[T Printable](t T) {
	fmt.Println(t)
}

func main() {
	var i PrintInt = 20
	PrintIt(i)

	var f PrintFloat = 10.23
	PrintIt(f)
}
