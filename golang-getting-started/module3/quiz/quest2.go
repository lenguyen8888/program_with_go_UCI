package main

import "fmt"

func main() {
	x := [...]int{4, 8, 5}
	y := x[0:2]
	z := x[1:3]
	y[0] = 1
	z[1] = 3
	fmt.Print(x)
}

// What is the output of the following program?
// [1 8 3]
