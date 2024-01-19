package main

import "fmt"

func main() {
	x := [...]int{1, 2, 3, 4, 5}
	y := x[0:2]
	z := x[1:4]
	fmt.Print(len(y), cap(y), len(z), cap(z))
}

// What is the output of the following program?
// 2 5 3 4
