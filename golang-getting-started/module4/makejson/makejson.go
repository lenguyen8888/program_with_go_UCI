package main

// Write a program which prompts the user to first enter a name, and then enter
// an address. Your program should create a map and add the name and address
// to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from
// the map, and then your program should print the JSON object.

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name, address string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a name:")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Enter an address:")
	address, _ = reader.ReadString('\n')
	address = strings.TrimSpace(address)

	var m map[string]string = make(map[string]string)
	m["name"] = name
	m["address"] = address

	jsonObj, _ := json.Marshal(m)
	fmt.Println(string(jsonObj))
}
