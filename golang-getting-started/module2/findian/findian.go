package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Write a program which prompts the user to enter a string. The program searches
// through the entered string for the characters ‘i’, ‘a’, and ‘n’.
// The program should print “Found!” if the entered string starts with the character
// ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should
//  print “Not Found!” otherwise. The program should not be case-sensitive,
// so it does not matter if the characters are upper-case or lower-case.

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string:")
	// Read all the input until a new line and store it in input
	input, _ = reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	found := true
	// Check if the input starts with i
	found = found && input[0] == 'i'
	// Check if the input contains a
	found = found && strings.Contains(input, "a")
	// Check if the input ends with n
	found = found && input[len(input)-1] == 'n'

	if found {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

// Examples: The program should print “Found!”
// for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
// The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
