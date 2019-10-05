package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func add3(a, b, c int) int {
	return a + b + c
}

func test() {
	fmt.Println("This is a fuction of type void")
}

func main() {
	ans := add(15, 18)

	fmt.Println(ans)

	ans2 := add3(15, 18, 28)
	fmt.Println(ans2)

	test()

}
