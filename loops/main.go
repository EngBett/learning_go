//Go has only one loop the for loop
package main

import "fmt"

func main() {
	i := 0

	//loop style 1
	for i <= 10{
		fmt.Println(i)
		i += 1
	}

	//line breaks
	fmt.Println()
	fmt.Println()

	//loop style 2
	for j:=0;j<10;j++{
		fmt.Println(j)
	}

	//line breaks
	fmt.Println()
	fmt.Println()

	//odd numbers only
	for k:=0;k<=10;k++{
		if k%2 == 0{
			continue
		}else {
			fmt.Println(k)
		}
	}

	//line breaks
	fmt.Println()
	fmt.Println()

	//infinite loop
	for{
		fmt.Println("Keep printing")
	}

	// use continue to skip and break to break the loop
}
