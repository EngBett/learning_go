package main

import "fmt"

//strings with spaces dont really work here
func main() {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scan(&input)
	fmt.Print(input)
}
