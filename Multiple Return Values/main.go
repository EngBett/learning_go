package main

import "fmt"

func add_sub(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	add_res, sub_res := add_sub(15, 12)

	fmt.Println("Addition", add_res)
	fmt.Println("Subtraction", sub_res)

	add_result, _ := add_sub(12, 15)

	fmt.Println(add_result)
}
