package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Go")

	//mathematical expressions
	fmt.Println("1+1 is equal to", 1 + 1)
	fmt.Println("1-1 is equal to", 1 - 1)
	fmt.Println("1*1 is equal to", 1 * 1)
	fmt.Println("1/1 is equal to", 1 / 1)
	fmt.Println("1%1 is equal to", 1 % 1)

	//floating points
	fmt.Println(1.15 - 0.0275)

	//booleans
	fmt.Println(5 == 5)
	fmt.Println(6 == 5)

	fmt.Println("The time is :", time.Now())
}
