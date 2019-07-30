package main

import "fmt"

func main() {
	i:=15

	switch i {
	case 1:
		fmt.Println("1")
	case 50,10:
		fmt.Println("50 or 10")
	case 100:
		fmt.Println("100")
	default:
		fmt.Println("Unknown")
	}

	j := 100

	switch {
	case j<500:
		fmt.Println("Less than 500")
	case j > 500:
		fmt.Println("Greater than 500")
	default:
		fmt.Println("Is 500")
	}
}
