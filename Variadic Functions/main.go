package main

import "fmt"

func mult(nums ...int) int {
	total := 10

	for _, num := range nums {
		total *= num
	}

	return total
}

func main() {
	// Println() is an example of a variadic function
	fmt.Println("this", "is", "an", "example", "of", "a", "variadic", "fuction")

	//call a variadic fuction
	fmt.Println(mult(1, 6, 8, 15, 25))

	fmt.Println("")

	nums := []int{1, 6, 8, 15, 25}
	fmt.Println(mult(nums...))

}
