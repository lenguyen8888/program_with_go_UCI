package main

// Write a program which reads information from a file and represents it in a slice
// of structs. Assume that there is a text file which contains a series of names.
// Each line of the text file has a first name and a last name, in that order,
// separated by a single space on the line.

// Your program will define a name struct which has two fields,
// fname for the first name, and lname for the last name.
// Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file.
// Your program will successively read each line of the text file and create
// a struct which contains the first and last names found in the file.
// Each struct created will be added to a slice, and after all lines have been
// read from the file, your program will have a slice containing one struct for
// each line in the file. After reading all lines from the file, your program
// should iterate through your slice of structs and print the
// first and last names found in each struct.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Enter file name: ")
	var fileName string
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	names := make([]Name, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		name := strings.Split(line, " ")
		names = append(names, Name{fname: name[0], lname: name[1]})
	}

	for _, name := range names {
		fmt.Printf("fname: %s, lname: %s\n", name.fname, name.lname)
	}
}
