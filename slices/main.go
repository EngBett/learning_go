package main

import "fmt"

func main() {
	s := make([]int, 3)
	s[1] = 3
	s[0] = 5
	s[2] = 8
	fmt.Println(s)
}
